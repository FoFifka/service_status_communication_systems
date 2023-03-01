package interactor

import (
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type smsInteractor struct {
	SMSRepository repository.SMSRepository
	SMSPresenter  presenter.SMSPresenter
}

type SMSInteractor interface {
	Get(sd *[]model.SMSData) (*[]model.SMSData, error)
}

func NewSMSInteractor(r repository.SMSRepository, p presenter.SMSPresenter) SMSInteractor {
	return &smsInteractor{r, p}
}

func (si *smsInteractor) Get(sd *[]model.SMSData) (*[]model.SMSData, error) {
	s, err := si.SMSRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return si.SMSPresenter.ResponseSMS(s), nil
}
