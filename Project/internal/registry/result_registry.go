package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewResultController() controller.ResultController {
	return controller.NewResultController(r.NewResultInteractor())
}

func (r *registry) NewResultInteractor() interactor.ResultInteractor {
	return interactor.NewResultInteractor(r.NewResultRepository(), r.NewResultPresenter())
}

func (r *registry) NewResultRepository() ur.ResultRepository {
	return ir.NewResultRepository()
}

func (r *registry) NewResultPresenter() up.ResultPresenter {
	return ip.NewResultPresenter()
}
