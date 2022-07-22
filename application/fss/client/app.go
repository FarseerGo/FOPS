package client

import (
	"fops/domain/fss"
	"fs/core/container"
	"fs/linq"
)

type app struct {
	repository fss.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[fss.Repository](),
	}
}

// GetCount 取出全局客户端数量
func (app *app) GetCount() int64 {
	return app.repository.GetClientCount()
}

// GetList 取出全局客户端列表
func (app *app) GetList() []Dto {
	lst := app.repository.GetClientList()
	return linq.FromOrder[Dto, int64](lst).OrderByDescending(func(item Dto) int64 {
		return item.ActivateAt.UnixMicro()
	})
}
