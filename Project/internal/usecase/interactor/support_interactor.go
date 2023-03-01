package interactor

import (
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type supportInteractor struct {
	SupportRepository repository.SupportRepository
	SupportPresenter  presenter.SupportPresenter
}

type SupportInteractor interface {
	Get() ([]byte, error)
}

func NewSupportInteractor(r repository.SupportRepository, p presenter.SupportPresenter) SupportInteractor {
	return &supportInteractor{r, p}
}

func (si *supportInteractor) Get() ([]byte, error) {
	s, err := si.SupportRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return si.SupportPresenter.ResponseSupport(s), nil
}
