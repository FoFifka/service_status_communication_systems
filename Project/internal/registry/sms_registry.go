package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewSMSController() controller.SMSController {
	return controller.NewSMSController(r.NewSMSInteractor())
}

func (r *registry) NewSMSInteractor() interactor.SMSInteractor {
	return interactor.NewSMSInteractor(r.NewSMSRepository(), r.NewSMSPresenter())
}

func (r *registry) NewSMSRepository() ur.SMSRepository {
	return ir.NewSMSRepository()
}

func (r *registry) NewSMSPresenter() up.SMSPresenter {
	return ip.NewSMSPresenter()
}
