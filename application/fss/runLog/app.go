package runLog

import (
	"fops/domain/fss"
	"github.com/farseernet/farseer.go/core"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/core/eumLogLevel"
)

type app struct {
	repository fss.IFssDevice
	fssServer  string
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[fss.IFssDevice](),
	}
}

// GetList 获取日志
func (app *app) GetList(jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[Dto] {
	return app.repository.GetRunLogList(app.fssServer, jobName, logLevel, pageSize, pageIndex)
}
