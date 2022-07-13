package pod

import "fops/domain/_/eumK8SKind"

type YamlTplVO struct {
	// 主键
	Id int
	// k8s类型
	K8SKindType eumK8SKind.Enum
	// 模板名称
	Name string
	// 模板内容
	Template string
}
