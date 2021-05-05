package watchman

import "strings"

// knownConfigs is a list of files to ignore even if they are not in .gitignore
var knownConfigs = []string{".git", ".idea", ".cache"}

// isKnownConfigs returns true if given file is supposed to be ignored
func isKnownConfigs(needle string) bool {
	for _, val := range knownConfigs {
		if val == needle {
			return true
		}
	}
	return false
}

// childOfIgnoredDirectory returns true for children of ignored directory
func childOfIgnoredDirectory(path string) bool {
	fragments := strings.Split(path,"/")
	if len(fragments) >= 1 {
		return isKnownConfigs(fragments[0])
	}

	return false
}
