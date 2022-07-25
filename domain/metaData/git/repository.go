package git

import "time"

type Repository interface {
	// ToList Git列表
	ToList() []DomainObject
	// ToListByIds Git列表
	ToListByIds(ids []int) []DomainObject
	// ToListByMaster Git列表
	ToListByMaster(masterGitId int, ids []int) []DomainObject
	// ToInfo Git信息
	ToInfo(id int) DomainObject
	// Count Git数量
	Count() int64
	// Add 添加GIT
	Add(do DomainObject) int
	// Update 修改GIT
	Update(id int, do DomainObject)
	// UpdateForTime 修改GIT的拉取时间
	UpdateForTime(id int, pullAt time.Time)
	// Delete 删除GIT
	Delete(id int)
}
