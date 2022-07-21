package build

// GitVO Git仓库
type GitVO struct {
	Id          int    // 主键
	Name        string // Git名称
	Hub         string // 托管地址
	Branch      string // Git分支
	UserName    string // 账户名称
	UserPwd     string // 账户密码
	ProjectPath string // 项目位置
	IsMaster    bool   // 是否主仓库
}
