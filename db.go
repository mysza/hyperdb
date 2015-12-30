package hyperdb

import (
	"errors"
	"github.com/boltdb/bolt"
	"time"
)

const defaultDataFileName = "hyper.db"

var (
	ErrCannotCloseNilDB     = errors.New("attempted to close nil database")
	ErrCannotCloseNilBoltDB = errors.New("attemted to close nil Bolt database")
)

// DB represents a collection of objects with multitude of attributes,
// that can be retrieved based all possible dimensions.
type DB struct {
	boltHandle *bolt.DB
}

// Open creates and opens a database at the given path.
// If the file doesn't exist, it is created automatically.
func Open(path string) (*DB, error) {
	if path == "" {
		path = defaultDataFileName
	}
	db, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	return &DB{boltHandle: db}, nil
}

// Close closes the database.
func (db *DB) Close() error {
	if db == nil {
		return ErrCannotCloseNilDB
	}
	if db.boltHandle == nil {
		return ErrCannotCloseNilBoltDB
	}
	return db.boltHandle.Close()
}

// Add adds an object to the given namespace - indexes all attributes and stores data
// accessible through GetData() method of the HyperObject.
// The function returns unique id of the object stored in hyperdb that can
// later be used to update or remove object in/from the namespace.
func (db *DB) Add(namespace string, obj *HyperObject) (string, error) {
	return "", nil
}

// Update updates the object with specified unique id in the given namespace.
func (db *DB) Update(namespace string, uid string, obj *HyperObject) error {
	return nil
}

// Remove removes passed object from the given namespace
// based on the unique id of the object.
func (db *DB) Remove(namespace string, uid string) ([]byte, error) {
	return nil, nil
}
