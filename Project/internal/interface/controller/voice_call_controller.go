package controller

import (
	"github.com/labstack/echo"
	"net/http"
	"skillsbox-diploma/internal/usecase/interactor"
)

type voiceCallController struct {
	VoiceCallInteractor interactor.VoiceCallInteractor
}

type VoiceCallController interface {
	GetVoiceCallData(ctx echo.Context) error
}

func NewVoiceCallController(vi interactor.VoiceCallInteractor) VoiceCallController {
	return &voiceCallController{vi}
}

func (vc *voiceCallController) GetVoiceCallData(ctx echo.Context) error {
	voiceCallResponse, err := vc.VoiceCallInteractor.Get()
	if err != nil {
		return err
	}
	return ctx.JSONBlob(http.StatusOK, voiceCallResponse)
}
