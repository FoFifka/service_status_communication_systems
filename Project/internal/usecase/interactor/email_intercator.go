package interactor

import (
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type emailInteractor struct {
	EmailRepository repository.EmailRepository
	EmailPresenter  presenter.EmailPresenter
}

type EmailInteractor interface {
	Get() ([]byte, error)
}

func NewEmailInteractor(r repository.EmailRepository, p presenter.EmailPresenter) EmailInteractor {
	return &emailInteractor{r, p}
}

func (ei *emailInteractor) Get() ([]byte, error) {
	e, err := ei.EmailRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return ei.EmailPresenter.ResponseEmail(e), nil
}
