package presenter

import (
	"encoding/json"
	"log"
	"skillsbox-diploma/internal/domain/model"
)

type mmsPresenter struct {
}

type MMSPresenter interface {
	ResponseMMS(*[]model.MMSData) []byte
}

func NewMMSPresenter() MMSPresenter {
	return &mmsPresenter{}
}

func (mp *mmsPresenter) ResponseMMS(md *[]model.MMSData) []byte {

	mmsData := *md

	mmsResponse, err := json.Marshal(mmsData)
	if err != nil {
		log.Fatalln(err)
	}
	return mmsResponse
}
