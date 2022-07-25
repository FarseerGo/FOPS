package yamlTpl

type Repository interface {
	// ToList Yaml模板列表
	ToList() []DomainObject
	// ToInfo Yaml模板信息
	ToInfo(id int) DomainObject
	// Count Yaml模板数量
	Count() int64
	// Add 添加Yaml模板
	Add(do DomainObject) int
	// Update 修改Yaml模板
	Update(id int, do DomainObject)
	// Delete 删除Yaml模板
	Delete(id int)
}
