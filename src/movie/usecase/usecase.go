package usecase

import (
	"context"

	"github.com/yanuar-nc/go-boiler-plate/src/movie/domain"
)

// MovieUsecase interface
type MovieUsecase interface {
	Save(ctx context.Context, param domain.Movie) error
}
