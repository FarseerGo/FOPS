package runLog

import (
	"fops/domain/_/eumLogLevel"
	"fops/domain/fss"
	"fs/core"
	"fs/core/container"
)

type app struct {
	repository fss.Repository
}

func NewApp() *app {
	return &app{
		repository: container.Resolve[fss.Repository](),
	}
}

// GetList 获取日志
func (app *app) GetList(jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[Dto] {
	return app.repository.GetRunLogList(jobName, logLevel, pageSize, pageIndex)
}
