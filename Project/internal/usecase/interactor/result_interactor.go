package interactor

import (
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type resultInteractor struct {
	ResultRepository repository.ResultRepository
	ResultPresenter  presenter.ResultPresenter
}

type ResultInteractor interface {
	Get() ([]byte, error)
}

func NewResultInteractor(r repository.ResultRepository, p presenter.ResultPresenter) ResultInteractor {
	return &resultInteractor{r, p}
}

func (ri *resultInteractor) Get() ([]byte, error) {
	r, err := ri.ResultRepository.Get()
	return ri.ResultPresenter.ResponseResult(r), err
}
