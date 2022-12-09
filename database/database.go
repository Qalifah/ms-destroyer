package database

import (
	"context"
	"destroyer"
)

type Repository interface {
	ListTarget(ctx context.Context) ([]destroyer.Target, error)
	GetTarget(ctx context.Context, id string) (*destroyer.Target, error)
}
