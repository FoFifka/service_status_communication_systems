package presenter

import (
	"encoding/json"
	"log"
	"skillsbox-diploma/internal/domain/model"
)

type emailPresenter struct {
}

type EmailPresenter interface {
	ResponseEmail(*[]model.EmailData) []byte
}

func NewEmailPresenter() EmailPresenter {
	return &emailPresenter{}
}

func (ep *emailPresenter) ResponseEmail(ed *[]model.EmailData) []byte {
	emailResponse, err := json.Marshal(ed)
	if err != nil {
		log.Fatalln(err)
	}

	return emailResponse
}
