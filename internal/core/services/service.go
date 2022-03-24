package services

import (
	"errors"

	"com.guilherme/hexagonal-example/internal/core/domain"
	"com.guilherme/hexagonal-example/internal/core/ports"
	"com.guilherme/hexagonal-example/pkg/uidgen"
)

type GameService struct {
	gamesRepository ports.GameRepository
	uidGen          uidgen.UIDGen
}

func New(gamesRepository ports.GameRepository) *GameService {
	return &GameService{
		gamesRepository: gamesRepository,
	}
}

func (service *GameService) Get(id string) (domain.Game, error) {
	game, err := service.gamesRepository.Get(id)

	if err != nil {
		return domain.Game{}, errors.New("Get game from repository has failed")
	}

	return game, nil
}

func (service *GameService) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs > size*size {
		return domain.Game{}, errors.New("Invalid amount of bombs")
	}

	game := domain.NewGame(service.uidGen.New(), name, size, bombs)

	if err := service.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New("Unable to create game")
	}

	return game, nil
}

func (service *GameService) Reveal(id string, now uint, col uint) (domain.Game, error) {
	return domain.Game{}, nil
}
