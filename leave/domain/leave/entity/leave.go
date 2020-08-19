package entity

import (
	"leave/domain/leave/entity/applicant"
	"leave/domain/leave/entity/approver"
	"leave/domain/shared"
	"time"
)

/*
请假单聚合根
 */
type Leave struct {
	Id string
	Applicant applicant.Applicant
	Approver approver.Approver
	LeaveType shared.LeaveType
	Status Status
	StartTime time.Time
	EndTime time.Time
	Duration  int64
	LeaveMaxLevel int //审批领导最大级别
	CurrentApprovalInfo approver.ApprovalType
	HistoryApprovalInfos []ApprovalInfo

}

