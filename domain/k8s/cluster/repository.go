package cluster

type Repository interface {
	// ToList 集群列表
	ToList() []DomainObject
	// ToInfo 集群信息
	ToInfo(id int) DomainObject
	// Count 集群数量
	Count() int
	// Add 添加集群
	Add(cluster DomainObject) int
	// Update 修改集群
	Update(id int, cluster DomainObject)
	// Delete 删除集群
	Delete(id int)
}
