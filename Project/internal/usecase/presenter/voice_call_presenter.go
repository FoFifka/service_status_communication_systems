package presenter

import "skillsbox-diploma/internal/domain/model"

type VoiceCallPresenter interface {
	ResponseVoiceCall(vd *[]model.VoiceCallData) []byte
}
