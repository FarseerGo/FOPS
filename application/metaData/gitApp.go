package metaData

import (
	"fops/domain/building/build"
	"fops/domain/building/device"
	"fops/domain/metaData/git"
	"fs/core/container"
	"fs/mapper"
	"time"
)

type GitApp struct {
	repository git.Repository
	device     device.IGitDevice
}

func NewGitApp() *GitApp {
	return &GitApp{
		repository: container.Resolve[git.Repository](),
		device:     container.Resolve[device.IGitDevice](),
	}
}

// ToList Git列表
func (app *GitApp) ToList() []GitDto {
	lst := app.repository.ToList()
	return mapper.Array[GitDto](lst)
}

// Add 添加GIT
func (app *GitApp) Add(dto GitDto) {
	do := mapper.Single[git.DomainObject](dto)
	app.repository.Add(do)
}

// UpdateForTime 修改GIT的拉取时间
func (app *GitApp) UpdateForTime(gitId int, time time.Time) {
	app.repository.UpdateForTime(gitId, time)
}

// Update 修改GIT
func (app *GitApp) Update(dto GitDto) {
	do := mapper.Single[git.DomainObject](dto)
	app.repository.Update(dto.Id, do)
}

// ToInfo Git信息
func (app *GitApp) ToInfo(id int) GitDto {
	do := app.repository.ToInfo(id)
	return mapper.Single[GitDto](do)
}

// Delete 删除GIT
func (app *GitApp) Delete(id int) {
	app.repository.Delete(id)
}

// Count Git数量
func (app *GitApp) Count() int {
	return app.repository.Count()
}

// Clear 清除仓库
func (app *GitApp) Clear(gitId int, progress chan string) bool {
	if progress == nil {
		progress = make(chan string, 10)
	}

	var do = app.repository.ToInfo(gitId)
	return app.device.Clear(do.Hub, progress)
}

// CloneOrPull 根据判断是否存在Git目录，来决定返回Clone or pull
func (app *GitApp) CloneOrPull(dto GitDto, lstLog []string) bool {
	progress := make(chan string, 10)
	for _, log := range lstLog {
		progress <- log
	}
	vo := mapper.Single[build.GitVO](dto)
	return app.device.CloneOrPull(vo, progress)
}
