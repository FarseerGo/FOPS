package dockerfileTpl

type DomainObject struct {
	// 主键
	Id int
	// 模板名称
	Name string
	// 模板内容
	Template string
}

func New() DomainObject {
	return DomainObject{}
}
