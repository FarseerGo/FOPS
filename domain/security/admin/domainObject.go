package admin

import (
	"github.com/farseernet/farseer.go/utils/encrypt"
	"time"
)

// DomainObject 管理员
type DomainObject struct {
	Id          int       // 主键
	UserName    string    // 管理员名称
	UserPwd     string    // 管理员密码
	IsEnable    bool      // 是否启用
	LastLoginAt time.Time // 上次登陆时间
	LastLoginIp string    // 上次登陆IP
	CreateAt    time.Time // 创建时间
	CreateUser  string    // 创建人
	CreateId    string    // 创建人ID
}

func New() DomainObject {
	return DomainObject{
		CreateAt:    time.Now(),
		LastLoginIp: "",
	}
}

// EncryptPwd 对密码加密
func (admin *DomainObject) EncryptPwd(pwd string) {
	if pwd == "" {
		if len(pwd) < 6 {
			panic("管理员密码长度不能小于6")
		}
		admin.UserPwd = encrypt.Md5(pwd)
	}
}

// SetLoginIp 修改登陆时间信息
func (admin *DomainObject) SetLoginIp(ip string) {
	admin.LastLoginIp = ip
	admin.LastLoginAt = time.Now()
}
