package domain

import "time"

// ClusterVerVO 集群镜像版本及部署时间
type ClusterVerVO struct {
	// 集群ID
	ClusterId int
	// 集群镜像版本
	DockerVer string
	// 上次部署成功时间
	DeploySuccessAt time.Time
	// 上次部署成功的构建ID
	BuildSuccessId int
	// 上次部署失败时间
	DeployFailAt time.Time
	// 上次部署失败的构建ID
	BuildFailId int
}
