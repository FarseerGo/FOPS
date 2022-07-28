package kubectlSetImage

import (
	"fops/domain/building/devicer"
	"fops/domain/k8s/cluster"
	"fops/domain/metaData/dockerHub"
	"fops/domain/metaData/project"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/parse"
)

type app struct {
	clusterRepository   cluster.Repository
	projectRepository   project.Repository
	dockerHubRepository dockerHub.Repository
	dockerDevice        devicer.IDockerDevice
	kubectlDevice       devicer.IKubectlDevice
}

func NewApp() *app {
	return &app{
		clusterRepository:   container.Resolve[cluster.Repository](),
		projectRepository:   container.Resolve[project.Repository](),
		dockerHubRepository: container.Resolve[dockerHub.Repository](),
		dockerDevice:        container.Resolve[devicer.IDockerDevice](),
		kubectlDevice:       container.Resolve[devicer.IKubectlDevice](),
	}
}

func (app *app) SyncImages(clusterId int, projectId int, progress chan string) bool {
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
