package vo

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
