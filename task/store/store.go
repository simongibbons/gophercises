package store

import (
	"encoding/binary"
	"encoding/json"
	"github.com/coreos/bbolt"
)

type Store struct {
	Db *bolt.DB
}

var taskBucket = []byte("tasks")

func New(filepath string) (*Store, error) {
	db, err := bolt.Open(filepath, 0600, nil)
	if err != nil {
		return nil, err
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})

	return &Store{db}, nil
}

func (s *Store) Close() {
	s.Db.Close()
}

func (s *Store) AddTask(t Task) (*Task, error) {
	err := s.updateOrCreateTask(t)

	if err != nil {
		return nil, err
	}

	return &t, err
}

func (s *Store) updateOrCreateTask(t Task) error {
	return s.Db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		if bucket == nil {
			panic("Bucket not initialised")
		}

		if t.Id == nil {
			id, _ := bucket.NextSequence()
			intId := int(id)
			t.Id = &intId
		}

		j, err := json.Marshal(t)
		if err != nil {
			return err
		}

		bucket.Put(itob(*t.Id), j)
		return nil
	})
}


func (s *Store) GetTasks() (ts []Task, err error) {
	err = s.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		if bucket == nil {
			panic("Bucket not initialised")
		}

		return bucket.ForEach(func(k, v []byte) error {
			t := &Task{}
			err := json.Unmarshal(v, t)
			if err != nil {
				return err
			}
			ts = append(ts, *t)
			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	return ts, nil
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

type Task struct {
	Content  string `json:"content"`
	Id       *int    `json:"id"`
	Complete bool   `json:"complete"`
}
