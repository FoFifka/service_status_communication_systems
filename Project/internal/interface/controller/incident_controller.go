package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/usecase/interactor"
)

type incidentController struct {
	IncidentInteractor interactor.IncidentInteractor
}

type IncidentController interface {
	GetIncidentData(ctx echo.Context) error
}

func NewIncidentController(ii interactor.IncidentInteractor) IncidentController {
	return &incidentController{ii}
}

func (ic *incidentController) GetIncidentData(ctx echo.Context) error {
	i, err := ic.IncidentInteractor.Get()
	if err != nil {
		return err
	}

	return ctx.JSONBlob(http.StatusOK, i)
}
