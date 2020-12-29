package bson

type DataFormat struct{}

func (DataFormat) Encode(obj interface{}) ([]byte, error) {
	//return bson.Marshal(obj)
	return []byte{}, nil
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	//return bson.Unmarshal(data, &value)
	return nil
}

func (DataFormat) String() string { return "bson" }
