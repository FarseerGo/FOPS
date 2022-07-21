package building

import (
	"context"
	"fops/domain/_/eumBuildStatus"
	"fops/domain/building/build"
	"fops/domain/building/device"
	"fs/core/container"
)

type BuildService struct {
	repository       build.Repository
	logWriteDevice   device.ILogWriteDevice
	dockerDevice     device.IDockerDevice
	directoryDevice  device.IDirectoryDevice
	gitDevice        device.IGitDevice
	kubectlDevice    device.IKubectlDevice
	copyToDistDevice device.ICopyToDistDevice
	progress         chan string
	ctx              context.Context
	cancel           context.CancelFunc
}

func NewBuildService() *BuildService {
	ctx, cancel := context.WithCancel(context.Background())
	return &BuildService{
		repository:       container.Resolve[build.Repository](),
		logWriteDevice:   container.Resolve[device.ILogWriteDevice](),
		dockerDevice:     container.Resolve[device.IDockerDevice](),
		directoryDevice:  container.Resolve[device.IDirectoryDevice](),
		gitDevice:        container.Resolve[device.IGitDevice](),
		kubectlDevice:    container.Resolve[device.IKubectlDevice](),
		copyToDistDevice: container.Resolve[device.ICopyToDistDevice](),
		ctx:              ctx,
		cancel:           cancel,
	}
}

// Build 构建
func (service *BuildService) Build() {
	var buildDo = service.repository.GetUnBuildInfo()
	if buildDo.Id < 1 {
		return
	}
	var isUpdate = service.repository.SetBuilding(buildDo.Id)

	// 没有更新成功，说明已经被抢了
	if isUpdate == 0 {
		return
	}

	service.progress = service.logWriteDevice.CreateProgress(buildDo.Id)

	// 定义环境变量
	var projectGitRoot = service.gitDevice.GetGitPath(buildDo.GetGitMaster().Hub)
	var dockerHub = service.dockerDevice.GetDockerHub(buildDo.Docker.Hub)
	var dockerImage = service.dockerDevice.GetDockerImage(buildDo.Docker.Hub, buildDo.Project.Name, buildDo.BuildNumber)
	var gitName = service.gitDevice.GetName(buildDo.GetGitMaster().Hub)
	buildDo.GenerateEnv(projectGitRoot, dockerHub, dockerImage, gitName)

	defer service.catch(buildDo)

	// 打印环境变量
	buildDo.PrintEnv(service.progress)

	// 前置检查
	service.directoryDevice.Check(service.progress)

	// 拉取主仓库及依赖仓库
	service.checkResult(service.gitDevice.CloneOrPullAndDependent(buildDo.Gits, service.progress, service.ctx), buildDo.Id)

	// 登陆镜像仓库(先登陆，如果失败了，后则面也不需要编译、打包了)
	service.checkResult(service.dockerDevice.Login(buildDo.Docker.Hub, buildDo.Docker.UserName, buildDo.Docker.UserPwd, service.progress, buildDo.Env, service.ctx), buildDo.Id)

	// 将需要打包的源代码，复制到dist目录
	service.copyToDistDevice.Copy(buildDo.Gits, buildDo.Env, service.progress)

	// 生成Dockerfile文件
	buildDo.GenerateDockerfileContent()
	service.dockerDevice.CreateDockerfile(buildDo.Project.Name, buildDo.Docker.DockerfileContent, service.ctx)

	// docker打包
	service.checkResult(service.dockerDevice.Build(buildDo.Env, service.progress, service.ctx), buildDo.Id)

	// docker上传
	service.checkResult(service.dockerDevice.Push(buildDo.Env, service.progress, service.ctx), buildDo.Id)

	// k8s更新
	service.checkResult(service.kubectlDevice.SetImages(buildDo.Cluster, buildDo.Project.Name, buildDo.Env.DockerImage, buildDo.Project.K8SControllersType, service.progress, service.ctx), buildDo.Id)

	service.success(buildDo, service.progress)
}

// CheckResult 检查结构
func (service *BuildService) checkResult(result bool, buildId int) {
	status := service.repository.GetStatus(buildId)
	if status == eumBuildStatus.Finish {
		panic("手动取消，退出构建。")
	}

	if !result {
		panic("exit")
	}
}

// 设置任务失败
func (service *BuildService) fail(buildDo build.DomainObject, progress chan string) {
	progress <- "---------------------------------------------------------"
	progress <- "执行失败，退出构建。"

	// 发布事件
	build.FinishedEvent{ProjectId: buildDo.Project.Id, BuildId: buildDo.Id, ClusterId: buildDo.Cluster.Id, IsSuccess: false}.PublishEvent()

	service.repository.Cancel(buildDo.Id)
}

// 设置任务成功
func (service *BuildService) success(buildDo build.DomainObject, progress chan string) {
	progress <- "---------------------------------------------------------"
	progress <- "构建完成。"

	// 发布事件
	build.FinishedEvent{ProjectId: buildDo.Project.Id, BuildId: buildDo.Id, ClusterId: buildDo.Cluster.Id, IsSuccess: true}.PublishEvent()

	service.repository.Success(buildDo.Id)
}

func (service *BuildService) catch(buildDo build.DomainObject) {
	if err := recover(); err != nil {
		service.cancel()
		msg := err.(string)
		if msg != "exit" {
			service.progress <- msg
		}
		service.fail(buildDo, service.progress)

	}
}
