package device

import "fops/domain/build/build"

type IDockerDevice interface {
	// GetDockerHub 取得dockerHub
	GetDockerHub(dockerHubAddress string) string
	// GetDockerImage 生成镜像名称
	GetDockerImage(dockerHubAddress string, projectName string, buildNumber int) string
	// Login 登陆容器仓库
	Login(dockerHub string, loginName string, loginPwd string, progress chan string, env build.EnvVO) bool
	// ExistsDockerfile dockerfile文件是否存在
	ExistsDockerfile(dockerfilePath string) bool
	// CreateDockerfile 生成Dockerfile文件
	// projectName dockerfile文件地址
	// dockerfileContent 文件内容
	CreateDockerfile(projectName string, dockerfileContent string)
	// Build 容器构建
	Build(env build.EnvVO, progress chan string) bool
	// Push 上传镜像
	Push(env build.EnvVO, progress chan string) bool
}
