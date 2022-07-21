package k8s

import (
	"fops/domain/building/device"
	"fops/domain/k8s/cluster"
	"fops/domain/metaData/dockerHub"
	"fops/domain/metaData/project"
	"fs/core/container"
	"fs/utils/parse"
)

type KubectlSetImageApp struct {
	clusterRepository   cluster.Repository
	projectRepository   project.Repository
	dockerHubRepository dockerHub.Repository
	dockerDevice        device.IDockerDevice
	kubectlDevice       device.IKubectlDevice
}

func NewKubectlSetImageApp() *KubectlSetImageApp {
	return &KubectlSetImageApp{
		clusterRepository:   container.Resolve[cluster.Repository](),
		projectRepository:   container.Resolve[project.Repository](),
		dockerHubRepository: container.Resolve[dockerHub.Repository](),
		dockerDevice:        container.Resolve[device.IDockerDevice](),
		kubectlDevice:       container.Resolve[device.IKubectlDevice](),
	}
}

func (app *KubectlSetImageApp) SyncImages(clusterId int, projectId int, progress chan string) bool {
	clusterDo := app.clusterRepository.ToInfo(clusterId)
	projectDo := app.projectRepository.ToInfo(projectId)
	dockerDo := app.dockerHubRepository.ToInfo(projectDo.DockerHub)

	// 组装镜像版本
	dockerVer := parse.Convert(projectDo.DockerVer, 0)
	dockerImage := app.dockerDevice.GetDockerImage(dockerDo.Hub, projectDo.Name, dockerVer)

	// 更新镜像
	result := app.kubectlDevice.SetImagesByClusterName(clusterDo.Name, clusterDo.Config, projectDo.Name, dockerImage, projectDo.K8SControllersType, progress)
	if result {
		progress <- "更新镜像版本完成。"
	}
	return result
}
