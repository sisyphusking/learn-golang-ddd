package event

import (
	"encoding/json"
	"leave/domain/leave/entity"
	"time"

	uuid "github.com/satori/go.uuid"
)

type LeaveEvent struct {
	LeaveEventType LeaveEventType
	Id             string
	Timestamp      time.Time
	Source         string
	Data           string
}

func NewLeaveEvent(leaveEventType LeaveEventType, leave *entity.Leave) *LeaveEvent {
	b, _ := json.Marshal(leave)
	id := uuid.NewV4().String()
	source := "leave event"
	return &LeaveEvent{leaveEventType, id, time.Now(), source, string(b)}
}
