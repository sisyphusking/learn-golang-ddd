package facade

import "leave/domain/leave/repository/po"

type LeaveDaoRepository interface {
	FindById(leaveId string) *po.LeavePO
	Save(leavePO *po.LeavePO) error
	QueryByApplicantId(applicantId string) []*po.LeavePO
	QueryByApproverId(approverId string) []*po.LeavePO
}
