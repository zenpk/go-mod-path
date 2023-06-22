package gmp

import (
	"errors"
	"os"
)

var (
	FailedToFind = errors.New("failed to find the path of go.mod")
)

// GetNearestPath returns the nearest path that contains a `go.mod`
func GetNearestPath() (string, error) {
	return GetPath(0)
}

// GetPath let user specify the level of the `go.mod`
// e.g. level=1 means get the parent module's `go.mod` path
func GetPath(level int) (string, error) {
	levelCnt := -1
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	dir += "/"
	for true {
		files, err := os.ReadDir(dir)
		if err != nil {
			return "", err
		}
		if binarySearchGoMod(files) {
			levelCnt++
			if levelCnt == level {
				return dir, nil
			}
		}
		dir += "../"
	}
	return "", FailedToFind
}

// binarySearchGoMod files is sorted by filename, so it is possible to use binary search
func binarySearchGoMod(files []os.DirEntry) bool {
	goModName := "go.mod"
	for i, j := 0, len(files)-1; i <= j; {
		mid := i + (j-i)/2
		file := files[mid]
		if file.Name() == goModName {
			return true
		}
		if file.Name() > goModName {
			j = mid
		}
		if file.Name() < goModName {
			i = mid
		}
	}
	return false
}
