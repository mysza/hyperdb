package hyperdb

import (
	"errors"
	"time"

	"github.com/boltdb/bolt"
)

// DefaultDataFileName is the default file name of the database file.
const DefaultDataFileName = "hyper.db"

var (
	// ErrCannotCloseNilDB is an error returned when attempted to close nil database.
	ErrCannotCloseNilDB = errors.New("attempted to close nil database")
	// ErrCannotCloseNilBoltDB is an error returned when attempted to close nil Bolt database.
	ErrCannotCloseNilBoltDB = errors.New("attempted to close nil Bolt database")
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
		path = DefaultDataFileName
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

// Get looks up the ids of the objects that fit the search criteria and returns them to the client.
func (db *DB) Get(searchCriteria string) ([]string, error) {
	return nil, nil
}

// Put adds an object to the database if it does not exist, or updates existing entries based on the new
// object definition.
// The function returns unique id of the object stored in hyperdb that can
// later be used to update or remove object in/from the namespace.
func (db *DB) Put(obj HyperObject) (string, error) {
	objID := obj.GetID()
	isNew := false
	if len(objID) == 0 {
		// new object, add to objects list
		objID = db.assureID(obj)
		isNew = true
	}
	// update the indexes
	db.updateIndexes(objID, obj.GetAttributes(), isNew)
	return objID, nil
}

// Remove removes passed object from the database
// based on the unique id of the object.
func (db *DB) Remove(uid string) error {
	return nil
}

// assureID assures that given objects gets a unique ID in the database.
func (db *DB) assureID(obj HyperObject) string {
	return ""
}

// updateIndexes updates indexes assigned to an object with ID = uid.
// If database has indexes on attributes not provided in the parameters, those indexes will be removed.
func (db *DB) updateIndexes(uid string, attributes []Attribute, skipExistingCheck bool) error {
	return nil
}
