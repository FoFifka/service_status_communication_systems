package repository

import "skillsbox-diploma/internal/domain/model"

type IncidentRepository interface {
	GetAll() (*[]model.IncidentData, error)
}
