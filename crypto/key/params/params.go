package params

import (
	xxhash "github.com/cespare/xxhash"
)

type Params map[uint64]Value

// PARAMETERS //////////////////////////////////////////////////////////////////
type dataType int

const (
	strType dataType = iota
	intType
	floatType
	boolType
	byteSliceType
)

type dataValue struct {
	stringValue    string
	integerValue   int
	floatValue     float64
	booleanValue   bool
	byteSliceValue []byte
}

type Value struct {
	dataType
	dataValue
}

////////////////////////////////////////////////////////////////////////////////
func hashKey(key string) uint64 {
	return xxhash.Sum64String(key)
}

// SET Parameter ///////////////////////////////////////////////////////////////
func (self Params) AddString(key string, value string) Params {
	self[hashKey(key)] = String(value)
	return self
}

func (self Params) AddInteger(key string, value int) Params {
	self[hashKey(key)] = Integer(value)
	return self
}

func (self Params) AddFloat(key string, value float64) Params {
	self[hashKey(key)] = Float(value)
	return self
}

func (self Params) AddBoolean(key string, value bool) Params {
	self[hashKey(key)] = Boolean(value)
	return self
}

func (self Params) AddBytes(key string, value []byte) Params {
	self[hashKey(key)] = Bytes(value)
	return self
}

// GET Parameter ///////////////////////////////////////////////////////////////
func (self Params) String(key string) string { return self[hashKey(key)].stringValue }
func (self Params) Integer(key string) int   { return self[hashKey(key)].integerValue }
func (self Params) Float(key string) float64 { return self[hashKey(key)].floatValue }
func (self Params) Boolean(key string) bool  { return self[hashKey(key)].booleanValue }
func (self Params) Bytes(key string) []byte  { return self[hashKey(key)].byteSliceValue }

////////////////////////////////////////////////////////////////////////////////
func String(value string) Value {
	return Value{
		dataType: strType,
		dataValue: dataValue{
			stringValue: value,
		},
	}
}

func Integer(value int) Value {
	return Value{
		dataType: intType,
		dataValue: dataValue{
			integerValue: value,
		},
	}
}

func Float(value float64) Value {
	return Value{
		dataType: floatType,
		dataValue: dataValue{
			floatValue: value,
		},
	}
}

func Boolean(value bool) Value {
	return Value{
		dataType: boolType,
		dataValue: dataValue{
			booleanValue: value,
		},
	}
}

func Bytes(value []byte) Value {
	return Value{
		dataType: byteSliceType,
		dataValue: dataValue{
			byteSliceValue: value,
		},
	}
}
