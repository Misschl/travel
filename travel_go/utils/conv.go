package utils

func Uint8SliceString(s []uint8) string {
	result := []byte{}
	for _, i := range s {
		result = append(result, byte(i))
	}
	return string(result)
}
