package watchman

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"time"
)

// Diff contains the difference between cached data
type Diff struct {
	size int64
	timestamp time.Duration
	changes []diffmatchpatch.Diff
}