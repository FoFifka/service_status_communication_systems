package repository

import "skillsbox-diploma/internal/domain/model"

type BillingRepository interface {
	GetAll() (*model.BillingData, error)
}
