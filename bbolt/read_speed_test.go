package bbolt

import (
	"go.etcd.io/bbolt"
	"strconv"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	valueSets := map[string][]byte{
		"empty":  []byte(""),
		"small":  randomSmall,
		"medium": randomMedium,
		"large":  randomLarge,
	}

	for name, value := range valueSets {
		b.Run(name, func(b *testing.B) {
			benchmarkReadWithValue(b, value)
		})
	}
}

func benchmarkReadWithValue(b *testing.B, value []byte) {
	ResetAndFill(b.N, value)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := db.View(func(tx *bbolt.Tx) error {
			bucket := tx.Bucket([]byte(bucketName))
			suffix := strconv.Itoa(i)
			_ = bucket.Get([]byte("key" + suffix))

			return nil
		})

		if err != nil {
			panic(err)
		}
	}
}
