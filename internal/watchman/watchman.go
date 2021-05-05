package watchman

import (
	"fmt"
	"os"
	"path/filepath"
)

// Watchman is responsible for the following actions
// - walking the directory tree starting from the current working directory.
// - index all files along with their metadata and save into in-memory cache
// - sync data with cloud provider
// - save last snapshot of cache locally

// AppFs allows working with different file systems using a clean API, so less
// worries about platform compatibility.

func Init() {

}

func IndexAllFiles(path string) {

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if isKnownConfigs(info.Name()) {
			fmt.Println(">> directly ignore : " + path)
		}

		if childOfIgnoredDirectory(path) {
			fmt.Println(">> ignore as child : " + path)
		} else {
			fmt.Println(">> >> >> >>  accept as child : " + path)
		}

		return nil
	})
}
