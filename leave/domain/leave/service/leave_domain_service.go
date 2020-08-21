package service

import (
	"leave/domain/leave/entity"
	"leave/domain/leave/entity/approver"
	event2 "leave/domain/leave/event"
	"leave/domain/leave/repository/facade"
	"leave/infrastructure/util/event"
)

type LeaveDomainService struct {
	publisher       *event.Publisher
	leaveRepository facade.LeaveRepository
	leaveFactory    *LeaveFactory
}

func NewLeaveDomainService(publisher *event.Publisher, leaveRepository facade.LeaveRepository, leaveFactory *LeaveFactory) *LeaveDomainService {
	return &LeaveDomainService{publisher: publisher, leaveRepository: leaveRepository, leaveFactory: leaveFactory}
}

func (l *LeaveDomainService) CreateLeave(leave *entity.Leave, leaderMaxLevel int, approver approver.Approver) {
	leave.LeaveMaxLevel = leaderMaxLevel
	leave.Approver = approver
	leave.Create()
	l.leaveRepository.Save(l.leaveFactory.CreateLeavePO(leave))

	//发布事件
	event := event2.NewLeaveEvent(event2.CreateEvent, leave)
	l.leaveRepository.SaveEvent(l.leaveFactory.CreateLeaveEventPO(event))

}
