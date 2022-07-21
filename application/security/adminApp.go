package security

import (
	"fops/domain/security/admin"
	"fs/core/container"
	"fs/mapper"
)

type AdminApp struct {
	repository admin.Repository
}

func NewAdminApp() *AdminApp {
	return &AdminApp{
		repository: container.Resolve[admin.Repository](),
	}
}

// Add 添加管理员
func (app *AdminApp) Add(dto AdminDto) {
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
func (app *AdminApp) Update(dto AdminDto) {
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
func (app *AdminApp) ToList() []AdminDto {
	lst := app.repository.ToList()
	return mapper.Array[AdminDto](lst)
}

// ToInfo Admin信息
func (app *AdminApp) ToInfo(id int) AdminDto {
	do := app.repository.ToInfo(id)
	dto := mapper.Single[AdminDto](do)
	dto.UserPwd = ""
	return dto
}
