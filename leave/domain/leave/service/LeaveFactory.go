package service

import (
	"leave/domain/leave/entity"
	"leave/domain/leave/event"
	"leave/domain/leave/repository/po"
	"time"
)

type LeaveFactory struct {
}

func (l *LeaveFactory) CreateLeavePO(leave *entity.Leave) *po.LeavePO {

	return &po.LeavePO{
		Id:                        leave.Id,
		ApplicantId:               leave.Applicant.PersonId,
		ApplicantName:             leave.Applicant.PersonName,
		ApplicantType:             leave.Applicant.PersonType,
		ApproverId:                leave.Approver.PersonId,
		ApproverName:              leave.Approver.PersonName,
		LeaveType:                 leave.LeaveType,
		Status:                    leave.Status,
		StartTime:                 time.Time{},
		EndTime:                   time.Time{},
		Duration:                  0,
		HistoryApprovalInfoPOList: nil,
	}
}

func (l *LeaveFactory) CreateLeaveEventPO(e *event.LeaveEvent) *po.LeaveEventPO {

	return &po.LeaveEventPO{
		Id:             e.Id,
		LeaveEventType: e.LeaveEventType,
		Timestamp:      e.Timestamp,
		Source:         e.Source,
		Data:           e.Data,
	}
}
