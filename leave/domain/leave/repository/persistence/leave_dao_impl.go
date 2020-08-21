package persistence

import (
	"leave/domain/leave/repository/po"

	"github.com/jinzhu/gorm"
)

type LeaveDaoImpl struct {
	client *gorm.DB
}

func (l LeaveDaoImpl) FindById(leaveId string) *po.LeavePO {
	panic("implement me")
}

func (l LeaveDaoImpl) Save(leavePO *po.LeavePO) error {
	//自动迁移模式
	l.client.AutoMigrate(&po.LeavePO{})
	l.client.Create(leavePO)
	return nil
}

func (l LeaveDaoImpl) QueryByApplicantId(applicantId string) []*po.LeavePO {
	panic("implement me")
}

func (l LeaveDaoImpl) QueryByApproverId(approverId string) []*po.LeavePO {
	panic("implement me")
}
