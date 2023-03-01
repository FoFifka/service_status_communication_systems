package interactor

import (
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type billingInteractor struct {
	BillingRepository repository.BillingRepository
	BillingPresenter  presenter.BillingPresenter
}

type BillingInteractor interface {
	Get() ([]byte, error)
}

func NewBillingInteractor(r repository.BillingRepository, p presenter.BillingPresenter) BillingInteractor {
	return &billingInteractor{r, p}
}

func (bi *billingInteractor) Get() ([]byte, error) {
	b, err := bi.BillingRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return bi.BillingPresenter.ResponseBilling(b), nil
}
