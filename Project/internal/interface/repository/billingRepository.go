package repository

import (
	"errors"
	"log"
	"math"
	"os"
	"skillsbox-diploma/internal/domain/model"
	"skillsbox-diploma/pkg/app_settings"
)

type billingRepository struct {
}

type BillingRepository interface {
	GetAll() (*model.BillingData, error)
}

func NewBillingRepository() BillingRepository {
	return &billingRepository{}
}

func (br *billingRepository) GetAll() (*model.BillingData, error) {
	data, err := br.readDataFromFile()
	if err != nil {
		return nil, err
	}
	splittedData, err := br.splitToArrayStrings(data)

	parsedData, err := br.parseDataToStruct(splittedData)

	return parsedData, nil // TODO billing repository
}

func (br *billingRepository) readDataFromFile() ([]byte, error) {
	file, err := os.Open(app_settings.AppSettings.Configuration.BillingDataFilePath)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	fileStat, _ := file.Stat()
	fileSize := fileStat.Size()
	buf := make([]byte, fileSize)
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (br *billingRepository) splitToArrayStrings(data []byte) ([]string, error) {
	var stringsArray []string
	for _, byteItem := range data {
		stringsArray = append(stringsArray, string(byteItem))
	}
	return stringsArray, nil
}

func (br *billingRepository) parseDataToStruct(data []string) (*model.BillingData, error) {
	var billingData model.BillingData

	err := br.validateData(data)
	if err != nil {
		return nil, err
	}
	var reversedData []string
	for i := len(data) - 1; i >= 0; i-- {
		reversedData = append(reversedData, data[i])
	}

	sumOfPowsOfTwo := 0
	for i, dataItem := range reversedData {
		if dataItem == "0" {
			continue
		}
		powOfTwo := int(math.Pow(2, float64(i)))
		sumOfPowsOfTwo += powOfTwo
	}
	billingData = br.parseToStruct(data)

	return &billingData, nil
}

func (br *billingRepository) validateData(data []string) error {
	if len(data) != 6 {
		return errors.New("data length is incorrect")
	}
	return nil
}

func (br *billingRepository) parseToStruct(data []string) model.BillingData {
	createCustomer := false
	purchase := false
	payout := false
	recurring := false
	fraudControl := false
	checkoutPage := false
	if data[0] == "1" {
		createCustomer = true
	}
	if data[1] == "1" {
		purchase = true
	}
	if data[2] == "1" {
		payout = true
	}
	if data[3] == "1" {
		recurring = true
	}
	if data[4] == "1" {
		fraudControl = true
	}
	if data[5] == "1" {
		checkoutPage = true
	}

	return model.BillingData{
		CreateCustomer: createCustomer,
		Purchase:       purchase,
		Payout:         payout,
		Recurring:      recurring,
		FraudControl:   fraudControl,
		CheckoutPage:   checkoutPage,
	}
}
