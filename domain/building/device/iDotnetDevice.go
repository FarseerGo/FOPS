package device

import "fops/domain/building/build/vo"

type IDotnetDevice interface {
	// GetReleasePath 获取编译保存的目录地址
	GetReleasePath(projectName string) string
	// CheckExistsSource 检查项目文件是否存在
	CheckExistsSource(env vo.EnvVO, progress chan string) bool
	// Publish 编译.net core
	Publish(savePath string, source string, progress chan string) bool
	// PublishByEnv 编译.net core
	PublishByEnv(env vo.EnvVO, progress chan string) bool
	// GetSourceDirRoot 获取项目源地址
	GetSourceDirRoot(github string, projectPath string) string
}
