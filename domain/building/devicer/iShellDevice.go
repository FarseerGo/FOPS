package devicer

import (
	"context"
	"fops/domain/building/build/vo"
)

type IShellDevice interface {
	// ExecShell 执行Shell脚本
	ExecShell(env vo.EnvVO, shellScript string, progress chan string, ctx context.Context) bool
}
