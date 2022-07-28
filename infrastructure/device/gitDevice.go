package device

import (
	"context"
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
)

func init() {
	_ = container.Register(func() devicer.IGitDevice { return &gitDevice{} })
}

type gitDevice struct {
}

func (gitDevice) GetGitPath(gitHub string) string {
	//TODO implement me
	panic("implement me")
}

func (gitDevice) RememberPassword(env vo.EnvVO, progress chan string) {
	//TODO implement me
	panic("implement me")
}

func (gitDevice) ExistsGitProject(gitPath string) bool {
	//TODO implement me
	panic("implement me")
}

func (gitDevice) Clear(gitHub string, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}

func (gitDevice) GetName(gitHub string) string {
	//TODO implement me
	panic("implement me")
}

func (gitDevice) CloneOrPull(git vo.GitVO, progress chan string) bool {
	//TODO implement me
	panic("implement me")
}

func (gitDevice) CloneOrPullAndDependent(lstGit []vo.GitVO, progress chan string, ctx context.Context) bool {
	//TODO implement me
	panic("implement me")
}
