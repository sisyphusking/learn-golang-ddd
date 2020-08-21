package po

import (
	"leave/domain/leave/entity"
	"leave/domain/shared"
	"time"
)

/*
对应数据库对象
*/

type LeavePO struct {
	Id                        string
	ApplicantId               string
	ApplicantName             string
	ApplicantType             shared.PersonType
	ApproverId                string
	ApproverName              string
	LeaveType                 shared.LeaveType
	Status                    entity.Status
	StartTime                 time.Time
	EndTime                   time.Time
	Duration                  int64
	HistoryApprovalInfoPOList []ApprovalInfoPO
}

func (l *LeavePO) GetHistoryApprovalInfoPOList() []ApprovalInfoPO {
	return l.HistoryApprovalInfoPOList
}
