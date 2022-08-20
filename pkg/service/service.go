package service

import (
	"github.com/AlexanderTurok/telegram-beaty-bot"
	"github.com/AlexanderTurok/telegram-beaty-bot/pkg/repository"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Participant interface {
	SetParticipantName(message *tgbotapi.Message) error

	IsParticipant(uuid int) (bool, error)
	GetParticipant(uuid int) (*telegram.Participant, error)
	GetAllParticipants() (*[]telegram.Participant, error)
	AddParticipant(uuid int) error
	UpdateParticipant(column, value string, uuid int) error
	DeleteParticipant(uuid int) error
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
}

type Voter interface {
	GetParticipant(uuid int) (*telegram.Participant, error)
	UpdateParticipant(uuid int) error
	GetCache(uuid int) (string, error)
	SetCache(uuid int, value string) error
	DeleteCache(uuid int) error
}

type Service struct {
	Participant
	Voter
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Participant: NewParticipantService(repository),
		Voter:       NewVoterService(repository),
	}
}
