// Package pickledb is an simple key-value db.
package pickledb

import (
	"fmt"
	"path/filepath"
)

// PickleDb represent a database object with the file path and the save option.
type PickleDb struct {
	Location string
	Option   string
	db       map[string]interface{}
}

// New return a empty PickleDb object. location is the path to the json file.
func New(location, option string) *PickleDb {
	return &PickleDb{
		Location: location,
		Option:   option,
		db:       map[string]interface{}{},
	}
}

// load data to PickleDb object from json file
func (p *PickleDb) load() error {
	m, err := ReadFromJSONFile(p.Location)
	if err != nil {
		return err
	}
	p.db = m.(map[string]interface{})
	return nil
}

// save data to json file from an PickleDb object
func (p *PickleDb) dump() error {
	return WriteToJSONFile(p.db, p.Location)
}

// Load read the data from json file, return true if read the data correctly,
// else return false.
func (p *PickleDb) Load() bool {
	exist := IsFileExisted(p.Location)
	if exist {
		err := p.load()
		if err != nil {
			return false
		}
	}
	return true
}

// Dump save the data to json file, return true if save the data correctly,
// else return false.
func (p *PickleDb) Dump() bool {
	exist := IsFileExisted(p.Location)
	if !exist {
		dir, _ := filepath.Split(p.Location)
		err := MakeDir(dir)
		if err != nil {
			return false
		}
	}

	err := p.dump()
	if err != nil {
		return false
	}
	return true
}

// Set set the value(string,int,whatever) of a key
func (p *PickleDb) Set(key string, value interface{}) bool {
	p.db[key] = value
	p.dump()
	return true
}

// Get return the value of a key
func (p *PickleDb) Get(key string) interface{} {
	return p.db[key]
}

// GetAll return a list of all keys in db
func (p *PickleDb) GetAll() []string {
	keys := []string{}
	for key := range p.db {
		keys = append(keys, key)
	}
	return keys
}

// Remove remove a key
func (p *PickleDb) Remove(key string) bool {
	delete(p.db, key)
	p.dump()
	return true
}

// Append add more to a key's value
func (p *PickleDb) Append(key string, more interface{}) bool {
	tmp := (p.db[key]).(string)
	value := fmt.Sprintf("%s%s", tmp, more.(string))

	p.db[key] = value
	p.dump()

	return true
}

// Destroy delete everything from the database
func (p *PickleDb) Destroy() bool {
	p.db = nil
	p.dump()

	return true
}

// Pair represents a tuple
type Pair struct {
	Key   interface{}
	Value interface{}
}

// ListCreate create a list
func (p *PickleDb) ListCreate(name string) bool {
	p.db[name] = []interface{}{}
	p.dump()
	return true
}

// ListAdd add a value to a list
func (p *PickleDb) ListAdd(name string, value interface{}) bool {
	l := (p.db[name]).([]interface{})
	l = append(l, value)
	p.db[name] = l
	p.dump()
	return true
}

// ListExtend extend a list with a sequence
func (p *PickleDb) ListExtend(name string, seq []interface{}) bool {
	l := (p.db[name]).([]interface{})
	l = append(l, seq)
	p.db[name] = l
	p.dump()
	return true
}

// ListGetAll return all values in a list
func (p *PickleDb) ListGetAll(name string) []interface{} {
	return (p.db[name]).([]interface{})
}

// ListGet return one value in a list by given postion
func (p *PickleDb) ListGet(name string, pos int) interface{} {
	return (p.db[name]).([]interface{})[pos]
}

// ListDel delete a list and all of its values
func (p *PickleDb) ListDel(name string) int {
	number := p.ListLen(name)
	p.Remove(name)
	return number
}

// ListPop remove one value in a list
func (p *PickleDb) ListPop(name string, pos int) interface{} {
	value := p.ListGet(name, pos)

	l := p.ListGetAll(name)
	l = append(l[:pos], l[pos+1:]...) // slice delete
	p.db[name] = l

	return value
}

// ListLen returns the length of the list
func (p *PickleDb) ListLen(name string) int {
	return len(p.ListGetAll(name))
}

// ListAppend add more to a val in a list
func (p *PickleDb) ListAppend(name string, pos int, more interface{}) bool {
	tmp := p.ListGet(name, pos)
	value := fmt.Sprintf("%s%s", tmp.(string), more.(string))

	(p.db[name]).([]interface{})[pos] = value
	p.dump()
	return true
}

// DictCreate create a dictionary
func (p *PickleDb) DictCreate(name string) bool {
	p.db[name] = map[interface{}]interface{}{}
	p.dump()
	return true
}

// DictAdd add a key-value pair to a dict, "pair" is a tuple.
func (p *PickleDb) DictAdd(name string, pair Pair) bool {
	dict := (p.db[name]).(map[interface{}]interface{})
	dict[pair.Key] = pair.Value
	p.db[name] = dict

	p.dump()

	return true
}

// DictGet return the value for a key in a dictionary
func (p *PickleDb) DictGet(name string, key interface{}) interface{} {
	return (p.db[name]).(map[interface{}]interface{})[key]
}

// DictGetAll return all key-value pairs from a dict
func (p *PickleDb) DictGetAll(name string) map[interface{}]interface{} {
	return (p.db[name]).(map[interface{}]interface{})
}

// DictRemove remove a dict and all of its pairs
func (p *PickleDb) DictRemove(name string) bool {
	p.Remove(name)
	p.dump()

	return true
}

// DictPop remove one key-value pair in a dict
func (p *PickleDb) DictPop(name string, key interface{}) interface{} {
	m := (p.db[name]).(map[interface{}]interface{})
	value := m[key]
	delete(m, key)
	p.db[name] = m
	p.dump()

	return value
}

// DictKeys return all the keys for a dictionary
func (p *PickleDb) DictKeys(name string) []interface{} {
	m := (p.db[name]).(map[interface{}]interface{})
	var keys []interface{}
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// DictValues return all the values for a dictionary
func (p *PickleDb) DictValues(name string) []interface{} {
	m := (p.db[name]).(map[interface{}]interface{})
	var values []interface{}
	for _, v := range m {
		values = append(values, v)
	}

	return values
}

// DictExists determine if a key exists or not
func (p *PickleDb) DictExists(name string, key interface{}) bool {
	_, exist := (p.db[name]).(map[interface{}]interface{})[key]
	return exist
}
