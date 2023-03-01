package presenter

import (
	"encoding/json"
	"log"
	"skillsbox-diploma/internal/domain/model"
)

type voiceCallPresenter struct {
}

type VoiceCallPresenter interface {
	ResponseVoiceCall(*[]model.VoiceCallData) []byte
}

func NewVoiceCallPreneter() VoiceCallPresenter {
	return &voiceCallPresenter{}
}

func (vp *voiceCallPresenter) ResponseVoiceCall(vcd *[]model.VoiceCallData) []byte {

	voiceCallResponse, err := json.Marshal(vcd)
	if err != nil {
		log.Fatalln(err)
	}

	return voiceCallResponse
}
