package utils

func Uint8SliceString(s []uint8) string {
	var result []byte
	for _, i := range s {
		result = append(result, i)
	}
	return string(result)
}
