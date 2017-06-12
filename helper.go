package pickledb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// ReadFromJSONFile reads the json file given by path to `data`
func ReadFromJSONFile(path string) (data interface{}, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, err
	}

	return
}

// WriteToJSONFile writes the `data` the the given path as json file
func WriteToJSONFile(data interface{}, path string) error {
	js, _ := json.Marshal(data)

	file, err := os.OpenFile(path, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	ioutil.WriteFile(path, js, 0644)

	return nil
}

// IsFileExisted return true if the file given by `path` is existed, or
// return false
func IsFileExisted(path string) bool {
	info, err := os.Stat(path)

	if err != nil {
		return false
	}
	if info.IsDir() {
		return false
	}

	return true
}

// MakeDir make the directory given by `dir`
func MakeDir(dir string) error {
	err := os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	return nil
}
