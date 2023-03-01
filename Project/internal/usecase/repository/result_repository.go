package repository

import "skillsbox-diploma/internal/domain/model"

type ResultRepository interface {
	Get() (model.ResultT, error)
}
