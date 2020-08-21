package po

import (
	"leave/domain/leave/event"
	"time"
)

type LeaveEventPO struct {
	Id             string
	LeaveEventType event.LeaveEventType
	Timestamp      time.Time
	Source         string
	Data           string
}
