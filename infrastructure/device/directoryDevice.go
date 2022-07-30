package device

import (
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/file"
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

	file.ClearFile(vo.DistRoot)

	// 自动创建目录
	progress <- "自动创建目录。"

	if !file.IsExists(vo.FopsRoot) {
		file.CreateDir766(vo.FopsRoot)
	}
	if !file.IsExists(vo.NpmModulesRoot) {
		file.CreateDir766(vo.NpmModulesRoot)
	}
	if !file.IsExists(vo.DistRoot) {
		file.CreateDir766(vo.DistRoot)
	}
	if !file.IsExists(vo.KubeRoot) {
		file.CreateDir766(vo.KubeRoot)
	}
	if !file.IsExists(vo.ShellRoot) {
		file.CreateDir766(vo.ShellRoot)
	}
	if !file.IsExists(vo.GitRoot) {
		file.CreateDir766(vo.GitRoot)
	}
	progress <- "前置检查通过。"
}
