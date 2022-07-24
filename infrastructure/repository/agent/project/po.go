package project

import (
	domain "fops/domain/_"
	"fops/domain/_/eumBuildType"
	"fops/domain/_/eumK8SControllers"
)

type PO struct {
	Id int `gorm:"primaryKey"`

	// 项目名称
	Name string
	// 应用ID（链路追踪）
	AppId string
	// 程序入口名称
	EntryPoint string
	// 程序启动端口
	EntryPort int
	// 访问域名
	Domain string
	// 项目组ID
	GroupId int
	// GIT
	GitId int
	// 依赖的GIT库（会同时拉取依赖的GIT库）
	//[Field(Name = "dependent_git_ids", StorageType = EumStorageType.Array)]
	DependentGitIds []int
	// DockerHub模板
	DockerHub int
	// DockerfileTpl模板
	DockerfileTpl int
	// K8SDeployment模板
	K8STplDeployment int
	// K8SIngress模板
	K8STplIngress int
	// K8SService模板
	K8STplService int
	// K8SConfig模板
	K8STplConfig int
	// K8S模板自定义变量(K1=V1,K2=V2)
	K8STplVariable string
	// 项目路径
	Path string
	// 镜像版本
	DockerVer string
	// 集群版本
	//[Field(Name = "cluster_ver", StorageType = EumStorageType.Json)]
	ClusterVer map[int]domain.ClusterVerVO
	// 构建方式
	BuildType eumBuildType.Enum
	// Shell脚本
	ShellScript string
	// K8S负载类型
	K8SControllersType eumK8SControllers.Enum
}
