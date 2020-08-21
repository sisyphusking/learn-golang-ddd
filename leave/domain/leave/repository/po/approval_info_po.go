package po

type ApprovalInfoPO struct {
	ApprovalInfoId string
	LeaveId        string
	ApplicantId    string
	ApproverId     string
	ApproverLevel  int
	ApproverName   string
	Msg            string
	Time           int64
}

func (a *ApprovalInfoPO) GetLeaveId() string {
	return a.LeaveId
}

func (a *ApprovalInfoPO) SetLeaveId(leaveId string) {
	a.LeaveId = leaveId
}
