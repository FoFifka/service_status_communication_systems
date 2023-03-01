package registry

import (
	"skillsbox-diploma/internal/interface/controller"
)

type registry struct {
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry() Registry {
	return &registry{}
}

func (r *registry) NewAppController() controller.AppController {
	appController := controller.NewAppController(
		r.NewSMSController(),
		r.NewMMSController(),
		r.NewEmailController(),
		r.NewVoiceCallController(),
		r.NewBillingController(),
		r.NewSupportController(),
		r.NewIncidentController(),
		r.NewResultController(),
	)

	return appController
}
