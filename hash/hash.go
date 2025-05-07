package hash

func rotateLeft(val uint64, shift uint64) uint64 {
	return (val << shift) | (val >> (64 - shift))
}

func xxhash64(input string) uint64 {
	var h64 uint64 = 0xAAAAAAAAAAAAAAAA
	lenInput := uint64(len(input))
	const limit = 32

	var i uint64
	for i = 0; i+limit <= lenInput; i += limit {
		block := input[i : i+limit]
		h64 += uint64(len(block))
		h64 = rotateLeft(h64, 27)
		h64 ^= mixBlock(block)
		h64 = h64*prime64v1 + prime64v2
	}

	remainder := input[i:]
	if len(remainder) > 0 {
		h64 ^= uint64(len(remainder))
		h64 = rotateLeft(h64, 23)
		h64 ^= mixBlock(remainder)
		h64 = h64 * prime64v3
	}

	h64 ^= h64 >> 33
	h64 *= prime64v2
	h64 ^= h64 >> 29
	h64 *= prime64v3
	h64 ^= h64 >> 32

	return h64
}

func mixBlock(block string) uint64 {
	var hash uint64 = 0
	for i := 0; i < len(block); i++ {
		hash ^= uint64(block[i])
		hash = rotateLeft(hash, 7)
		hash *= prime64v1
	}
	return hash
}

func GenerateHash(input string) string {
	hashVal := xxhash64(input)
	reduced := hashVal % mod
	hashString := base62Conversion(reduced)

	for len(hashString) < 10 {
		hashString = "0" + hashString
	}

	return hashString
}
