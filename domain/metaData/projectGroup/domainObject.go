package projectGroup

type DomainObject struct {
	// 主键
	Id int
	// 集群ID
	ClusterIds []int
	// 项目组名称
	Name string
}

func New() DomainObject {
	return DomainObject{}
}
