package bson

type dataFormat struct{}

func (dataFormat) Encode(obj interface{}) ([]byte, error) {
	//return bson.Marshal(obj)
	return []byte{}, nil
}

func (dataFormat) Decode(data []byte, value interface{}) error {
	//return bson.Unmarshal(data, &value)
	return nil
}

func (dataFormat) String() string { return "bson" }
