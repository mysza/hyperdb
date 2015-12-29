package hyperdb

// DB represents a collection of objects with multitude of attributes,
// that can be retrieved based all possible dimensions.
type DB struct {
}

// Open creates and opens a database at the given path.
// If the file doesn't exist, it is created automatically.
func Open(path string) (*DB, error) {
	return nil, nil
}

// Add adds an object to the given namespace - indexes all attributes and stores data
// accessible through GetData() method of the HyperObject.
// The function returns unique id of the object stored in hyperdb that can
// later be used to update or remove object in/from the namespace.
func (db *DB) Add(namespace string, obj *HyperObject) (uint64, error) {
	return nil
}

// Update updates the object with specified unique id in the given namespace.
func (db *DB) Update(namespace string, uid uint64, obj *HyperObject) error {
	return nil
}

// Remove removes passed object from the given namespace
// based on the unique id of the object.
func (db *DB) Remove(namespace string, uid uint64) ([]byte, error) {
	return nil, nil
}
