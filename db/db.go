package db

import (
	"github.com/boltdb/bolt"
	"github.com/cumedang/go/utils"
)

var db *bolt.DB

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocskBucket = "blocks"
	checkpoint   = "checkpoint"
)

func DB() *bolt.DB {
	if db == nil {
		dbPorinter, err := bolt.Open(dbName, 0600, nil)
		db = dbPorinter
		utils.HandelERror(err)
		err = db.Update(func(t *bolt.Tx) error {
			_, err := t.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandelERror(err)
			_, err = t.CreateBucketIfNotExists([]byte(blocskBucket))
			utils.HandelERror(err)
			return err
		})
		utils.HandelERror(err)

	}
	return db
}

func SaveBlock(hash string, data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocskBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandelERror(err)

}

func Close() {
	DB().Close()
}

func SaveCheckPoint(data []byte) {
	err := DB().Update(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandelERror(err)
}

func Checkpoint() []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(dataBucket))
		data = bucket.Get([]byte(checkpoint))
		return nil
	})
	return data
}

func Block(hash string) []byte {
	var data []byte
	DB().View(func(t *bolt.Tx) error {
		bucket := t.Bucket([]byte(blocskBucket))
		data = bucket.Get([]byte(hash))
		return nil
	})
	return data
}
