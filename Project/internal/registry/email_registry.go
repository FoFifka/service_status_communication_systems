package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewEmailController() controller.EmailController {
	return controller.NewEmailController(r.NewEmailInteractor())
}

func (r *registry) NewEmailInteractor() interactor.EmailInteractor {
	return interactor.NewEmailInteractor(r.NewEmailRepository(), r.NewEmailPresenter())
}

func (r *registry) NewEmailRepository() ur.EmailRepository {
	return ir.NewEmailRepository()
}

func (r *registry) NewEmailPresenter() up.EmailPresenter {
	return ip.NewEmailPresenter()
}
