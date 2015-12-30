package hyperdb

// HyperObject is an interface expected to be implemented by the objects
// that are about to be stored in the hyperdb.
type HyperObject interface {
	// GetId returns a unique ID of the object the hyperdb client wants to
	// use as a reference for the object. If 0 is returned, hyperdb assigns its
	// own ID.
	GetId() string
	// GetAttributes returns a slice of all the attributes of the object.
	GetAttributes() []Attribute
	// GetData retuns a raw representation of the object.
	// The client of the hyperdb is responsible for objects serialization.
	GetData() []byte
}