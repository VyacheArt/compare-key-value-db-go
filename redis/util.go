package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"strconv"
	"strings"
)

// const addr = "localhost:6379"
const addr = "/opt/homebrew/opt/redis/redis.sock"

var (
	db  *redis.Client
	ctx = context.Background()
)

func Reset() {
	if db != nil {
		_ = db.Close()
	}

	//open db
	db = redis.NewClient(&redis.Options{
		Network: "unix",
		Addr:    addr,
		DB:      0,
	})

	//ping
	_, err := db.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	//flush
	_, err = db.FlushDB(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func ResetAndFill(count int, value []byte) {
	Reset()

	pipe := db.Pipeline()
	for i := 0; i < count; i++ {
		pipe.Set(ctx, "key"+strconv.Itoa(i), value, 0)
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func DeleteEveryNth(n int) {
	pipe := db.Pipeline()
	for i := 0; i < n; i++ {
		pipe.Del(ctx, "key"+strconv.Itoa(i))
	}

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}
}

func GetUsedMemoryBytes() int64 {
	info, err := db.Info(ctx, "memory").Result()
	if err != nil {
		panic(err)
	}

	return parseInfo(info, "used_memory:")
}

func parseInfo(info string, key string) int64 {
	lines := strings.Split(info, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, key) {
			value, err := strconv.ParseInt(strings.Trim(line[len(key):], "\r"), 10, 64)
			if err != nil {
				panic(err)
			}

			return value
		}
	}

	return 0
}
