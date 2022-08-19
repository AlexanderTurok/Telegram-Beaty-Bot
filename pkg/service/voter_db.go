package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
)

type VoterDBService struct {
	repository *repository.Repository
}

func NewVoterDBService(repository *repository.Repository) *VoterDBService {
	return &VoterDBService{
		repository: repository,
	}
}

func (v *VoterDBService) GetParticipant(uuid int) (*telegram.Participant, error) {
	return nil, nil
}

func (v *VoterDBService) UpdateParticipant(uuid int) error {
	return nil
}
