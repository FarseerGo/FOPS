package git

import (
	"fops/infrastructure/repository/context"
	"time"
)

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList Git列表
func (agent) ToList() []PO {
	return context.NewContext().Git.Desc("Id").ToList()
}

// ToListByIds Git列表
func (agent) ToListByIds(ids []int) []PO {
	return context.NewContext().Git.Where("Id in ?", ids).ToList()
}

// ToInfo Git信息
func (agent) ToInfo(id int) PO {
	return context.NewContext().Git.Where("Id = ?", id).ToEntity()
}

// Count Git数量
func (agent) Count() int64 {
	return context.NewContext().Git.Count()
}

// Add 添加GIT
func (agent) Add(po PO) int {
	context.NewContext().Git.Insert(&po)
	return po.Id
}

// Update 修改GIT
func (agent) Update(id int, po PO) {
	context.NewContext().Git.Where("Id = ?", id).Update(po)
}

// UpdatePullAt 修改GIT的拉取时间
func (agent) UpdatePullAt(id int, pullAt time.Time) {
	context.NewContext().Git.Where("Id = ?", id).UpdateValue("PullAt", pullAt)
}

// Delete 删除GIT
func (agent) Delete(id int) {
	context.NewContext().Git.Where("Id = ?", id).Delete()
}
