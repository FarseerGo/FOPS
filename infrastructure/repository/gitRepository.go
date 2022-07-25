package repository

import (
	"fops/domain/metaData/git"
	"fops/infrastructure/repository/agent/gitAgent"
	"fs/core/container"
	"fs/mapper"
	"time"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() git.Repository { return &gitRepository{} })
}

type gitRepository struct {
}

// ToList Git列表
func (repository gitRepository) ToList() []git.DomainObject {
	lst := gitAgent.ToList()
	return mapper.Array[git.DomainObject](lst)
}

// ToListByIds Git列表
func (repository gitRepository) ToListByIds(ids []int) []git.DomainObject {
	lst := gitAgent.ToListByIds(ids)
	return mapper.Array[git.DomainObject](lst)
}

// ToListByMaster Git列表
func (repository gitRepository) ToListByMaster(masterGitId int, ids []int) []git.DomainObject {
	return repository.ToListByIds(append(ids, masterGitId))
}

// ToInfo Git信息
func (repository gitRepository) ToInfo(id int) git.DomainObject {
	po := gitAgent.ToInfo(id)
	return mapper.Single[git.DomainObject](po)
}

// Count Git数量
func (repository gitRepository) Count() int64 {
	return gitAgent.Count()
}

// Add 添加GIT
func (repository gitRepository) Add(do git.DomainObject) int {
	po := mapper.Single[gitAgent.PO](do)
	return gitAgent.Add(po)
}

// Update 修改GIT
func (repository gitRepository) Update(id int, do git.DomainObject) {
	po := mapper.Single[gitAgent.PO](do)
	gitAgent.Update(id, po)
}

// UpdateForTime 修改GIT的拉取时间
func (repository gitRepository) UpdateForTime(id int, pullAt time.Time) {
	gitAgent.UpdatePullAt(id, pullAt)
}

// Delete 删除GIT
func (repository gitRepository) Delete(id int) {
	gitAgent.Delete(id)
}
