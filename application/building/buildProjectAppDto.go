package building

import (
	domain "fops/domain/_"
	"fops/domain/_/eumK8SControllers"
)

type BuildProjectAppDto struct {
	// 主键
	Id int
	// 项目名称
	Name string
	// Dockerfile
	DockerfileName string
	// DockerfileTpl模板
	DockerfileTpl int
	// 镜像版本
	DockerVer string
	// K8S负载类型
	K8sControllersType eumK8SControllers.Enum
	// 集群版本
	ClusterVer map[int]domain.ClusterVerVO
}
