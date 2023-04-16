package leveldb

import (
	"github.com/VyacheArt/compare-key-value-db-go/util"
	"os"
	"testing"
)

var (
	randomSmall  = util.RandomBytes(10)
	randomMedium = util.RandomBytes(100)
	randomLarge  = util.RandomBytes(1000)
)

func TestMain(m *testing.M) {
	status := m.Run()

	if db != nil {
		_ = db.Close()
	}

	os.Exit(status)
}
