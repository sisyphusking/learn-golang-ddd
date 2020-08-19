package entity

import "leave/domain/leave/entity/approver"

type ApprovalInfo struct {
	ApprovalInfoId string
	Approver approver.Approver
	ApprovalType approver.ApprovalType
}
