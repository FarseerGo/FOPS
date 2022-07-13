package device

import "fops/domain/build/build"

type ICopyToDistDevice interface {
	Copy(lstGit []build.GitVO, env build.EnvVO, progress chan string)
}
