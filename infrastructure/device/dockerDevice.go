package device

import (
	"context"
	"fops/domain/building/build/event"
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/exec"
	"github.com/farseernet/farseer.go/utils/file"
	"github.com/farseernet/farseer.go/utils/str"
	"os"
	"strconv"
)

func init() {
	_ = container.Register(func() devicer.IDockerDevice { return &dockerDevice{} })
}

type dockerDevice struct {
}

func (dockerDevice) GetDockerHub(dockerHubAddress string) string {
	var dockerHub = "localhost"
	if dockerHubAddress != "" {
		dockerHub = dockerHubAddress
		dockerHub = str.CutRight(dockerHub, "/")
	}
	return dockerHub
}

func (device dockerDevice) GetDockerImage(dockerHubAddress string, projectName string, buildNumber int) string {
	return device.GetDockerHub(dockerHubAddress) + ":" + projectName + "-" + strconv.Itoa(buildNumber)
}

func (dockerDevice) Login(dockerHub string, loginName string, loginPwd string, progress chan string, env vo.EnvVO, ctx context.Context) bool {
	progress <- "---------------------------------------------------------"
	if dockerHub != "" && loginName != "" {
		var result = exec.RunShellContext("docker login "+dockerHub+" -u "+loginName+" -p "+loginPwd, progress, env.ToMap(), "", ctx)
		if result != 0 {
			progress <- "镜像仓库登陆失败。"
			return false
		}
	}

	progress <- "镜像仓库登陆成功。"
	return true
}

func (dockerDevice) ExistsDockerfile(dockerfilePath string) bool {
	return file.IsExists(dockerfilePath)
}

func (dockerDevice) CreateDockerfile(projectName string, dockerfileContent string, ctx context.Context) {
	if file.IsExists(vo.DockerfilePath) {
		_ = os.RemoveAll(vo.DockerfilePath)
	}
	file.WriteString(vo.DockerfilePath, dockerfileContent)
}

func (dockerDevice) Build(env vo.EnvVO, progress chan string, ctx context.Context) bool {
	progress <- "---------------------------------------------------------"
	progress <- "开始镜像打包。"

	// 打包
	var result = exec.RunShellContext("docker build -t "+env.DockerImage+" --network=host -f "+vo.DockerfilePath+" "+vo.DistRoot, progress, env.ToMap(), vo.DistRoot, ctx)
	if result == 0 {
		progress <- "镜像打包完成。"
	} else {
		progress <- "镜像打包出错了。"
	}
	return result == 0
}

func (dockerDevice) Push(env vo.EnvVO, progress chan string, ctx context.Context) bool {
	// 上传
	var result = exec.RunShellContext("docker push "+env.DockerImage, progress, env.ToMap(), "", ctx)

	if result == 0 {
		progress <- "镜像上传完成。"

		// 上传成功后，需要更新项目中的镜像版本属性
		event.DockerPushedEvent{
			BuildNumber: env.BuildNumber,
			ProjectId:   env.ProjectId,
		}.PublishEvent()
		return true
	}

	progress <- "镜像上传出错了。"
	return false
}
