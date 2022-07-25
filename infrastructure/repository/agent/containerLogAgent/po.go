package containerLogAgent

type PO struct {
	Id int `gorm:"primaryKey"`
	// 应用名称
	//[Keyword]
	AppName string
	// 容器名称
	//[Keyword]
	ContainerName string
	// 镜像名称
	//[Keyword]
	ContainerImage string
	// 容器IP
	//[Keyword]
	ContainerIp string
	// 环境变量
	//[Flattened]
	ContainerEnv map[string]string
	// 节点名称
	//[Keyword]
	NodeName string
	// 节点IP
	//[Keyword]
	NodeIp string
	// 日志级别
	//[Keyword]
	LogLevel string
	// 日志内容
	//[Text]
	Content string
	// 日志时间
	//[Date]
	CreateAt int64
}
