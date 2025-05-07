package hash

func base62Conversion(num uint64) string {
	const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if num == 0 {
		return "0"
	}
	encoded := ""
	base := uint64(62)

	for num > 0 {
		remainder := num % base
		encoded = string(base62Chars[remainder]) + encoded
		num /= base
	}
	return encoded
}
