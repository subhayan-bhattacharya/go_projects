package commands

import (
	domain "command-bus/domain"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type CreateUser struct {
	Marker[domain.UserId]
	Name  string
	Email string
}

type CreateUserHandler struct {
	Repository domain.UserRepository
	Now        func() time.Time
}

func (c CreateUserHandler) Handle(ctx context.Context, cmd CreateUser) (domain.UserId, error) {
	var zeroUserId domain.UserId
	exists, err := c.Repository.ExistsByEmail(ctx, cmd.Email)
	if err != nil {
		return zeroUserId, errors.New("could not check email")
	}
	if exists {
		return zeroUserId, errors.New("this email exists in the repository")
	}
	userId := fmt.Sprintf("%x", rand.Uint64())
	user := domain.User{
		UserId: domain.UserId(userId),
		Name:   cmd.Name,
		Email:  cmd.Email,
	}
	err = c.Repository.Save(ctx, user)
	if err != nil {
		return zeroUserId, errors.New("could not save user")
	}
	return user.UserId, nil
}
