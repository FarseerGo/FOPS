package queue

import "time"

type BuildLogVO struct {
	BuildId int
	Log     string
	LogAt   time.Time
}
