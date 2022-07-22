package vo

import (
	"fops/domain/_/eumK8SControllers"
)

// ProjectVO 构建时的项目
type ProjectVO struct {
	Id                 int                    // 主键
	Name               string                 // 项目名称
	EntryPoint         string                 // 程序入口名称
	EntryPort          int                    // 程序启动端口
	Domain             string                 // 访问域名
	Path               string                 // 项目路径
	K8STplVariable     string                 // K8S模板自定义变量(K1=V1K2=V2)
	K8SControllersType eumK8SControllers.Enum // K8S负载类型
	ShellScript        string                 // Shell脚本
}
