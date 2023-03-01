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

type voiceCallRepository struct {
}

type VoiceCallRepository interface {
	GetAll() (*[]model.VoiceCallData, error)
}

func NewVoiceCallRepository() VoiceCallRepository {
	return &voiceCallRepository{}
}

func (vr *voiceCallRepository) GetAll() (*[]model.VoiceCallData, error) {
	data, err := vr.readFromFile()
	if err != nil {
		return nil, err
	}
	splittedData, err := vr.splitToArrayStrings(data)
	if err != nil {
		return nil, err
	}
	parsedData, err := vr.parseDataToStruct(splittedData)

	return parsedData, nil
}

func (vr *voiceCallRepository) readFromFile() ([]byte, error) {
	file, err := os.Open(app_settings.AppSettings.Configuration.VoiceCallDataFilePath)
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

func (vr *voiceCallRepository) splitToArrayStrings(undividedData []byte) ([]string, error) {
	strData := string(undividedData)
	linesData := strings.Split(strData, "\n")

	if len(linesData) == 0 {
		return nil, errors.New("error split data to array strings")
	}

	return linesData, nil
}

func (vr *voiceCallRepository) parseDataToStruct(smsList []string) (*[]model.VoiceCallData, error) {
	var voiceCallData []model.VoiceCallData
	for _, itemString := range smsList {
		voiceCallDataArrayString := strings.Split(itemString, ";")
		err := validateVoiceCallData(validateVoiceCallParams{data: voiceCallDataArrayString, skipParams: []string{"smsData"}}) // first check length of voiceCallDataArrayString array
		if err != nil {
			continue
		}
		// if data length check accepted, check data string array as a struct
		smsDataValidateParam := model.VoiceCallData{
			Country:  voiceCallDataArrayString[0],
			Provider: voiceCallDataArrayString[3],
		}
		err = validateVoiceCallData(validateVoiceCallParams{voiceCallData: smsDataValidateParam, skipParams: []string{"data"}})
		if err != nil {
			continue
		}

		countries := *app_settings.AppSettings.Countries
		country, ok := countries[voiceCallDataArrayString[0]] // find country if exists in map of countries
		if !ok {
			continue
		}
		connectionStabilityF64, err := strconv.ParseFloat(voiceCallDataArrayString[4], 2)
		if err != nil {
			continue
		}
		connectionStabilityF32 := float32(connectionStabilityF64) // convert float64 to float32
		ttfb, err := strconv.Atoi(voiceCallDataArrayString[5])
		if err != nil {
			continue
		}
		voicePurity, err := strconv.Atoi(voiceCallDataArrayString[6])
		if err != nil {
			continue
		}
		medianOfCallsTime, err := strconv.Atoi(voiceCallDataArrayString[7])
		if err != nil {
			continue
		}

		voiceCallDataItem := model.VoiceCallData{
			Country:             country.Alpha2Code,
			Bandwidth:           voiceCallDataArrayString[1],
			ResponseTime:        voiceCallDataArrayString[2],
			Provider:            voiceCallDataArrayString[3],
			ConnectionStability: connectionStabilityF32,
			TTFB:                ttfb,
			VoicePurity:         voicePurity,
			MedianOfCallsTime:   medianOfCallsTime,
		}
		voiceCallData = append(voiceCallData, voiceCallDataItem)
	}
	return &voiceCallData, nil
}

func validateVoiceCallData(params validateVoiceCallParams) error {

	for _, skipParam := range params.skipParams {
		// check skip params
		if skipParam != "data" {
			if len(params.data) != 8 {
				return errors.New("data len is incorrect ")
			}
		}
		if skipParam != "smsData" {
			countries := *app_settings.AppSettings.Countries
			providers := *app_settings.AppSettings.VoiceCallProviders

			_, ok := countries[params.voiceCallData.Country]
			if !ok {
				return errors.New("country is incorrect")
			}
			_, ok = providers[params.voiceCallData.Provider]
			if !ok {
				return errors.New("provider is incorrect")
			}
		}
	}
	return nil
}

type validateVoiceCallParams struct {
	data          []string
	voiceCallData model.VoiceCallData
	skipParams    []string // "data" or "smsData", or empty
}
