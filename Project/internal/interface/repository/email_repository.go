package repository

import (
	"errors"
	"log"
	"os"
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/pkg/app_settings"
	"strconv"
	"strings"
)

type emailRepository struct {
}

type EmailRepository interface {
	GetAll() (*[]model.EmailData, error)
}

func NewEmailRepository() EmailRepository {
	return &emailRepository{}
}

func (er *emailRepository) GetAll() (*[]model.EmailData, error) {
	data, err := er.readFromFile()
	if err != nil {
		return nil, err
	}
	splittedData, err := er.splitToArrayStrings(data)
	if err != nil {
		return nil, err
	}
	parsedData, err := er.parseDataToStruct(splittedData)
	if err != nil {
		return nil, err
	}
	return parsedData, nil
}

func (er *emailRepository) readFromFile() ([]byte, error) {
	file, err := os.Open(app_settings.AppSettings.Configuration.EmailDataFilePath)
	if err != nil {
		log.Panicln("Ошибка при открытии файла", err)
		return nil, err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	fileSize := fileStat.Size()
	if fileSize == 0 {
		return nil, errors.New("file is empty")
	}
	buf := make([]byte, fileSize)
	_, err = file.Read(buf)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}

	return buf, nil
}

func (er *emailRepository) splitToArrayStrings(undividedData []byte) ([]string, error) {
	strData := string(undividedData)
	linesData := strings.Split(strData, "\n")

	if len(linesData) == 0 {
		return nil, errors.New("error split data to array strings")
	}

	return linesData, nil
}

func (er *emailRepository) parseDataToStruct(emailList []string) (*[]model.EmailData, error) {
	var emailData []model.EmailData
	for _, itemString := range emailList {
		emailDataArrayString := strings.Split(itemString, ";")
		err := validateEmailData(validateEmailParams{data: emailDataArrayString, skipParams: []string{"emailData"}}) // first check length of emailDataArrayString array
		if err != nil {
			continue
		}
		// if data length check accepted, check data string array as a struct
		emailDataValidateParam := model.EmailData{
			Country:  emailDataArrayString[0],
			Provider: emailDataArrayString[1],
		}
		err = validateEmailData(validateEmailParams{emailData: emailDataValidateParam, skipParams: []string{"data"}})
		if err != nil {
			continue
		}

		countries := *app_settings.AppSettings.Countries
		country, ok := countries[emailDataArrayString[0]] // find country if exists in map of countries
		if !ok {
			continue
		}
		deliveryTime, err := strconv.Atoi(emailDataArrayString[2])
		if err != nil {
			continue
		}
		emailDataItem := model.EmailData{
			Country:      country.Alpha2Code,
			Provider:     emailDataArrayString[1],
			DeliveryTime: deliveryTime,
		}
		emailData = append(emailData, emailDataItem)
	}
	return &emailData, nil
}

func validateEmailData(params validateEmailParams) error {

	for _, skipParam := range params.skipParams {
		// check skip params
		if skipParam != "data" {
			if len(params.data) != 3 {
				return errors.New("data len is incorrect ")
			}
		}
		if skipParam != "emailData" {
			countries := *app_settings.AppSettings.Countries
			providers := *app_settings.AppSettings.EmailProviders

			_, ok := countries[params.emailData.Country]
			if !ok {
				return errors.New("country is incorrect")
			}
			_, ok = providers[params.emailData.Provider]
			if !ok {
				return errors.New("provider is incorrect")
			}
		}
	}
	return nil
}

type validateEmailParams struct {
	data       []string
	emailData  model.EmailData
	skipParams []string // "data" or "emailData", or empty
}
