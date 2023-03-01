package presenter

import (
	"encoding/json"
	"log"
	"skillsbox-diploma/internal/domain/model"
)

type resultPresenter struct {
}

type ResultPresenter interface {
	ResponseResult(model.ResultT) []byte
}

func NewResultPresenter() ResultPresenter {
	return &resultPresenter{}
}

func (rp *resultPresenter) ResponseResult(r model.ResultT) []byte {
	responseResult, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return responseResult
}
