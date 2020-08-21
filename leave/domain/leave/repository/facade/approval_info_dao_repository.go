package facade

import "leave/domain/leave/repository/po"

type ApprovalInfoDaoRepository interface {
	SaveAll(list []po.ApprovalInfoPO) error
	QueryByApplicantId(appliantId string) []*po.ApprovalInfoPO
	QueryByLeaveId(leaveId string) []*po.ApprovalInfoPO
}
