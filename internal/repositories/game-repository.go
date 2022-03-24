package repository

import (
	"encoding/json"
	"errors"

	"com.guilherme/hexagonal-example/internal/core/domain"
)

type InMemoryRepository struct {
	data map[string][]byte
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: map[string][]byte{}}
}

func (repository InMemoryRepository) Get(id string) (domain.Game, error) {
	if value, ok := repository.data[id]; ok {
		game := domain.Game{}
		if err := json.Unmarshal(value, &game); err != nil {
			return domain.Game{}, errors.New("Fail to get value from key value store")
		}

		return game, nil
	}

	return domain.Game{}, errors.New("Game not found in key value store")
}

func (repository InMemoryRepository) Save(domain.Game) error {
	return nil
}
