package controller

type appController struct {
	SMSController
	MMSController
	EmailController
	VoiceCallController
	BillingController
	SupportController
	IncidentController
	ResultController
}

type AppController interface {
	SMSController
	MMSController
	EmailController
	VoiceCallController
	BillingController
	SupportController
	IncidentController
	IncidentController
	ResultController
}

func NewAppController(sms SMSController, mms MMSController, email EmailController, voiceCall VoiceCallController, billing BillingController, support SupportController, incident IncidentController, result ResultController) AppController {
	return &appController{
		sms,
		mms,
		email,
		voiceCall,
		billing,
		support,
		incident,
		result,
	}
}
