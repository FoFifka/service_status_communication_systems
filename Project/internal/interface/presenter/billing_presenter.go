package presenter

import (
	"encoding/json"
	"log"
	"skillsbox-diploma/internal/domain/model"
)

type billingPresenter struct {
}

type BillingPresenter interface {
	ResponseBilling(*model.BillingData) []byte
}

func NewBillingPresenter() BillingPresenter {
	return &billingPresenter{}
}

func (bp *billingPresenter) ResponseBilling(bd *model.BillingData) []byte {
	responseBilling, err := json.Marshal(bd)
	if err != nil {
		log.Fatalln(err)
	}
	return responseBilling
}
