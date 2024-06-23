package tinyindexdb

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/danjacques/gofslock/fslock"
)

func Query(table_name, index_name string) (map[string]interface{}, error) {
	initCache()

	table_name = sanitiseString(table_name)
	index_name = sanitiseString(index_name)

	if !doesTableExist(table_name) {
		return nil, errors.New("the specified table does not exist")
	}

	indexPath := filepath.Join(getCachePath(), ".tables", table_name, index_name)

	if _, err := os.Stat(indexPath); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("the specified record does not exist")
	}

	record_content, err := os.ReadFile(indexPath)
	if err != nil {
		return nil, err
	}

	if !json.Valid(record_content) {
		return nil, errors.New("corrupted record data")
	}

	var map_ map[string]interface{}

	if err := json.Unmarshal(record_content, &map_); err != nil {
		return nil, err
	}

	return map_, nil
}

func Write(table_name, index_name string, data map[string]interface{}) error {
	initCache()

	table_name = sanitiseString(table_name)
	index_name = sanitiseString(index_name)

	if !doesTableExist(table_name) {
		return errors.New("the specified table does not exist")
	}

	indexPath := filepath.Join(getCachePath(), ".tables", table_name, index_name)

	// self note: fslock.WithBlocking creates the file if it does not exist
	return fslock.WithBlocking(indexPath, func() error {
		// the file is locked, keep retrying each 10 ms till the file is unlocked
		time.Sleep(10 * time.Millisecond)

		return nil
	}, func() error {
		f, err := os.OpenFile(indexPath, os.O_CREATE|os.O_WRONLY, 0777)
		if err == nil {
			defer f.Close()

			jsonEnc, err := json.Marshal(data)
			if err != nil {
				return err
			}

			_, err = f.WriteString(string(jsonEnc))
			if err != nil {
				return err
			}
		} else {
			return err
		}

		return nil
	})
}

func DeleteRecord(table_name, index_name string) error {
	initCache()

	table_name = sanitiseString(table_name)
	index_name = sanitiseString(index_name)

	if !doesTableExist(table_name) {
		return errors.New("the specified table does not exist")
	}

	indexPath := filepath.Join(getCachePath(), ".tables", table_name, index_name)

	if _, err := os.Stat(indexPath); errors.Is(err, os.ErrNotExist) {
		return errors.New("the specified record does not exist")
	}

	err := os.RemoveAll(indexPath)
	if err != nil {
		return err
	}

	return nil
}
