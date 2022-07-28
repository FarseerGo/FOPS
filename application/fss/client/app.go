package client

import (
	"fops/domain/fss"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/linq"
)

type app struct {
	repository fss.IFssDevice
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[fss.IFssDevice](),
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
