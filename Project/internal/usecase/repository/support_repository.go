package repository

import "skillsbox-diploma/internal/domain/model"

type SupportRepository interface {
	GetAll() (*[]model.SupportData, error)
}
