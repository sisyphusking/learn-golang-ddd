package facade

import "leave/domain/leave/repository/po"

type LeaveEventDaoRepository interface {
	Save(po *po.LeaveEventPO) error
}
