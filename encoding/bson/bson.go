package bson

import (
	bson "go.mongodb.org/mongo-driver/bson"
)

type DataFormat struct{}

func (DataFormat) Encode(value interface{}) ([]byte, error) {
	return bson.Marshal(value)
}

func (DataFormat) Decode(data []byte, value interface{}) error {
	return bson.Unmarshal(data, value)
}

func (DataFormat) String() string { return "bson" }
