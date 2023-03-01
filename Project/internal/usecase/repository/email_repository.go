package repository

import "skillsbox-diploma/internal/domain/model"

type EmailRepository interface {
	GetAll() (*[]model.EmailData, error)
}
