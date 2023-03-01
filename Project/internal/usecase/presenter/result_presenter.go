package presenter

import "skillsbox-diploma/internal/domain/model"

type ResultPresenter interface {
	ResponseResult(model.ResultT) []byte
}
