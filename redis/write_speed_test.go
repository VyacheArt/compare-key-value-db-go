package redis

import (
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
		err := db.Set(ctx, "key"+strconv.Itoa(i), value, 0).Err()
		if err != nil {
			panic(err)
		}
	}
}

func benchmarkWriteSingleTransactionWithValue(b *testing.B, value []byte) {
	Reset()

	b.ResetTimer()

	pipe := db.Pipeline()
	for i := 0; i < b.N; i++ {
		pipe.Set(ctx, "key"+strconv.Itoa(i), value, 0)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
}
