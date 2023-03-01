package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewSupportController() controller.SupportController {
	return controller.NewSupportController(r.NewSupportInteractor())
}

func (r *registry) NewSupportInteractor() interactor.SupportInteractor {
	return interactor.NewSupportInteractor(r.NewSupportRepository(), r.NewSupportPresenter())
}

func (r *registry) NewSupportPresenter() up.SupportPresenter {
	return ip.NewSupportPresenter()
}

func (r *registry) NewSupportRepository() ur.SupportRepository {
	return ir.NewSupportRepository()
}
