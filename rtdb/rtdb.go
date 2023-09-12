package realtimedb

import (
	"fmt"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	defaultDBName = "default"
)

type M map[string]string

type Collection struct {
	*bbolt.Bucket
}

type RealTimeDb struct {
	db *bbolt.DB
}

func New() (*RealTimeDb, error) {
	dbname := fmt.Sprintf("%s.realtimedb", defaultDBName)
	db, err := bbolt.Open(dbname, 0666, nil)
	if err != nil {
		return nil, err
	}
	return &RealTimeDb{
		db: db,
	}, nil
}

func (rtdb *RealTimeDb) CreateCollection(name string) (*Collection, error) {
	tx, err := rtdb.db.Begin(true)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// bucket := tx.Bucket([]byte(name))
	// if bucket != nil {
	// 	return &Collection{Bucket: bucket}, nil
	// }
	bucket, err := tx.CreateBucketIfNotExists([]byte(name))
	if err != nil {
		return nil, err
	}
	return &Collection{Bucket: bucket}, nil
}

// func (rtdb *RealTimeDb) Insert(collec *Collection, k, v any) // for `any` supported types
func (rtdb *RealTimeDb) Insert(collecName string, data M) (uuid.UUID, error) {
	id := uuid.New()
	tx, err := rtdb.db.Begin(true)
	if err != nil {
		return id, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte(collecName))
	if err != nil {
		return id, err
	}
	for k, v := range data {
		if err := bucket.Put([]byte(k), []byte(v)); err != nil {
			return id, err
		}
	}
	if err := bucket.Put([]byte("id"), []byte(id.String())); err != nil {
		return id, err
	}
	return id, tx.Commit()
}

// get http://localhost:7777/users?eq.name=(sumit@gmail.com) ----> general case
// get http://localhost:7777/users?eq.id=23423rbe8w449-823y78w-uwehf7643 ----> in our case
// func (rtdb *RealTimeDb) Select(collec string, query M) (M, error) {
// 	tx, err := rtdb.db.Begin(false)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bucket := tx.Bucket([]byte(collec))
// 	if bucket == nil {
// 		return nil, fmt.Errorf("Collection (%s) not found", collec)
// 	}
// }
