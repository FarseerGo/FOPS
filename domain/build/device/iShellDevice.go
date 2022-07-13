package device

import "fops/domain/build/build"

type IShellDevice interface {
	// ExecShell 执行Shell脚本
	ExecShell(env build.EnvVO, shellScript string, progress chan string) bool
}
