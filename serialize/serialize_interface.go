package serialize

type Marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

type Unmarshaler interface {
	Unmarshal(data []byte, v interface{}) error
}

type Serializer interface {
	Marshaler
	Unmarshaler
	GetName() string
}
