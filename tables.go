package tinyindexdb

import (
	"errors"
	"os"
	"path/filepath"
)

func doesTableExist(table_name string) bool {
	initCache()

	table_name = sanitiseString(table_name)
	tblPath := filepath.Join(getCachePath(), ".tables", table_name)

	if stat, err := os.Stat(tblPath); err == nil && stat.IsDir() {
		return true
	}

	return false
}

func NewTable(table_name string) error {
	initCache()

	table_name = sanitiseString(table_name)
	tblPath := filepath.Join(getCachePath(), ".tables", table_name)

	if _, err := os.Stat(tblPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(tblPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func DestroyTable(table_name string) error {
	initCache()

	table_name = sanitiseString(table_name)
	tblPath := filepath.Join(getCachePath(), ".tables", table_name)

	err := os.RemoveAll(tblPath)
	if err != nil {
		return err
	}

	return nil
}

func TableExists(table_name string) bool {
	return doesTableExist(table_name)
}
