package device

type ICopyToDistDevice interface {
	Copy(lstGit []vo.GitVO, env vo.EnvVO, progress chan string)
}
