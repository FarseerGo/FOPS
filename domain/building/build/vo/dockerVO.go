package vo

type DockerVO struct {
	Id                int    // 主键
	Name              string // 仓库名称
	Hub               string // 托管地址
	UserName          string // 账户名称
	UserPwd           string // 账户密码
	DockerfileContent string // Dockerfile内容
}

func NewDocker(docker DockerVO, dockerfileContent string) DockerVO {
	return DockerVO{
		Id:                docker.Id,
		Name:              docker.Name,
		Hub:               docker.Hub,
		UserName:          docker.UserName,
		UserPwd:           docker.UserPwd,
		DockerfileContent: dockerfileContent,
	}
}
