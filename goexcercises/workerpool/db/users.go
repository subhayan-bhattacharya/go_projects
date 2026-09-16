package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	bolt "go.etcd.io/bbolt"
)

var userbucket = []byte("users")

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Username  string
	Email     string
	FirstName string
	LastName  string
}

// Repository defines the interface for task storage operations
type Repository interface {
	AddUser(user User) error
	DeleteUser(username string) error
	GetUser(username string) (User, error)
}

type BoltRepository struct {
	db *bolt.DB
}

// NewBoltRepository creates a new BoltRepository and initializes the database
func NewBoltRepository(dbPath string) (*BoltRepository, error) {
	boltDB, err := bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}

	repo := &BoltRepository{db: boltDB}
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(userbucket)
		return err
	})
	if err != nil {
		boltDB.Close()
		return nil, err
	}

	return repo, nil
}

// Close closes the database connection
func (r *BoltRepository) Close() error {
	return r.db.Close()
}

func (r *BoltRepository) AddUser(user User) error {
	usernameKey := []byte(user.Username)
	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("could not encode data into byte %v", err)
	}
	return r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(userbucket)
		return b.Put(usernameKey, data)
	})
}

func (r *BoltRepository) GetUser(username string) (User, error) {
	usernameKey := []byte(username)
	var user User
	err := r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(userbucket)
		userDetails := b.Get(usernameKey)
		if userDetails == nil {
			return ErrUserNotFound
		}
		return json.Unmarshal(userDetails, &user)
	})
	return user, err
}

func (r *BoltRepository) DeleteUser(username string) error {
	usernameKey := []byte(username)
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(userbucket)
		return bucket.Delete(usernameKey)
	})
}
