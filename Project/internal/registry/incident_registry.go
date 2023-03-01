package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewIncidentController() controller.IncidentController {
	return controller.NewIncidentController(r.NewIncidentInteractor())
}

func (r *registry) NewIncidentInteractor() interactor.IncidentInteractor {
	return interactor.NewIncidentInteractor(r.NewIncidentRepository(), r.NewIncidentPresenter())
}

func (r *registry) NewIncidentRepository() ur.IncidentRepository {
	return ir.NewIncidentRepository()
}

func (r *registry) NewIncidentPresenter() up.IncidentPresenter {
	return ip.NewIncidentPresenter()
}
