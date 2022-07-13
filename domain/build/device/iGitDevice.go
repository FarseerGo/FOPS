package device

import "fops/domain/build/build"

type IGitDevice interface {
	// GetGitPath 获取项目GIT源代码存的位置/var/lib/fops/git/{gitName}/
	GetGitPath(gitHub string) string
	// RememberPassword 记住密码
	RememberPassword(env build.EnvVO, progress chan string)
	// ExistsGitProject 项目GIT是否存在
	ExistsGitProject(gitPath string) bool
	// Clear 消除仓库
	Clear(gitHub string, progress chan string) bool
	// GetName gitName，项目文件夹名称
	GetName(gitHub string) string
	// CloneOrPull 根据判断是否存在Git目录，来决定返回Clone or pull
	CloneOrPull(git build.GitVO, progress chan string) bool
	// CloneOrPullAndDependent 拉取主仓库及依赖仓库
	CloneOrPullAndDependent(lstGit []build.GitVO, progress chan string) bool
}
