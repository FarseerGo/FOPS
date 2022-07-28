package devicer

import (
	"context"
	"fops/domain/_/eumK8SControllers"
	"fops/domain/building/build/vo"
)

type IKubectlDevice interface {
	// GetConfigFile 获取存储k8s Config的路径
	GetConfigFile(clusterName string) string
	// CreateConfigFile 创建用于K8S远程管理的配置文件
	CreateConfigFile(clusterName string, clusterConfig string) string
	// SetYaml 生成yaml文件，并执行kubectl apply命令
	SetYaml(clusterName string, projectName string, yamlContent string, progress chan string) bool
	// SetImages 更新k8s的镜像版本
	SetImages(cluster vo.ClusterVO, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string, ctx context.Context) bool
	// SetImagesByClusterName 更新k8s的镜像版本
	SetImagesByClusterName(clusterName string, clusterConfig string, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string) bool
}
