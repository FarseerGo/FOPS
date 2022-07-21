package build

import (
	"fmt"
	"strconv"
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
	// Fops根目录
	FopsRoot string
	// kubectlConfig配置
	KubeRoot string
	// NpmModules
	NpmModulesRoot string
	// 编译保存的根目录
	DistRoot string
	// GIT根目录
	GitRoot string
	// Dockerfile文件地址
	DockerfilePath string
	// DockerIgnore文件地址
	DockerIgnorePath string
	// 生成Shell脚本的存放路径
	ShellRoot string
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

func NewEnvVO() EnvVO {
	return EnvVO{
		// Fops根目录
		FopsRoot: "/var/lib/fops/",
		// kubectlConfig配置
		KubeRoot: "/var/lib/fops/kube/",
		// NpmModules
		NpmModulesRoot: "/var/lib/fops/npm",
		// 编译保存的根目录
		DistRoot: "/var/lib/fops/dist/",
		// GIT根目录
		GitRoot: "/var/lib/fops/git/",
		// Dockerfile文件地址
		DockerfilePath: "/var/lib/fops/dist/Dockerfile",
		// DockerIgnore文件地址
		DockerIgnorePath: "/var/lib/fops/dist/.dockerignore",
		// 生成Shell脚本的存放路径
		ShellRoot: "/var/lib/fops/shell/",
	}
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
		"FopsRoot":           env.FopsRoot,
		"NpmModulesRoot":     env.NpmModulesRoot,
		"DistRoot":           env.DistRoot,
		"KubeRoot":           env.KubeRoot,
		"Git_Root":           env.GitRoot,
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
