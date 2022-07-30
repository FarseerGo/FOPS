package vo

import (
	"fmt"
	"strconv"
)

const (
	// FopsRoot Fops根目录
	FopsRoot = "/var/lib/fops/"
	// KubeRoot kubectlConfig配置
	KubeRoot = "/var/lib/fops/kube/"
	// NpmModulesRoot NpmModules
	NpmModulesRoot = "/var/lib/fops/npm"
	// DistRoot 编译保存的根目录
	DistRoot = "/var/lib/fops/dist/"
	// GitRoot GIT根目录
	GitRoot = "/var/lib/fops/git/"
	// DockerfilePath Dockerfile文件地址
	DockerfilePath = "/var/lib/fops/dist/Dockerfile"
	// DockerIgnorePath DockerIgnore文件地址
	DockerIgnorePath = "/var/lib/fops/dist/.dockerignore"
	// ShellRoot 生成Shell脚本的存放路径
	ShellRoot = "/var/lib/fops/shell/"
)

// EnvVO 构建时的环境变量
type EnvVO struct {
	// 构建版本号
	BuildId int
	// 构建版本号
	BuildNumber int
	// 项目ID
	ProjectId int
	// 项目名称
	ProjectName string
	// 项目访问域名
	ProjectDomain string
	// 程序入口名
	ProjectEntryPoint string
	// 程序启动端口
	ProjectEntryPort int
	// Git仓库源代码根目录
	// /var/lib/fops/git/{gitName}/
	ProjectGitRoot string
	// Git仓库地址
	GitHub string
	// Git名称（项目的目录名称）
	GitName string
	// Docker仓库地址
	DockerHub string
	// Docker镜像
	DockerImage string
}

// Print 打印环境变量
func (env *EnvVO) Print(progress chan string) {
	// 打印环境变量
	progress <- "---------------------------------------------------------"
	progress <- "环境变量："

	for k, v := range env.ToMap() {
		progress <- fmt.Sprint(k, "=", v)
	}
}

// ToMap 转成字典
func (env *EnvVO) ToMap() map[string]string {
	return map[string]string{
		"FopsRoot":           FopsRoot,
		"NpmModulesRoot":     NpmModulesRoot,
		"DistRoot":           DistRoot,
		"KubeRoot":           KubeRoot,
		"Git_Root":           GitRoot,
		"Git_Hub":            env.GitHub,
		"Build_Number":       strconv.Itoa(env.BuildNumber),
		"Project_Name":       env.ProjectName,
		"Project_Domain":     env.ProjectDomain,
		"Project_GitRoot":    env.ProjectGitRoot,
		"Project_EntryPoint": env.ProjectEntryPoint,
		"Project_EntryPort":  strconv.Itoa(env.ProjectEntryPort),
		"Docker_Hub":         env.DockerHub,
		"Docker_Image":       env.DockerImage,
		"Git_Name":           env.GitName,
	}
}
