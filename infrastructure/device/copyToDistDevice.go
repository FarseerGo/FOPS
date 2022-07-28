package device

import (
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/directory"
	"strings"
)

func init() {
	_ = container.Register(func() devicer.ICopyToDistDevice { return &copyToDistDevice{} })
}

type copyToDistDevice struct {
}

// Copy 将项目的源代码，复制到目标目录中。用于后续的打包
func (copyToDistDevice) Copy(lstGit []vo.GitVO, env vo.EnvVO, progress chan string) {
	progress <- "---------------------------------------------------------"
	for _, git := range lstGit {
		gitArr := strings.Split(git.ProjectPath, "/")
		progress <- "源文件" + git.ProjectPath + " 复制到 " + env.DistRoot + gitArr[len(gitArr)-2]
		directory.CopyFolder(git.ProjectPath, env.DistRoot+gitArr[len(gitArr)-2])
	}
}
