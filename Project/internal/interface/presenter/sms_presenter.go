package presenter

import "skillsbox-diploma/internal/domain/model"

type smsPresenter struct {
}
type SMSPresenter interface {
	ResponseSMS(*[]model.SMSData) *[]model.SMSData
}

func NewSMSPresenter() SMSPresenter {
	return &smsPresenter{}
}

func (sp *smsPresenter) ResponseSMS(s *[]model.SMSData) *[]model.SMSData {
	//for _, smsItem := range s {
	//	todo presenter
	//}
	return s
}
