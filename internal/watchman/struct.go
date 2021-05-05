package watchman

// Node represents a file on disk, with few attributes that assist in tracking changes
type Node struct {
	filename     string
	absolutePath string
	lastEdited   int64
	lastParsed   int64
	previousSize int
	
}
