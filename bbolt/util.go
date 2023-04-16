package bbolt

import (
	"go.etcd.io/bbolt"
	"os"
	"strconv"
)

const (
	path       = "temp-bolt.db"
	bucketName = "test"
)

var (
	db *bbolt.DB
)

func Reset() {
	if db != nil {
		_ = db.Close()
	}

	//cleanup
	_ = os.RemoveAll(path)

	//open db
	var err error
	db, err = bbolt.Open(path, 0666, nil)
	if err != nil {
		panic(err)
	}

	//create bucket
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})

	if err != nil {
		panic(err)
	}
}

func ResetAndFill(count int, value []byte) {
	Reset()

	err := db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		for i := 0; i < count; i++ {
			suffix := strconv.Itoa(i) //to prevent same keys
			err := bucket.Put([]byte("key"+suffix), value)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func DeleteEveryNth(n int) {
	counter := 0
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		c := bucket.Cursor()

		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			if counter%n == 0 {
				err := bucket.Delete(k)
				if err != nil {
					return err
				}
			}

			counter++
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
}

func GetFilesystemSizeBytes() int64 {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	fi, err := file.Stat()
	if err != nil {
		panic(err)
	}

	return fi.Size()
}
