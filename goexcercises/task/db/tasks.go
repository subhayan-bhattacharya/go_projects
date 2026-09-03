package db

import (
	"encoding/binary"
	"time"

	bolt "go.etcd.io/bbolt"
)

var taskbucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
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

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskbucket)
		idx, _ := bucket.NextSequence()
		id = int(idx)
		key := itob(id)
		return bucket.Put(key, []byte(task))
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
			task := Task{
				Key:   btoi(k),
				Value: string(v),
			}
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
