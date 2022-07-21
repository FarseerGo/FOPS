package fss

import (
	"fops/domain/_/eumLogLevel"
	"fops/domain/fss"
	"fs/core"
	"fs/core/container"
)

type runLogApp struct {
	repository fss.Repository
}

func NewRunLogApp() *runLogApp {
	return &runLogApp{
		repository: container.Resolve[fss.Repository](),
	}
}

// GetList 获取日志
func (app *runLogApp) GetList(jobName string, logLevel eumLogLevel.Enum, pageSize int, pageIndex int) core.PageList[RunLogDTO] {
	return app.repository.GetRunLogList(jobName, logLevel, pageSize, pageIndex)
}
