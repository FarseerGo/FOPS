package containerLog

import (
	"fops/domain/appLog/containerLog"
	"fs/core/container"
	"fs/linq"
	"fs/mapper"
)

type app struct {
	repository containerLog.Repository
}

func NewApp() *app {
	return &app{repository: container.Resolve[containerLog.Repository]()}
}

// ToList 读取前500条日志
func (app *app) ToList() []Dto {
	lstDo := app.repository.ToList(100)
	lstDto := mapper.Array[Dto](lstDo)
	return linq.FromOrder[Dto, int64](lstDto).OrderBy(func(item Dto) int64 {
		return item.CreateAt.UnixMicro()
	})
}