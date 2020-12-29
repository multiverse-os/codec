package format

type DataType int

const (
	Raw DataType = iota
	JSON
	BSON
	CBOR
	GOB
)

type Formater interface {
	Encode(input interface{}) ([]byte, error)
	Decode(data []byte, output interface{}) error
}

func Load(dataType DataType) Formater {
	switch dataType {
	case JSON:
		return jsonFormat{}
	case BSON:
		return bsonFormat{}
	case CBOR:
		return cborFormat{}
	case GOB:
		return gobFormat{}
	default: // Raw
		return rawFormat{}
	}

}
