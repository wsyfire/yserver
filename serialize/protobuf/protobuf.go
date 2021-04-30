package protobuf

import (
	"errors"
	"github.com/gogo/protobuf/proto"
)

var (
	ErrWrongType = errors.New("protobuf: convert on wrong type")
)

type Serializer struct {
}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) Marshal(v interface{}) ([]byte, error) {
	pb, ok := v.(proto.Message)
	if !ok {
		return nil, ErrWrongType
	}

	return proto.Marshal(pb)
}

func (s *Serializer) Unmarshal(data []byte, v interface{}) error {
	pb, ok := v.(proto.Message)
	if !ok {
		return ErrWrongType
	}

	return proto.Unmarshal(data, pb)
}

func (s *Serializer) GetName() string {
	return "protobuf"
}
