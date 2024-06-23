package tinyindexdb

import (
	"errors"
	"os"
	"path/filepath"
)

func getCachePath() string {
	path, _ := filepath.Abs("./")
	cachePath := filepath.Join(path, ".tinyindexdb_cache")

	return cachePath
}

func doesCacheExist() bool {
	path := getCachePath()

	if stat, err := os.Stat(path); err == nil && stat.IsDir() {
		return true
	}

	return false
}

func initCache() bool {
	path := getCachePath()

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return false
		}
	}

	return true
}

func destroyCache() bool {
	if !doesCacheExist() {
		return true
	}

	path := getCachePath()

	err := os.RemoveAll(path)

	return err == nil
}

func ClearCache() {
	destroyCache()
	initCache()
}
