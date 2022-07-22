package admin

import (
	"fops/domain/security/admin"
	"fs/core/container"
	"fs/mapper"
)

type app struct {
	repository admin.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[admin.Repository](),
	}
}

// Add 添加管理员
func (app *app) Add(dto Dto) {
	if dto.UserName == "" || dto.UserPwd == "" {
		panic("管理员名称、登陆密码必须填写。")
	}

	do := mapper.Single[admin.DomainObject](dto)
	do.EncryptPwd(do.UserPwd)
	var isExists = app.repository.IsExists(do.UserName)
	if isExists {
		panic("管理员名称已存在")
	}
	app.repository.Add(do)
}

// Update 修改管理员
func (app *app) Update(dto Dto) {
	if dto.Id < 1 {
		panic("管理员不存在。")
	}
	if dto.UserName == "" {
		panic("管理员名称必须填写。")
	}

	do := mapper.Single[admin.DomainObject](dto)
	do.EncryptPwd(do.UserPwd)

	var isExists = app.repository.IsExistsWithoutSelf(do.UserName, do.Id)
	if isExists {
		panic("管理员名称已存在")
	}

	app.repository.Update(do.Id, do)
}

// ToList Admin列表
func (app *app) ToList() []Dto {
	lst := app.repository.ToList()
	return mapper.Array[Dto](lst)
}

// ToInfo Admin信息
func (app *app) ToInfo(id int) Dto {
	do := app.repository.ToInfo(id)
	dto := mapper.Single[Dto](do)
	dto.UserPwd = ""
	return dto
}
