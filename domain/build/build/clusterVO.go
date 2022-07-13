package build

// ClusterVO 集群
type ClusterVO struct {
	Id     int    // 主键
	Name   string // 集群名称
	Config string // 本地kubectl配置地址内容
}
