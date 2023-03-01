package repository

import (
	"errors"
	"os"
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/pkg/app_settings"
	"strings"
)

type smsRepository struct {
}

type SMSRepository interface {
	GetAll() (*[]model.SMSData, error)
}

func NewSMSRepository() SMSRepository {
	return &smsRepository{}
}

func (sr *smsRepository) GetAll() (*[]model.SMSData, error) {
	data, err := sr.readFromFile()
	if err != nil {
		return nil, err
	}
	splittedData, err := sr.splitToArrayStrings(data)
	if err != nil {
		return nil, err
	}
	parsedData, err := sr.parseDataToStruct(splittedData) // todo smsData variable
	if err != nil {
		return nil, err
	}
	return parsedData, nil
}

func (sr *smsRepository) readFromFile() ([]byte, error) {
	file, err := os.Open(app_settings.AppSettings.Configuration.SMSDataFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileStat.Size()
	if fileSize == 0 {
		return nil, errors.New("file is empty")
	}
	buf := make([]byte, fileSize)
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (sr *smsRepository) splitToArrayStrings(undividedData []byte) ([]string, error) {
	strData := string(undividedData)
	linesData := strings.Split(strData, "\n")

	if len(linesData) == 0 {
		return nil, errors.New("error split data to array strings")
	}

	return linesData, nil
}

func (sr *smsRepository) parseDataToStruct(smsList []string) (*[]model.SMSData, error) {
	var smsData []model.SMSData
	for _, itemString := range smsList {
		smsDataArrayString := strings.Split(itemString, ";")
		err := validateData(validateParams{data: smsDataArrayString, skipParams: []string{"smsData"}}) // first check length of smsDataArrayString array
		if err != nil {
			continue
		}
		// if data length check accepted, check data string array as a struct
		smsDataValidateParam := model.SMSData{
			Country:  smsDataArrayString[0],
			Provider: smsDataArrayString[3],
		}
		err = validateData(validateParams{smsData: smsDataValidateParam, skipParams: []string{"data"}})
		if err != nil {
			continue
		}

		countries := *app_settings.AppSettings.Countries
		country, ok := countries[smsDataArrayString[0]] // find country if exists in map of countries
		if !ok {
			continue
		}

		smsDataItem := model.SMSData{
			Country:      country.Alpha2Code,
			Bandwidth:    smsDataArrayString[1],
			ResponseTime: smsDataArrayString[2],
			Provider:     smsDataArrayString[3],
		}
		smsData = append(smsData, smsDataItem)
	}
	return &smsData, nil
}

func validateData(params validateParams) error {

	for _, skipParam := range params.skipParams {
		// check skip params
		if skipParam != "data" {
			if len(params.data) != 4 {
				return errors.New("data len is incorrect ")
			}
		}
		if skipParam != "smsData" {
			countries := *app_settings.AppSettings.Countries
			providers := *app_settings.AppSettings.Providers

			_, ok := countries[params.smsData.Country]
			if !ok {
				return errors.New("country is incorrect")
			}
			_, ok = providers[params.smsData.Provider]
			if !ok {
				return errors.New("provider is incorrect")
			}
		}
	}
	return nil
}

type validateParams struct {
	data       []string
	smsData    model.SMSData
	skipParams []string // "data" or "smsData", or empty
}
