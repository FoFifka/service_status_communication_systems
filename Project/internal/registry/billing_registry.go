package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewBillingController() controller.BillingController {
	return controller.NewBillingController(r.NewBillingInteractor())
}

func (r *registry) NewBillingInteractor() interactor.BillingInteractor {
	return interactor.NewBillingInteractor(r.NewBillingRepository(), r.NewBillingPresenter())
}

func (r *registry) NewBillingRepository() ur.BillingRepository {
	return ir.NewBillingRepository()
}

func (r *registry) NewBillingPresenter() up.BillingPresenter {
	return ip.NewBillingPresenter()
}
