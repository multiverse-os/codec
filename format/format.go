package format

import (
	bson "./bson"
	cbor "./cbor"
	gob "./gob"
	json "./json"
	raw "./raw"
)

type Type int

const (
	Raw Type = iota
	JSON
	BSON
	CBOR
	GOB
)

type Formater interface {
	Encode(input interface{}) ([]byte, error)
	Decode(data []byte, output interface{}) error
}

func Load(dataType Type) Formater {
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
