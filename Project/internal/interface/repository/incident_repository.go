package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/pkg/app_settings"
)

type incidentRepository struct {
}

type IncidentRepository interface {
	GetAll() (*[]model.IncidentData, error)
}

func NewIncidentRepository() IncidentRepository {
	return &incidentRepository{}
}

func (ir *incidentRepository) GetAll() (*[]model.IncidentData, error) {
	data, err := ir.requestToURLForGetData()
	if err != nil {
		return nil, err
	}
	var incidentData []model.IncidentData
	err = json.Unmarshal(data, &incidentData)
	if err != nil {
		return nil, err
	}

	return &incidentData, nil
}

func (ir *incidentRepository) requestToURLForGetData() ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, app_settings.AppSettings.Configuration.IncidentDataUrlAddress, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
