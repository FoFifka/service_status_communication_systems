package presenter

import "skillsbox-diploma/internal/domain/model"

type MMSPresenter interface {
	ResponseMMS(md *[]model.MMSData) []byte
}
