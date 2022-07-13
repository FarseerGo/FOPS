package pod

type Repository interface {

	// ToListByGroupId 获取当前POD列表
	ToListByGroupId(groupId int) []DomainObject
	// ToList 获取当前POD列表
	ToList() []DomainObject
	// Update 更新模板ID
	Update(pod DomainObject)
}
