package device

import (
	"context"
	"fops/domain/_/eumK8SControllers"
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
)

func init() {
	_ = container.Register(func() devicer.IKubectlDevice { return &kubectlDevice{} })
}

type kubectlDevice struct {
}

func (kubectlDevice) GetConfigFile(clusterName string) string {
	//TODO implement me
	panic("implement me")
}

func (kubectlDevice) CreateConfigFile(clusterName string, clusterConfig string) string {
	//TODO implement me
	panic("implement me")
}

func (kubectlDevice) SetYaml(clusterName string, projectName string, yamlContent string, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}

func (kubectlDevice) SetImages(cluster vo.ClusterVO, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string, ctx context.Context) bool {
	//TODO implement me
	panic("implement me")
}

func (kubectlDevice) SetImagesByClusterName(clusterName string, clusterConfig string, projectName string, dockerImages string, k8SControllersType eumK8SControllers.Enum, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}
