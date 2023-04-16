package badger

import (
	"github.com/dgraph-io/badger/v4"
	"os"
	"path/filepath"
	"strconv"
)

const (
	directoryPath = "temp-badger"
)

var (
	db *badger.DB
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
	opts := badger.DefaultOptions(directoryPath).
		WithValueLogFileSize(100 * 1024 * 1024)

	//disable logging
	opts.Logger = nil

	db, err = badger.Open(opts)
	if err != nil {
		panic(err)
	}
}

func ResetAndFill(count int, value []byte) {
	Reset()

	batch := db.NewWriteBatch()
	defer batch.Cancel()

	for i := 0; i < count; i++ {
		batch.Set([]byte("key"+strconv.Itoa(i)), value)
	}

	err := batch.Flush()
	if err != nil {
		panic(err)
	}
}

func DeleteEveryNth(n int) {
	err := db.Update(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		count := 0

		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			if count%n == 0 {
				err := txn.Delete(item.Key())
				if err != nil {
					return err
				}
			}

			count++
		}

		return nil
	})

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
