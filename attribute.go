package hyperdb

// Attribute is a set property of an object that needs to be indexed in the db
// to be easily findable.
type Attribute interface {
	GetType() string
	GetName() string
	GetValue() []byte
}

// AttributeBase is a naive hyperdb implementation of the Attribute interface.
type AttributeBase struct {
	Type  string
	Name  string
	Value []byte
}

// GetType returns the type of the AttributeBase object.
func (ab AttributeBase) GetType() string {
	return ab.Type
}

// GetName returns the name of the AttributeBase object.
func (ab AttributeBase) GetName() string {
	return ab.Name
}

// GetValue returns the value of the AttributeBase object.
func (ab AttributeBase) GetValue() []byte {
	return ab.Value
}
