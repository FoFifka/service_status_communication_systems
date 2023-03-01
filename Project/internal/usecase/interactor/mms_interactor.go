package interactor

import (
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type mmsInteractor struct {
	MMSRepository repository.MMSRepository
	MMSPresenter  presenter.MMSPresenter
}

type MMSInteractor interface {
	Get() ([]byte, error)
}

func NewMMSInteractor(r repository.MMSRepository, p presenter.MMSPresenter) MMSInteractor {
	return &mmsInteractor{r, p}
}

func (mi *mmsInteractor) Get() ([]byte, error) {
	m, err := mi.MMSRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return mi.MMSPresenter.ResponseMMS(m), err
}
