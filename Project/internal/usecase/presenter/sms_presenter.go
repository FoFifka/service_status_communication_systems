package presenter

import (
	"skillsbox-diploma/internal/domain/model"
)

type SMSPresenter interface {
	ResponseSMS(sd *[]model.SMSData) *[]model.SMSData
}
