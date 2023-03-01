package presenter

import (
	"encoding/json"
	"log"
	"skillsbox-diploma/internal/domain/model"
)

type supportPresenter struct {
}

type SupportPresenter interface {
	ResponseSupport(*[]model.SupportData) []byte
}

func NewSupportPresenter() SupportPresenter {
	return &supportPresenter{}
}

func (sp *supportPresenter) ResponseSupport(sd *[]model.SupportData) []byte {
	responseSupport, err := json.Marshal(sd)
	if err != nil {
		log.Fatalln(err)
	}
	return responseSupport
}
