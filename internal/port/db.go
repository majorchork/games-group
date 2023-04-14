package port

import (
	"context"
	"github.com/majorchork/tech-crib-africa/internal/models"
)

type DB interface {
	CreateUser(ctx context.Context, userRequest models.UserRequest) error
	GetUserByEmail(ctx context.Context, email string) (*models.Admin, error)
	CreateGuests(ctx context.Context, guests []models.PeopleRequest) error
	GetGuestsByGroup(ctx context.Context, group int) (*[]models.Guest, error)
	GetGuestByEmail(ctx context.Context, email string) (*models.Guest, error)
	GetGuests(ctx context.Context) (*[]models.Guest, error)
}
