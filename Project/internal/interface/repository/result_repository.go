package repository

import (
	"skillsbox-diploma/internal/domain/model"
)

type resultRepository struct {
}

type ResultRepository interface {
	Get() (model.ResultT, error)
}

func NewResultRepository() ResultRepository {
	return &resultRepository{}
}

func (rp *resultRepository) Get() (model.ResultT, error) {
	// TODO Result repository
	var resultT model.ResultT
	resultT.Status = false

	// SMS
	sms := smsRepository{}
	smsData, err := sms.GetAll()
	if err != nil {
		return rp.returnError(resultT, err)
	}
	resultT.Data.SMS = *smsData

	// MMS

	mms := mmsRepository{}
	mmsData, err := mms.GetAll()
	if err != nil {
		return rp.returnError(resultT, err)
	}
	resultT.Data.MMS = *mmsData

	// Email

	email := emailRepository{}
	emailData, err := email.GetAll()
	if err != nil {
		return rp.returnError(resultT, err)
	}
	resultT.Data.Email = *emailData

	// VoiceCall

	voiceCall := voiceCallRepository{}
	voiceCallData, err := voiceCall.GetAll()
	if err != nil {
		return rp.returnError(resultT, err)
	}
	resultT.Data.VoiceCall = *voiceCallData

	// Billing
	billing := billingRepository{}
	billingData, err := billing.GetAll()
	if err != nil {
		return rp.returnError(resultT, err)
	}
	resultT.Data.Billing = *billingData

	// Support
	support := supportRepository{}
	supportData, err := support.GetAll()
	if err != nil {
		return rp.returnError(resultT, err)
	}
	resultT.Data.Support = *supportData

	// Incident
	incident := incidentRepository{}
	incidentData, err := incident.GetAll()
	if err != nil {
		return rp.returnError(resultT, err)
	}
	resultT.Data.Incidents = *incidentData

	resultT.Status = true
	return resultT, nil
}

func (rr *resultRepository) returnError(resultT model.ResultT, err error) (model.ResultT, error) {
	resultT.Data = model.ResultSetT{}
	resultT.Error = err.Error()
	return resultT, err
}
