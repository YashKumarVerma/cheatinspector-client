package watchman

import (
	"os"
	"path/filepath"
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
		} else {
			focusedFiles = append(focusedFiles, path)
		}
		return nil
	})

	return focusedFiles, ignoredFiles
}
