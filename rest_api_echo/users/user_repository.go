package users

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository
type IRepository interface {
	GetSth() (string, error)
}

type UserRepository struct {
	client *mongo.Client
}

func (r *UserRepository) GetSth() (string, error) {
    // Get data from MongoDB
	return "TODO next", fmt.Errorf("TODO next")
}

func NewUserRepository( client *mongo.Client ) UserRepository {
	return UserRepository{client: client}
}