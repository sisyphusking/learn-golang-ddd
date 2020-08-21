package event

import (
	"fmt"
	"leave/domain/leave/event"
)

type Publisher struct {
}

func (p *Publisher) Publish(event *event.LeaveEvent) error {

	fmt.Sprintf("event: %v", event)
	return nil
}
