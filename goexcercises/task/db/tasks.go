package db

import (
	"encoding/binary"
	"encoding/json"
	"time"

	bolt "go.etcd.io/bbolt"
)

var taskbucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key      int    `json:"-"`
	Value    string `json:"value"`
	Priority int    `json:"priority"`
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskbucket)
		return err
	})
}

func ClearTasks() error {
	err := db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(taskbucket)
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists(taskbucket)
		return err
	})
	return err
}

func CreateTask(task string, priority int) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
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

func DeleteTask(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskbucket)
		return bucket.Delete(itob(key))
	})
}

func AllTasks() ([]Task, error) {
	var tasks []Task
	err := db.View(func(tx *bolt.Tx) error {
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
