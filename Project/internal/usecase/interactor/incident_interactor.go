package interactor

import (
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type incidentInteractor struct {
	IncidentRepository repository.IncidentRepository
	IncidentPresenter  presenter.IncidentPresenter
}

type IncidentInteractor interface {
	Get() ([]byte, error)
}

func NewIncidentInteractor(r repository.IncidentRepository, p presenter.IncidentPresenter) IncidentInteractor {
	return &incidentInteractor{r, p}
}

func (ii *incidentInteractor) Get() ([]byte, error) {
	i, err := ii.IncidentRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return ii.IncidentPresenter.ResponseIncident(i), nil
}
