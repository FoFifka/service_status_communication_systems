package presenter

import "skillsbox-diploma/internal/domain/model"

type EmailPresenter interface {
	ResponseEmail(ed *[]model.EmailData) []byte
}
