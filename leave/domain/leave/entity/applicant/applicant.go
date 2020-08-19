package applicant

import "leave/domain/shared"

/*
请假申请人值对象
 */
type Applicant struct {
	PersonId string
	PersonName string
	PersonType shared.PersonType
}

func NewApplicant(personId string, personName string, personType shared.PersonType) *Applicant {
	return &Applicant{PersonId: personId, PersonName: personName, PersonType: personType}
}

