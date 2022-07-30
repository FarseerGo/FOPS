package device

import (
	"context"
	"fops/domain/building/build/vo"
	"fops/domain/building/devicer"
	"github.com/farseernet/farseer.go/core/container"
	"github.com/farseernet/farseer.go/utils/exec"
	"github.com/farseernet/farseer.go/utils/file"
	"strconv"
)

func init() {
	_ = container.Register(func() devicer.IShellDevice { return &shellDevice{} })
}

type shellDevice struct {
}

func (shellDevice) ExecShell(env vo.EnvVO, shellScript string, progress chan string, ctx context.Context) bool {
	// 每次执行时，需要生成shell脚本
	path := vo.ShellRoot + "fops_" + strconv.Itoa(env.BuildId) + ".sh"
	file.WriteString(path, shellScript)

	// 执行脚本
	var exitCode = exec.RunShellContext("/bin/sh -xe {path}", progress, env.ToMap(), vo.DistRoot, ctx)
	if exitCode == 0 {
		progress <- "执行脚本完成。"
	} else {

		progress <- "执行脚本出错了。"
	}
	return exitCode == 0
}
