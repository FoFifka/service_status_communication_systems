package presenter

import "skillsbox-diploma/internal/domain/model"

type BillingPresenter interface {
	ResponseBilling(bd *model.BillingData) []byte
}
