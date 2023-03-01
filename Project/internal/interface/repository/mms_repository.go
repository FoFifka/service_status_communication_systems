package repository

import (
	"encoding/json"
	"io"
	"net/http"
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/pkg/app_settings"
)

type mmsRepository struct {
}

type MMSRepository interface {
	GetAll() (*[]model.MMSData, error)
}

func NewMMSRepository() MMSRepository {
	return &mmsRepository{}
}

func (sr *mmsRepository) GetAll() (*[]model.MMSData, error) {
	data, err := sr.requestToURLForGetData()
	if err != nil {
		return nil, err
	}
	var parsedData []model.MMSData
	err = json.Unmarshal(data, &parsedData)
	if err != nil {
		return nil, err
	}

	filteredData, err := sr.filterData(parsedData)
	return filteredData, nil
}

func (mr *mmsRepository) requestToURLForGetData() ([]byte, error) {
	request, err := http.NewRequest(http.MethodGet, app_settings.AppSettings.Configuration.MMSDataUrlAddress, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func (sr *mmsRepository) filterData(mmsList []model.MMSData) (*[]model.MMSData, error) {
	var mmsData []model.MMSData

	for _, mmsDataItem := range mmsList {
		countries := *app_settings.AppSettings.Countries
		country, ok := countries[mmsDataItem.Country]
		if !ok {
			continue
		}
		providers := *app_settings.AppSettings.Providers
		provider, ok := providers[mmsDataItem.Provider]
		if !ok {
			continue
		}
		mmsData = append(mmsData, model.MMSData{
			Country:      country.Alpha2Code,
			Bandwidth:    mmsDataItem.Bandwidth,
			ResponseTime: mmsDataItem.ResponseTime,
			Provider:     provider.Name,
		})
	}

	return &mmsData, nil
}
