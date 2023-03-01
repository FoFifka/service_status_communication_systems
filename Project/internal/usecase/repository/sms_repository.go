package repository

import (
	"skillsbox-diploma/internal/domain/model"
)

type SMSRepository interface {
	GetAll() (*[]model.SMSData, error)
}
