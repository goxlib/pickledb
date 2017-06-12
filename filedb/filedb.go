package filedb

import (
	"os"
	"path/filepath"

	"pickledb"
)

var dbDir string

// SetDBDir sets db file store dir
func SetDBDir(dir string) {
	dbDir = dir
}

// FileDB wrap pickledb and provide unified file based database operations
type FileDB struct {
	Path string
	DB   *pickledb.PickleDb
}

// New returns a FileDB instance with the given path
func New(path string, abs bool) *FileDB {
	if !abs {
		path = filepath.Join(dbDir, path)
	}

	// is dbDir exists?
	err := os.MkdirAll(dbDir, 0664)
	if err != nil {
		panic(err)
	}

	db := pickledb.New(path, "")
	db.Load()

	return &FileDB{
		Path: path,
		DB:   db,
	}
}

// Get returns the value by the given key
func (f *FileDB) Get(key string) interface{} {
	return f.DB.Get(key)
}

// Set sets the key-value into the database
func (f *FileDB) Set(key string, value interface{}) bool {
	return f.DB.Set(key, value)
}

// Remove deltes the key and all its values
func (f *FileDB) Remove(key string) bool {
	return f.DB.Remove(key)
}

// GetAll returns the inner instance of PickleDb
func (f *FileDB) GetAll() *pickledb.PickleDb {
	return f.DB
}

// Close shutdowns the connection to the file database
func (f *FileDB) Close() {
	return
}
