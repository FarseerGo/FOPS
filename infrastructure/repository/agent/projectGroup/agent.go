package projectGroup

import "fops/infrastructure/repository/context"

type agent struct {
}

func NewAgent() agent { return agent{} }

// ToList 项目组列表
func (agent) ToList() []PO {
	return context.NewContext().ProjectGroup.Desc("Id").ToList()
}

// ToInfo 项目组信息
func (agent) ToInfo(id int) PO {
	return context.NewContext().ProjectGroup.Where("Id = ?", id).ToEntity()
}

// Count 项目组数量
func (agent) Count() int64 {
	return context.NewContext().ProjectGroup.Count()
}

// Add 添加项目组
func (agent) Add(po PO) int {
	context.NewContext().ProjectGroup.Insert(&po)
	return po.Id
}

// Update 修改项目组
func (agent) Update(id int, po PO) {
	context.NewContext().ProjectGroup.Where("Id = ?", id).Update(po)
}

// Delete 删除项目组
func (agent) Delete(id int) {
	context.NewContext().ProjectGroup.Where("Id = ?", id).Delete()
}
