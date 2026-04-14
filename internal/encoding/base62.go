package encoding

var base62_str string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

func EncodeBase62(num int64) string {
	var result []byte
	for num > 0 {
		result = append(result, base62_str[num % 62])
		num = num / 62 // reduce the num
	}

	return string(result)
}