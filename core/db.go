package core

// DB represents a collection of objects with multitude of attributes,
// that can be retrieved based all possible dimensions.
type DB struct {
}

// Open creates and opens a database at the given path.
// If the file doesn't exist, it is created automatically.
func Open(path string) (*DB, error) {
	return nil, nil
}
