package bbolt

import (
	"go.etcd.io/bbolt"
	"strconv"
	"testing"
)

func BenchmarkWriteSeparatedTransactions(b *testing.B) {
	valueSets := map[string][]byte{
		"empty":  []byte(""),
		"small":  randomSmall,
		"medium": randomMedium,
		"large":  randomLarge,
	}

	for name, value := range valueSets {
		b.Run(name, func(b *testing.B) {
			benchmarkWriteSeparatedTransactionsWithValue(b, value)
		})
	}
}

func BenchmarkWriteSingleTransaction(b *testing.B) {
	valueSets := map[string][]byte{
		"empty":  []byte(""),
		"small":  randomSmall,
		"medium": randomMedium,
		"large":  randomLarge,
	}

	for name, value := range valueSets {
		b.Run(name, func(b *testing.B) {
			benchmarkWriteSingleTransactionWithValue(b, value)
		})
	}
}

func benchmarkWriteSeparatedTransactionsWithValue(b *testing.B, value []byte) {
	Reset()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.Update(func(tx *bbolt.Tx) error {
			bucket := tx.Bucket([]byte(bucketName))
			suffix := strconv.Itoa(i) //to prevent same keys
			return bucket.Put([]byte("key"+suffix), value)
		})

		if err != nil {
			panic(err)
		}
	}
}

func benchmarkWriteSingleTransactionWithValue(b *testing.B, value []byte) {
	Reset()

	b.ResetTimer()

	err := db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		for i := 0; i < b.N; i++ {
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
