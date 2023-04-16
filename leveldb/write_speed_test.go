package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"strconv"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	valueSets := map[string][]byte{
		"empty":  []byte(""),
		"small":  randomSmall,
		"medium": randomMedium,
		"large":  randomLarge,
	}

	for name, value := range valueSets {
		b.Run(name, func(b *testing.B) {
			benchmarkWriteWithValue(b, value)
		})
	}
}

func BenchmarkBatchWrite(b *testing.B) {
	valueSets := map[string][]byte{
		"empty":  []byte(""),
		"small":  randomSmall,
		"medium": randomMedium,
		"large":  randomLarge,
	}

	for name, value := range valueSets {
		b.Run(name, func(b *testing.B) {
			benchmarkBatchWriteWithValue(b, value)
		})
	}
}

func benchmarkWriteWithValue(b *testing.B, value []byte) {
	Reset()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		prefix := strconv.Itoa(i) //to prevent same keys
		err := db.Put([]byte("key"+prefix), value, nil)
		if err != nil {
			panic(err)
		}
	}
}

func benchmarkBatchWriteWithValue(b *testing.B, value []byte) {
	Reset()

	b.ResetTimer()

	batch := new(leveldb.Batch)
	for i := 0; i < b.N; i++ {
		prefix := strconv.Itoa(i) //to prevent same keys
		batch.Put([]byte("key"+prefix), value)
	}

	err := db.Write(batch, nil)
	if err != nil {
		panic(err)
	}
}
