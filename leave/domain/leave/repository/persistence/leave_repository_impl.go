package persistence

import (
	"fmt"
	"leave/domain/leave/repository/facade"
	"leave/domain/leave/repository/po"
)

type LeaveRepositoryImpl struct {
	//初始化资源
	leaveDao        facade.LeaveDaoRepository
	approvalInfoDao facade.ApprovalInfoDaoRepository
	leaveEventDao   facade.LeaveEventDaoRepository
}

func (l *LeaveRepositoryImpl) Save(leavePO *po.LeavePO) error {
	l.leaveDao.Save(leavePO)

	for _, historyApprovalInfo := range leavePO.HistoryApprovalInfoPOList {
		historyApprovalInfo.SetLeaveId(leavePO.Id)
	}
	l.approvalInfoDao.SaveAll(leavePO.GetHistoryApprovalInfoPOList())
	panic("implement me")
}

func (l *LeaveRepositoryImpl) SaveEvent(eventPO *po.LeaveEventPO) error {
	fmt.Sprintf("save event: %v", eventPO)
	panic("implement me")
}

func (l *LeaveRepositoryImpl) FindById(id string) (*po.LeavePO, error) {
	panic("implement me")
}

func (l *LeaveRepositoryImpl) QueryByApplicantId(applicantId string) ([]*po.LeavePO, error) {
	panic("implement me")
}

func (l *LeaveRepositoryImpl) QueryByApproverId(approverId string) ([]*po.LeavePO, error) {
	panic("implement me")
}
