package device

import "fops/domain/building/build"

type ICopyToDistDevice interface {
	Copy(lstGit []build.GitVO, env build.EnvVO, progress chan string)
}
