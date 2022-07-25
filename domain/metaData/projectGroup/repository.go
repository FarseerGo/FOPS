package projectGroup

// Repository 项目管理仓储
type Repository interface {
	// ToList 项目组列表
	ToList() []DomainObject
	// ToInfo 项目组信息
	ToInfo(id int) *DomainObject
	// Count 项目组数量
	Count() int64
	// Add 添加项目组
	Add(vo DomainObject) int
	// Update 修改项目组
	Update(id int, projectGroup DomainObject)
	// Delete 删除项目组
	Delete(id int)
}
