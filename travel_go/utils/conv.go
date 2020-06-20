package utils

import (
	"strconv"
)

func Uint8SliceToInt(s []uint8) (int, error) {
	result := []byte{}
	for _, i := range s {
		result = append(result, byte(i))
	}
	return strconv.Atoi(string(result))
}
