package devicer

import "fops/domain/building/build/vo"

type ICopyToDistDevice interface {
	Copy(lstGit []vo.GitVO, env vo.EnvVO, progress chan string)
}
