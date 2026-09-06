package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	bolt "go.etcd.io/bbolt"
)

var taskbucket = []byte("tasks")

type Task struct {
	Key      int    `json:"-"`
	Value    string `json:"value"`
	Priority int    `json:"priority"`
}

// Repository defines the interface for task storage operations
type Repository interface {
	ClearTasks() error
	CreateTask(task string, priority int) (int, error)
	DeleteTask(key int) error
	AllTasks() ([]Task, error)
	Close() error
}

// BoltRepository implements Repository using BoltDB
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
		_, err := tx.CreateBucketIfNotExists(taskbucket)
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

// ClearTasks removes all tasks from the database
func (r *BoltRepository) ClearTasks() error {
	err := r.db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(taskbucket)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(taskbucket)
		return err
	})
	return err
}

// CreateTask adds a new task to the database
func (r *BoltRepository) CreateTask(task string, priority int) (int, error) {
	var id int
	err := r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskbucket)
		idx, _ := bucket.NextSequence()
		id = int(idx)

		value, err := json.Marshal(Task{Value: task, Priority: priority})
		if err != nil {
			return err
		}
		return bucket.Put(itob(id), value)
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

// DeleteTask removes a task from the database
func (r *BoltRepository) DeleteTask(key int) error {
	return r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskbucket)
		return bucket.Delete(itob(key))
	})
}

// AllTasks retrieves all tasks from the database
func (r *BoltRepository) AllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskbucket)
		cursor := bucket.Cursor()
		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			var task Task
			if err := json.Unmarshal(v, &task); err != nil {
				return err
			}
			task.Key = btoi(k)
			tasks = append(tasks, task)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
