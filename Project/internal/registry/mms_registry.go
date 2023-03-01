package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewMMSController() controller.MMSController {
	return controller.NewMMSController(r.NewMMSInteractor())
}

func (r *registry) NewMMSInteractor() interactor.MMSInteractor {
	return interactor.NewMMSInteractor(r.NewMMSRepository(), r.NewMMSPresenter())
}

func (r *registry) NewMMSRepository() ur.MMSRepository {
	return ir.NewMMSRepository()
}

func (r *registry) NewMMSPresenter() up.MMSPresenter {
	return ip.NewMMSPresenter()
}
