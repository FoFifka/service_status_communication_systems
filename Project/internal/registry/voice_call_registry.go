package registry

import (
	"skillsbox-diploma/internal/interface/controller"
	ip "skillsbox-diploma/internal/interface/presenter"
	ir "skillsbox-diploma/internal/interface/repository"
	"skillsbox-diploma/internal/usecase/interactor"
	up "skillsbox-diploma/internal/usecase/presenter"
	ur "skillsbox-diploma/internal/usecase/repository"
)

func (r *registry) NewVoiceCallController() controller.VoiceCallController {
	return controller.NewVoiceCallController(r.NewVoiceCallInteractor())
}

func (r *registry) NewVoiceCallInteractor() interactor.VoiceCallInteractor {
	return interactor.NewVoiceCallInteractor(r.NewVoiceCallRepository(), r.NewVoiceCallPresenter())
}

func (r *registry) NewVoiceCallPresenter() up.VoiceCallPresenter {
	return ip.NewVoiceCallPreneter()
}

func (r *registry) NewVoiceCallRepository() ur.VoiceCallRepository {
	return ir.NewVoiceCallRepository()
}
