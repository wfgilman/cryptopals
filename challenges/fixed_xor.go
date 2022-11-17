package challenges

func HexStringFixedXOR(a string, b string) string {
	da := DecodeHexString(a)
	db := DecodeHexString(b)
	return EncodeHexToString(FixedXOR(da, db))
}

func FixedXOR(a []byte, b []byte) []byte {
	if len(a) != len(b) {
		panic("byte slices are not equal lengths")
	}

	buf := make([]byte, len(a))

	for i := 0; i < len(a); i++ {
		buf[i] = a[i] ^ b[i]
	}

	return buf
}
