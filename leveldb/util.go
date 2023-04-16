package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"os"
	"path/filepath"
	"strconv"
)

const (
	directoryPath = "temp-leveldb"
)

var (
	db *leveldb.DB
)

func Reset() {
	if db != nil {
		_ = db.Close()
	}

	//cleanup
	err := os.RemoveAll(directoryPath)
	if err != nil {
		panic(err)
	}

	//open db
	db, err = leveldb.OpenFile(directoryPath, nil)
	if err != nil {
		panic(err)
	}
}

func ResetAndFill(count int, value []byte) {
	Reset()

	batch := new(leveldb.Batch)
	for i := 0; i < count; i++ {
		batch.Put([]byte("key"+strconv.Itoa(i)), value)
	}

	err := db.Write(batch, nil)
	if err != nil {
		panic(err)
	}
}

func DeleteEveryNth(n int) {
	iter := db.NewIterator(nil, nil)
	defer iter.Release()

	batch := new(leveldb.Batch)
	for i := 0; iter.Next(); i++ {
		if i%n == 0 {
			batch.Delete(iter.Key())
		}
	}

	err := db.Write(batch, nil)
	if err != nil {
		panic(err)
	}
}

func GetFilesystemSizeBytes() int64 {
	var size int64
	_ = filepath.Walk(directoryPath, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		size += info.Size()
		return nil
	})

	return size
}
