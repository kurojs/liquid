package store

import (
	"context"
)

type Store interface {
	ValidateUser(ctx context.Context, user *User) (bool, error)
}
