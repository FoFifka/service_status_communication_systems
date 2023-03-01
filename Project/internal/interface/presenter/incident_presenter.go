package presenter

import (
	"encoding/json"
	"log"
	"skillsbox-diploma/internal/domain/model"
)

type incidentPresenter struct {
}

type IncidentPresenter interface {
	ResponseIncident(*[]model.IncidentData) []byte
}

func NewIncidentPresenter() IncidentPresenter {
	return &incidentPresenter{}
}

func (ip *incidentPresenter) ResponseIncident(id *[]model.IncidentData) []byte {
	responseIncident, err := json.Marshal(id)
	if err != nil {
		log.Fatal(err)
	}
	return responseIncident
}
