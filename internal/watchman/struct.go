package watchman

import (
	"time"
)

// Diff contains the difference between cached data
type Diff struct {
	size      int64
	timestamp time.Duration
	changes   int
}
