package device

import (
	"fops/domain/building/devicer"
)

func init() {
	_ = container.Register(func() devicer.ILogWriteDevice { return &logWriteDevice{} })
}

type logWriteDevice struct {
}

func (logWriteDevice) View(buildId int) []string {
	//TODO implement me
	panic("implement me")
}

func (logWriteDevice) CreateProgress(buildId int) chan string {
	//TODO implement me
	panic("implement me")
}
