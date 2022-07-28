package repository

import (
	"fops/domain/metaData/git"
	"fops/infrastructure/repository/context"
	"fops/infrastructure/repository/model"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/data"
	"github.com/farseernet/farseer.go/mapper"

	"time"
)

func init() {
	// 注册项目组仓储
	_ = container.Register(func() git.Repository { return &gitRepository{data.Init[context.MysqlContext]("fops").Git} })
}

type gitRepository struct {
	data.TableSet[model.GitPO]
}

// ToList Git列表
func (repository gitRepository) ToList() []git.DomainObject {
	lst := repository.Desc("Id").ToList()
	return mapper.Array[git.DomainObject](lst)
}

// ToListByIds Git列表
func (repository gitRepository) ToListByIds(ids []int) []git.DomainObject {
	lst := repository.Where("Id in ?", ids).ToList()
	return mapper.Array[git.DomainObject](lst)
}

// ToListByMaster Git列表
func (repository gitRepository) ToListByMaster(masterGitId int, ids []int) []git.DomainObject {
	return repository.ToListByIds(append(ids, masterGitId))
}

// ToInfo Git信息
func (repository gitRepository) ToInfo(id int) git.DomainObject {
	po := repository.Where("Id = ?", id).ToEntity()
	return mapper.Single[git.DomainObject](po)
}

//// Count Git数量
//func (repository gitRepository) Count() int64 {
//	return repository.Count()
//}

// Add 添加GIT
func (repository gitRepository) Add(do git.DomainObject) int {
	po := mapper.Single[model.GitPO](do)
	repository.Insert(&po)
	return po.Id
}

// Update 修改GIT
func (repository gitRepository) Update(id int, do git.DomainObject) {
	po := mapper.Single[model.GitPO](do)
	repository.Where("Id = ?", id).Update(po)
}

// UpdateForTime 修改GIT的拉取时间
func (repository gitRepository) UpdateForTime(id int, pullAt time.Time) {
	repository.Where("Id = ?", id).UpdateValue("PullAt", pullAt)
}

// Delete 删除GIT
func (repository gitRepository) Delete(id int) {
	repository.Where("Id = ?", id).Delete()
}
