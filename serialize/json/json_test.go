package json

import (
	"reflect"
	"testing"
)

type yServer struct {
	X int    `json:"x"`
	Y string `json:"y"`
}

func TestSerializer_Marshal(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		s       *Serializer
		args    args
		want    []byte
		wantErr bool
	}{
		{"Ok", NewSerializer(), args{yServer{X: 20210416, Y: "YONG"}}, ([]byte)(`{"x":20210416,"y":"YONG"}`), false},
	}
	for _, tt := range tests {
		s := &Serializer{}
		got, err := s.Marshal(tt.args.v)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Serializer.Marshal() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Serializer.Marshal() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestNewSerializer(t *testing.T) {
	tests := []struct {
		name string
		want *Serializer
	}{
		{"ok", NewSerializer()},
	}
	for _, tt := range tests {
		if got := NewSerializer(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. NewSerializer() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestSerializer_Unmarshal(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		s       *Serializer
		args    args
		wantErr bool
	}{
		{"OK", NewSerializer(), args{data: ([]byte)(`{"x":20210416,"y":"YONG"}`), v: &yServer{}}, false},
	}
	for _, tt := range tests {
		s := &Serializer{}
		if err := s.Unmarshal(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
			t.Errorf("%q. Serializer.Unmarshal() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestSerializer_GetName(t *testing.T) {
	tests := []struct {
		name string
		s    *Serializer
		want string
	}{
		{"Ok", NewSerializer(), "json"},
	}
	for _, tt := range tests {
		s := &Serializer{}
		if got := s.GetName(); got != tt.want {
			t.Errorf("%q. Serializer.GetName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
