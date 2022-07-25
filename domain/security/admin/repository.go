package admin

type Repository interface {
	// IsExists 管理员是否存在
	IsExists(adminName string) bool
	// IsExistsWithoutSelf 管理员是否存在
	IsExistsWithoutSelf(adminName string, adminId int) bool
	// Add 添加管理员
	Add(do DomainObject) int
	// Update 修改管理员
	Update(id int, do DomainObject)
	// ToList Admin列表
	ToList() []DomainObject
	// ToInfo Admin信息
	ToInfo(id int) DomainObject
	// ToInfoByUsername Admin信息
	ToInfoByUsername(username string, pwd string) DomainObject
	// Count Admin数量
	Count() int64
	// Delete 删除管理员
	Delete(id int)
}
