package repository

import "skillsbox-diploma/internal/domain/model"

type MMSRepository interface {
	GetAll() (*[]model.MMSData, error)
}
