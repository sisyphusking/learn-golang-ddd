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
	Id                   string
	Applicant            applicant.Applicant
	Approver             approver.Approver
	LeaveType            shared.LeaveType
	Status               Status
	StartTime            time.Time
	EndTime              time.Time
	Duration             int64
	LeaveMaxLevel        int //审批领导最大级别
	CurrentApprovalInfo  approver.ApprovalType
	HistoryApprovalInfos []ApprovalInfo
}

/*
充血模型
*/

//请求单审批时长
func (l *Leave) GetDuration() int64 {
	return int64(l.EndTime.Sub(l.StartTime))
}

//添加审批记录
func (l *Leave) AddHistoryApprovalInfo(info ApprovalInfo) {
	l.HistoryApprovalInfos = append(l.HistoryApprovalInfos, info)
}

//创建审批单
func (l *Leave) Create() {
	l.Status = APPROVING
	l.StartTime = time.Now()
}

//同意后设置下一个审批人
func (l *Leave) Agree(nextApprover approver.Approver) {
	l.Status = APPROVING
	l.Approver = nextApprover
}

func (l *Leave) Reject(approver approver.Approver) {
	l.Approver = approver
	l.Status = REJECTED
}

func (l *Leave) Finish() *Leave {
	l.Approver = approver.Approver{}
	l.Status = APPROVED
	l.EndTime = time.Now()
	l.Duration = l.GetDuration()
	return l
}
