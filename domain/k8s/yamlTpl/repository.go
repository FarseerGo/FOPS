package yamlTpl

type Repository interface {
	// ToList Yaml模板列表
	ToList() []DomainObject
	// ToInfo Yaml模板信息
	ToInfo(id int) DomainObject
	// Count Yaml模板数量
	Count() int
	// Add 添加Yaml模板
	Add(yamlTpl DomainObject) int
	// Update 修改Yaml模板
	Update(id int, yamlTpl DomainObject)
	// Delete 删除Yaml模板
	Delete(id int)
}
