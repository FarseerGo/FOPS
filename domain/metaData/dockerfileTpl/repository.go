package dockerfileTpl

type Repository interface {
	// ToList Dockerfile模板列表
	ToList() []DomainObject
	// ToInfo Dockerfile模板信息
	ToInfo(id int) DomainObject
	// Count Dockerfile模板数量
	Count() int64
	// Add 添加Dockerfile模板
	Add(do DomainObject)
	// Update 修改Dockerfile模板
	Update(id int, do DomainObject)
	// Delete 删除Dockerfile模板
	Delete(id int)
}
