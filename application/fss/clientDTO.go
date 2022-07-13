package fss

import "time"

type ClientDTO struct {
	// 客户端ID
	Id int64
	// 客户端IP
	ClientIp string
	// 客户端名称
	ClientName string
	// 客户端能执行的任务
	Jobs []string
	// 活动时间
	ActivateAt time.Time
}
