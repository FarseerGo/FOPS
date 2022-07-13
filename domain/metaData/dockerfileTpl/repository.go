package dockerfileTpl

type Repository interface {
	// ToList Dockerfile模板列表
	ToList() []DomainObject
	// ToInfo Dockerfile模板信息
	ToInfo(id int) DomainObject
	// Count Dockerfile模板数量
	Count() int
	// Add 添加Dockerfile模板
	Add(dockerfileTpl DomainObject)
	// Update 修改Dockerfile模板
	Update(id int, dockerfileTpl DomainObject)
	// Delete 删除Dockerfile模板
	Delete(id int)
}
