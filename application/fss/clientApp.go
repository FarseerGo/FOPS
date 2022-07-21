package fss

import (
	"fops/domain/fss"
	"fs/core/container"
	"fs/linq"
)

type ClientApp struct {
	repository fss.Repository
}

func NewClientApp() *ClientApp {
	return &ClientApp{
		repository: container.Resolve[fss.Repository](),
	}
}

// GetCount 取出全局客户端数量
func (app *ClientApp) GetCount() int64 {
	return app.repository.GetClientCount()
}

// GetList 取出全局客户端列表
func (app *ClientApp) GetList() []ClientDto {
	lst := app.repository.GetClientList()
	return linq.FromOrder[ClientDto, int64](lst).OrderByDescending(func(item ClientDto) int64 {
		return item.ActivateAt.UnixMicro()
	})
}
