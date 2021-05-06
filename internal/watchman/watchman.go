package watchman

import (
	"fmt"
	"github.com/YashKumarVerma/hentry-client/internal/fs"
	"github.com/sergi/go-diff/diffmatchpatch"
	"os"
	"path/filepath"
	"time"
)

// Watchman is responsible for the following actions
// - walking the directory tree starting from the current working directory.
// - index all files along with their metadata and save into in-memory cache
// - sync data with cloud provider
// - save last snapshot of cache locally

// IndexAllFiles walks the passed path, and returns the list of ignored and
// focused files
func IndexAllFiles(path string) (filesNotIgnored []string, filesIgnored []string) {
	var focusedFiles []string
	var ignoredFiles []string

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if isKnownConfigs(info.Name()) {
			ignoredFiles = append(ignoredFiles, path)
		} else if childOfIgnoredDirectory(path) {
			ignoredFiles = append(ignoredFiles, path)
		} else if info.IsDir() {
			ignoredFiles = append(ignoredFiles, path)
		} else {
			focusedFiles = append(focusedFiles, path)
		}
		return nil
	})

	return focusedFiles, ignoredFiles
}

// ProcessFile calculates the diff between subsequent calls
func ProcessFile(file fs.FileDetails) bool {
	_, newDetails := fs.AnalyzeFile(file.Path)
	var difference Diff
	diffCalculator := diffmatchpatch.New()

	// check if entry exist in hashmap
	if oldDetails, ok := index[file.Path]; ok {
		difference.size = oldDetails.Size - newDetails.Size
		difference.timestamp = newDetails.LastModified.Sub(oldDetails.LastModified)
		difference.changes = diffCalculator.DiffMain(oldDetails.Contents, newDetails.Contents, true)
	} else {
		difference.size = newDetails.Size
		difference.timestamp = newDetails.LastModified.Sub(time.Now())
		difference.changes = diffCalculator.DiffMain(newDetails.Contents, "", true)
	}

	fmt.Println(difference)
	return true
}
