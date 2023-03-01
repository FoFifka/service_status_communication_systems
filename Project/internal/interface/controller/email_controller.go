package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/usecase/interactor"
)

type emailController struct {
	EmailInteractor interactor.EmailInteractor
}

type EmailController interface {
	GetEmailData(ctx echo.Context) error
}

func NewEmailController(ei interactor.EmailInteractor) EmailController {
	return &emailController{ei}
}

func (ec *emailController) GetEmailData(ctx echo.Context) error {
	emailResponse, err := ec.EmailInteractor.Get()
	if err != nil {
		return err
	}
	return ctx.JSONBlob(http.StatusOK, emailResponse)
}
