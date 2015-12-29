package core

// HyperObject is an interface expected to be implemented by the objects
// that are about to be stored in the hyperdb.
type HyperObject interface {
	// GetAttributes returns a slice of all the attributes of the object
	GetAttributes() []Attribute
	// GetData retuns a raw representation of the object.
	// The client of the hyperdb is responsible for objects serialization.
	GetData() []byte
}
