package ports

import "com.guilherme/hexagonal-example/internal/core/domain"

type GameRepository interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}

type GameService interface {
	Get(id string) (domain.Game, error)
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, now uint, col uint) (domain.Game, error)
}
