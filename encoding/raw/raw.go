package raw

type DataFormat struct{}

func (DataFormat) Encode(input interface{}) ([]byte, error)    { return input.([]byte), nil }
func (DataFormat) Decode(data []byte, value interface{}) error { return nil }
func (DataFormat) String() string                              { return "raw" }
