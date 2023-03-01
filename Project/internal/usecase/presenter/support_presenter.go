package presenter

import "skillsbox-diploma/internal/domain/model"

type SupportPresenter interface {
	ResponseSupport(*[]model.SupportData) []byte
}
