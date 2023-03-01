package interactor

import (
	"skillsbox-diploma/internal/usecase/presenter"
	"skillsbox-diploma/internal/usecase/repository"
)

type voiceCallInteractor struct {
	VoiceCallRepository repository.VoiceCallRepository
	VoiceCallPresenter  presenter.VoiceCallPresenter
}

type VoiceCallInteractor interface {
	Get() ([]byte, error)
}

func NewVoiceCallInteractor(r repository.VoiceCallRepository, p presenter.VoiceCallPresenter) VoiceCallInteractor {
	return &voiceCallInteractor{r, p}
}

func (vi *voiceCallInteractor) Get() ([]byte, error) {
	v, err := vi.VoiceCallRepository.GetAll()
	if err != nil {
		return nil, err
	}

	return vi.VoiceCallPresenter.ResponseVoiceCall(v), nil
}
