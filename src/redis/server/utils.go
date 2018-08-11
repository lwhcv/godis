package server

import (
	"fmt"
	"strconv"
)

func MustString(any interface{}) string {
	if any == nil {
		return ""
	}
	switch val := any.(type) {
	case int, uint, int64, uint64, uint32, int32, uint8, int8, int16, uint16:
		return fmt.Sprintf("%d", val)
	case string:
		return val
	case []byte:
		return string(val)
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)

	default:
		return fmt.Sprintf("%v", val)
	}

	return ""
}

func InArray(s string, arr []string) bool {
	for _, v := range arr {
		if s == v {
			return true
		}
	}
	return false
}
