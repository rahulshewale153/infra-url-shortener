package utils

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateEncodeBase62(length int) string {
	result := []byte{}
	for length > 0 {
		result = append([]byte{alphabet[length%62]}, result...)
		length /= 62
	}
	return string(result)
}
