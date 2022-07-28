package device

import (
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
)

func init() {
	_ = container.Register(func() devicer.IDotnetDevice { return &dotnetDevice{} })
}

type dotnetDevice struct {
}

func (dotnetDevice) GetReleasePath(projectName string) string {
	//TODO implement me
	panic("implement me")
}

func (dotnetDevice) CheckExistsSource(env vo.EnvVO, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}

func (dotnetDevice) Publish(savePath string, source string, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}

func (dotnetDevice) PublishByEnv(env vo.EnvVO, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}

func (dotnetDevice) GetSourceDirRoot(github string, projectPath string) string {
	//TODO implement me
	panic("implement me")
}
