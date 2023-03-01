package app_settings

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

var AppSettings *Settings

type Configuration struct {
	Host                   string `json:"host"`
	Port                   string `json:"port"`
	SMSDataFilePath        string `json:"sms_data_file_path"`
	MMSDataUrlAddress      string `json:"mms_data_url_address"`
	VoiceCallDataFilePath  string `json:"voice_call_data_file_path"`
	EmailDataFilePath      string `json:"email_data_file_path"`
	BillingDataFilePath    string `json:"billing_data_file_path"`
	SupportDataUrlAddress  string `json:"support_data_url_address"`
	IncidentDataUrlAddress string `json:"incident_data_url_address"`
}

type Settings struct {
	Configuration      *Configuration
	Countries          *map[string]Country
	Providers          *map[string]Provider
	VoiceCallProviders *map[string]Provider
	EmailProviders     *map[string]Provider
}

type Country struct {
	Alpha2Code string
}

type Provider struct {
	Name string
}

func SetSettings() *Settings {
	settings := Settings{
		Configuration:      getConfigurationFromFile(),
		Countries:          getCountryCodesFromFile(),
		Providers:          getProvidersFromFile(),
		VoiceCallProviders: getVoiceCallProvidersFromFile(),
		EmailProviders:     getEmailProvidersFromFile(),
	}
	return &settings
}

func getConfigurationFromFile() *Configuration {
	file, err := os.Open("../pkg/app_settings/configuration.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Fatalln(err)
	}
	return &configuration
}

func getCountryCodesFromFile() *map[string]Country {
	// read from file
	file, err := os.Open("../pkg/app_settings/countryCodes.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	fileSize := fileStat.Size()
	buf := make([]byte, fileSize)
	_, err = file.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// parse data to struct

	countryCodes := make(map[string]Country)
	strData := string(buf)
	strDataArray := strings.Split(strData, "\n")
	for _, countryDataStringItem := range strDataArray {
		countryData := strings.Split(countryDataStringItem, ";")
		countryCodes[countryData[0]] = Country{
			Alpha2Code: countryData[0],
		}
	}

	return &countryCodes
}

func getProvidersFromFile() *map[string]Provider {
	// read from file
	file, err := os.Open("../pkg/app_settings/providersList.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	fileSize := fileStat.Size()
	buf := make([]byte, fileSize)
	_, err = file.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// parse data to struct
	providers := make(map[string]Provider)
	strData := string(buf)

	strDataArray := strings.Split(strData, "\n")

	for _, providerStringItem := range strDataArray {
		providerStringItem = strings.TrimSpace(providerStringItem) // without it, map is confuse
		providers[providerStringItem] = Provider{
			Name: providerStringItem,
		}
	}
	return &providers
}

func getVoiceCallProvidersFromFile() *map[string]Provider {
	// read from file
	file, err := os.Open("../pkg/app_settings/voiceCallProvidersList.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	fileSize := fileStat.Size()
	buf := make([]byte, fileSize)
	_, err = file.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// parse data to struct
	providers := make(map[string]Provider)
	strData := string(buf)

	strDataArray := strings.Split(strData, "\n")

	for _, providerStringItem := range strDataArray {
		providerStringItem = strings.TrimSpace(providerStringItem) // without it, map is confuse
		providers[providerStringItem] = Provider{
			Name: providerStringItem,
		}
	}
	return &providers
}

func getEmailProvidersFromFile() *map[string]Provider {
	// read from file
	file, err := os.Open("../pkg/app_settings/emailProvidersList.txt")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}
	fileSize := fileStat.Size()
	buf := make([]byte, fileSize)
	_, err = file.Read(buf)
	if err != nil {
		log.Fatalln(err)
	}

	// parse data to struct
	providers := make(map[string]Provider)
	strData := string(buf)

	strDataArray := strings.Split(strData, "\n")

	for _, providerStringItem := range strDataArray {
		providerStringItem = strings.TrimSpace(providerStringItem) // without it, map is confuse
		providers[providerStringItem] = Provider{
			Name: providerStringItem,
		}
	}
	return &providers
}
