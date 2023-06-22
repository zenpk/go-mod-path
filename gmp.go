package gmp

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrFailedToFind = errors.New("failed to find the path of go.mod")
)

// GetNearestPath returns the nearest path that contains a `go.mod`
func GetNearestPath() (string, error) {
	return GetPath(0)
}

// GetPath returns a path that contains a `go.mod`, allowing user to specify the level of the `go.mod`
// e.g. level=1 means get the parent module's `go.mod` path
func GetPath(level int) (string, error) {
	levelCnt := -1
	dir, err := getDir()
	if err != nil {
		return "", err
	}
	for strings.Contains(dir, "/") {
		find, err := binarySearchGoMod(dir)
		if err != nil {
			return "", err
		}
		if find {
			levelCnt++
			if levelCnt == level {
				return dir, nil
			}
		}
		dir, err = truncateLastPath(dir)
		if err != nil {
			return "", err
		}
	}
	return "", ErrFailedToFind
}

// GetFolderPath allows user to find the go-mod-path by the folder name
func GetFolderPath(name string) (string, error) {
	name += "/"
	dir, err := getDir()
	if err != nil {
		return "", err
	}
	for strings.Contains(dir, "/") {
		if strings.HasSuffix(dir, name) {
			find, err := binarySearchGoMod(dir)
			if err != nil {
				return "", err
			}
			if find {
				return dir, nil
			}
		}
		dir, err = truncateLastPath(dir)
		if err != nil {
			return "", err
		}
	}
	return "", ErrFailedToFind
}

func getDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	// Windows
	dir = strings.ReplaceAll(dir, "\\", "/")
	dir += "/"
	return dir, nil
}

// binarySearchGoMod binary search for `go.mod` in current directory
func binarySearchGoMod(dir string) (bool, error) {
	// the files returned by os.ReadDir is sorted by the filename,
	// so it is possible to use binary search
	files, err := os.ReadDir(dir)
	if err != nil {
		return false, err
	}
	goModName := "go.mod"
	for i, j := 0, len(files); i < j; {
		mid := i + (j-i)/2
		file := files[mid]
		if file.Name() == goModName {
			return true, nil
		}
		if file.Name() > goModName {
			j = mid
		}
		if file.Name() < goModName {
			i = mid + 1
		}
	}
	return false, nil
}

// truncateLastPath equals cd ../
func truncateLastPath(dir string) (string, error) {
	if len(dir) == 1 { // dir == "/"
		return "", ErrFailedToFind
	}
	// remove the last "/"
	dir = dir[:len(dir)-1]
	lastIndex := strings.LastIndex(dir, "/")
	return dir[:lastIndex], nil
}
