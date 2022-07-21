package appLog

import (
	"fops/domain/appLog/containerLog"
	"fs/core/container"
	"fs/linq"
	"fs/mapper"
)

type containerLogApp struct {
	repository containerLog.Repository
}

func NewContainerLogApp() *containerLogApp {
	return &containerLogApp{repository: container.Resolve[containerLog.Repository]()}
}

// ToList 读取前500条日志
func (app *containerLogApp) ToList() []ContainerLogDto {
	lstDo := app.repository.ToList(100)
	lstDto := mapper.Array[ContainerLogDto](lstDo)
	return linq.FromOrder[ContainerLogDto, int64](lstDto).OrderBy(func(item ContainerLogDto) int64 {
		return item.CreateAt.UnixMicro()
	})
}
