package core

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

func (AttributeBase ab) GetType() string {
	return ab.Type
}

func (AttributeBase ab) GetName() string {
	return ab.Name
}

func (AttributeBase ab) GetValue() []byte {
	return ab.Value
}
