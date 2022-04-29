package hash

import (
	"errors"
	"strconv"
)

var (
	ErrUnSupportedType = errors.New("unsupported type encoding, only number can do encode this hash")
	ErrHash            = errors.New("invalid hash")
	ErrData            = errors.New("hash data is nil")
)

func (h *hash) Encode(v interface{}) (string, error) {
	if h == nil {
		return "", ErrData
	}
	
	var val []int
	
	// check type of parameters
	switch v.(type) {
	case int32:
		val = append(val, int(v.(int32)))
	case int64:
		val = append(val, int(v.(int64)))
	case int:
		val = append(val, v.(int))
	case string:
		s, err := strconv.Atoi(v.(string))
		if err != nil {
			return "", err
		}
		
		val = append(val, s)
	default:
		// if the type is map, struct, float e.t.c will return an error unsupported type
		return "", ErrUnSupportedType
	}
	
	// encode value ids
	en, err := h.h.Encode(val)
	if err != nil {
		return "", err
	}
	
	return en, nil
}
