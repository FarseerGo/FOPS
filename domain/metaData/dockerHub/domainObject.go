package dockerHub

type DomainObject struct {
	// 主键
	Id int
	// 仓库名称
	Name string
	// 托管地址
	Hub string
	// 账户名称
	UserName string
	// 账户密码
	UserPwd string
}

func New() DomainObject {
	return DomainObject{}
}
