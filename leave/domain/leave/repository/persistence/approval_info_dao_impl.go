package persistence

import (
	"fmt"
	"leave/domain/leave/repository/po"

	"github.com/jinzhu/gorm"
)

type ApprovalInfoImpl struct {
	client *gorm.DB
}

func (a ApprovalInfoImpl) SaveAll(list []*po.ApprovalInfoPO) error {
	fmt.Sprintf("save approval info list: %v", list)
	panic("implement me")
}

func (a ApprovalInfoImpl) QueryByApplicantId(appliantId string) []*po.ApprovalInfoPO {
	panic("implement me")
}

func (a ApprovalInfoImpl) QueryByLeaveId(leaveId string) []*po.ApprovalInfoPO {
	panic("implement me")
}
