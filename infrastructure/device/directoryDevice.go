package device

import (
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/directory"
	"os"
)

func init() {
	_ = container.Register(func() devicer.IDirectoryDevice { return &directoryDevice{} })
}

type directoryDevice struct {
}

// Check 检查目录
func (directoryDevice) Check(progress chan string) {
	progress <- "---------------------------------------------------------"
	progress <- "前置检查。"
	// 先删除之前编译的目标文件
	progress <- "先删除之前的目标文件。"

	directory.ClearFile(vo.DistRoot)

	// 自动创建目录
	progress <- "自动创建目录。"

	if !directory.IsExists(vo.FopsRoot) {
		os.MkdirAll(vo.FopsRoot, 0766)
	}
	if !directory.IsExists(vo.NpmModulesRoot) {
		os.MkdirAll(vo.NpmModulesRoot, 0766)
	}
	if !directory.IsExists(vo.DistRoot) {
		os.MkdirAll(vo.DistRoot, 0766)
	}
	if !directory.IsExists(vo.KubeRoot) {
		os.MkdirAll(vo.KubeRoot, 0766)
	}
	if !directory.IsExists(vo.ShellRoot) {
		os.MkdirAll(vo.ShellRoot, 0766)
	}
	if !directory.IsExists(vo.GitRoot) {
		os.MkdirAll(vo.GitRoot, 0766)
	}
	progress <- "前置检查通过。"
}
