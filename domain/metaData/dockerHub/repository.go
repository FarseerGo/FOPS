package dockerHub

type Repository interface {
	// ToList DockerHub列表
	ToList() []DomainObject
	// ToInfo DockerHub信息
	ToInfo(id int) DomainObject
	// Count DockerHub数量
	Count() int64
	// Add 添加仓库
	Add(do DomainObject)
	// Update 修改仓库
	Update(id int, do DomainObject)
	// Delete 删除仓库
	Delete(id int)
}
