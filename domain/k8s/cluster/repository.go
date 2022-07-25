package cluster

type Repository interface {
	// ToList 集群列表
	ToList() []DomainObject
	// ToInfo 集群信息
	ToInfo(id int) DomainObject
	// Count 集群数量
	Count() int64
	// Add 添加集群
	Add(do DomainObject) int
	// Update 修改集群
	Update(id int, do DomainObject)
	// Delete 删除集群
	Delete(id int)
}
