package randHash

// Tong Hash Function 64
func TongHash64(str string) uint64 {
	list := []byte(str)
	var seed uint64 = 131 // 31 131 1313 13131 131313 etc..
	var hash uint64 = 0
	for i := 0; i < len(list); i++ {
		hash = hash * seed + uint64(list[i])
	}
	return hash & 0x7FFFFFFFFFFFFFFF
}

