package project

import (
	domain "fops/domain/_"
	"fops/domain/_/eumBuildType"
	"fops/domain/_/eumK8SControllers"
	"strings"
	"time"
)

type DomainObject struct {
	//     主键
	Id int
	//     项目名称
	Name string
	//     应用ID（链路追踪）
	AppId string
	//     程序入口名称
	EntryPoint string
	//     程序启动端口
	EntryPort int
	//     访问域名
	Domain string
	//     项目路径
	Path string
	//     项目组ID
	GroupId int
	//     GIT
	GitId int
	//     依赖的GIT库（会同时拉取依赖的GIT库）
	DependentGitIds []int
	//     DockerHub模板
	DockerHub int
	//     DockerfileTpl模板
	DockerfileTpl int
	//     K8SDeployment模板
	K8STplDeployment int
	//     K8SIngress模板
	K8STplIngress int
	//     K8SService模板
	K8STplService int
	//     K8SConfig模板
	K8STplConfig int
	//     K8S模板自定义变量(K1=V1,K2=V2)
	K8STplVariable string
	//     镜像版本
	DockerVer string
	//     构建方式
	BuildType eumBuildType.Enum
	//     Shell脚本
	ShellScript string
	//     K8S负载类型
	K8SControllersType eumK8SControllers.Enum
	// 集群版本
	ClusterVer map[int]*domain.ClusterVerVO
}

func New() DomainObject {
	return DomainObject{
		ClusterVer: map[int]*domain.ClusterVerVO{},
	}
}

// CheckForSave 保存项目前，需对属性作检查
func (receiver *DomainObject) CheckForSave() {
	if receiver.Path == "" {
		receiver.Path = "/"
	} else if !strings.HasPrefix("/", receiver.Path) {
		receiver.Path = "/" + receiver.Path
	}

	if receiver.ShellScript == "" {
		receiver.ShellScript = ""
	}

	if receiver.Domain != "" {
		receiver.Domain = strings.ToLower(receiver.Domain)
		receiver.Domain = strings.Replace(receiver.Domain, "http://", "", -1)
		receiver.Domain = strings.Replace(receiver.Domain, "https://", "", -1)
	}

	if receiver.ClusterVer == nil {
		receiver.ClusterVer = map[int]*domain.ClusterVerVO{}
	}
}

// UpdateBuildVer 当构建失败时，记录失败时间、失败时的构建ID
func (receiver *DomainObject) UpdateBuildVer(isSuccess bool, clusterId int, buildId int) {
	if _, ok := receiver.ClusterVer[clusterId]; !ok {
		receiver.ClusterVer[clusterId] = &domain.ClusterVerVO{
			DockerVer: "0",
		}
	}

	// 当构建成功时，记录发布时间、发布时的构建ID
	if isSuccess {
		receiver.ClusterVer[clusterId].DockerVer = receiver.DockerVer
		receiver.ClusterVer[clusterId].DeploySuccessAt = time.Now()
		receiver.ClusterVer[clusterId].BuildSuccessId = buildId
	} else // 当构建失败时，记录失败时间、失败时的构建ID
	{
		receiver.ClusterVer[clusterId].DeployFailAt = time.Now()
		receiver.ClusterVer[clusterId].BuildFailId = buildId
	}
}
