package facade

import "leave/domain/leave/repository/po"

type LeaveRepository interface {
	Save(leavePO *po.LeavePO) error
	SaveEvent(eventPO *po.LeaveEventPO) error
	FindById(id string) (*po.LeavePO, error)
	QueryByApplicantId(applicantId string) ([]*po.LeavePO, error)
	QueryByApproverId(approverId string) ([]*po.LeavePO, error)
}
