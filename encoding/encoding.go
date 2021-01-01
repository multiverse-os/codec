package encoding

import (
	bson "github.com/multiverse-os/levelup/data/codec/encoding/bson"
	cbor "github.com/multiverse-os/levelup/data/codec/encoding/cbor"
	gob "github.com/multiverse-os/levelup/data/codec/encoding/gob"
	json "github.com/multiverse-os/levelup/data/codec/encoding/json"
	raw "github.com/multiverse-os/levelup/data/codec/encoding/raw"
)

type Type int

const (
	Raw Type = iota
	JSON
	BSON
	CBOR
	GOB
)

type Codec interface {
	Encode(input interface{}) ([]byte, error)
	Decode(data []byte, output interface{}) error
	String() string
}

func Format(dataType Type) Codec {
	switch dataType {
	case JSON:
		return json.DataFormat{}
	case BSON:
		return bson.DataFormat{}
	case CBOR:
		return cbor.DataFormat{}
	case GOB:
		return gob.DataFormat{}
	default: // Raw
		return raw.DataFormat{}
	}
}
