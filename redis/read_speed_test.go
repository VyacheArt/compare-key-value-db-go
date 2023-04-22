package redis

import (
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
		_, err := db.Get(ctx, "key"+strconv.Itoa(i)).Result()
		if err != nil {
			panic(err)
		}
	}
}
