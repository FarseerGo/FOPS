package yamlTplAgent

import "fops/domain/_/eumK8SKind"

type PO struct {
	Id int `gorm:"primaryKey"`
	// k8s类型
	K8SKindType eumK8SKind.Enum
	// 模板名称
	Name string
	// 模板内容
	Template string
}
