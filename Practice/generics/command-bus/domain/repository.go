package domain

import "context"

var Database = map[string]User{}

type UserRepository interface {
	ExistsByEmail(ctx context.Context, email string) (bool, error)
	Save(ctx context.Context, u User) error
}

type Repository struct{}

func (r *Repository) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	err := ctx.Err()
	if err != nil {
		return false, err
	}
	_, ok := Database[email]
	if ok {
		return true, nil
	}
	return false, nil
}

func (r *Repository) Save(ctx context.Context, user User) error {
	err := ctx.Err()
	if err != nil {
		return err
	}
	Database[user.Email] = user
	return nil
}
