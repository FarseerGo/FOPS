package device

import (
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
)

func init() {
	_ = container.Register(func() devicer.IShellDevice { return &shellDevice{} })
}

type shellDevice struct {
}

func (shellDevice) ExecShell(env vo.EnvVO, shellScript string, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}
