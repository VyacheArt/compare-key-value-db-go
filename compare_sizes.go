package main

import (
	"fmt"
	"github.com/VyacheArt/compare-key-value-db-go/badger"
	"github.com/VyacheArt/compare-key-value-db-go/bbolt"
	"github.com/VyacheArt/compare-key-value-db-go/leveldb"
	"github.com/VyacheArt/compare-key-value-db-go/redis"
	"github.com/VyacheArt/compare-key-value-db-go/util"
	"github.com/fatih/color"
)

type (
	resetAndFillFunc           func(count int, value []byte)
	deleteEveryNthFunc         func(n int)
	getFilesystemSizeBytesFunc func() int64
)

var counts = []int{100, 1000, 10_000, 100_000}

func main() {
	valueSets := [][]byte{
		util.RandomBytes(10),
		util.RandomBytes(100),
		util.RandomBytes(1000),
	}

	for _, valueSet := range valueSets {
		color.Cyan("=====================================")
		color.Cyan("Compare with value size %d", len(valueSet))
		color.Cyan("=====================================")

		//run comparing for bbolt
		runComparing("bbolt", valueSet, bbolt.ResetAndFill, bbolt.DeleteEveryNth, bbolt.GetFilesystemSizeBytes)

		//run comparing for leveldb
		runComparing("leveldb", valueSet, leveldb.ResetAndFill, leveldb.DeleteEveryNth, leveldb.GetFilesystemSizeBytes)

		//run comparing for badger
		runComparing("badger", valueSet, badger.ResetAndFill, badger.DeleteEveryNth, badger.GetFilesystemSizeBytes)

		//run comparing for redis
		runComparing("redis", valueSet, redis.ResetAndFill, redis.DeleteEveryNth, redis.GetUsedMemoryBytes)
	}
}

func runComparing(libName string, fillValue []byte, resetAndFill resetAndFillFunc, deleteEveryNth deleteEveryNthFunc, getFilesystemSizeBytes getFilesystemSizeBytesFunc) {
	color.Green("=====================================")
	color.Green("Compare %s", libName)
	color.Green("=====================================")

	color.Blue("Fill")

	for _, count := range counts {
		fmt.Println("Fill with", count, "items")
		resetAndFill(count, fillValue)

		color.Red("Size before 0 KB, after %d KB", getFilesystemSizeBytes()/1024)
	}

	const deleteEveryN = 10
	color.Blue("Delete of every %d items", deleteEveryN)

	for _, count := range counts {
		resetAndFill(count, fillValue)
		size := getFilesystemSizeBytes() / 1024

		fmt.Println("Delete every", deleteEveryN, "items from", count, "items")
		deleteEveryNth(deleteEveryN)

		color.Red("Size before %d KB, after %d KB", size, getFilesystemSizeBytes()/1024)
	}
}
