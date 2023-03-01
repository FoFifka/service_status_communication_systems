package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/pkg/app_settings"
)

type supportRepository struct {
}

type SupportRepository interface {
	GetAll() (*[]model.SupportData, error)
}

func NewSupportRepository() SupportRepository {
	return &supportRepository{}
}

func (sr *supportRepository) GetAll() (*[]model.SupportData, error) {
	data, err := sr.requestToURLForGetData()
	if err != nil {
		return nil, err
	}

	var supportData []model.SupportData
	err = json.Unmarshal(data, &supportData)
	if err != nil {
		return nil, err
	}

	return &supportData, nil // TODO Support repository
}

func (sr *supportRepository) requestToURLForGetData() ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, app_settings.AppSettings.Configuration.SupportDataUrlAddress, nil)
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
