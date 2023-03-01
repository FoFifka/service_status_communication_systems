package repository

import "skillsbox-diploma/internal/domain/model"

type VoiceCallRepository interface {
	GetAll() (*[]model.VoiceCallData, error)
}
