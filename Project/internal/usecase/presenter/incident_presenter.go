package presenter

import "skillsbox-diploma/internal/domain/model"

type IncidentPresenter interface {
	ResponseIncident(*[]model.IncidentData) []byte
}
