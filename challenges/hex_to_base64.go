package challenges

func HexToBase64(s string) string {
	decodedS := DecodeHexString(s)
	b := []byte(padLength(string(decodedS)))
	return string(encodeBase64(b))
}

func DecodeHexString(s string) []byte {
	res := []byte(s)
	i, j := 0, 1
	// Move in two byte increments getting the decimal corresponding
	// to each hex character. Update the result with the character of the
	// decimal. Return only the portion of the byte slice containing decoded
	// characters.
	for ; j < len(s); j += 2 {
		first := fromHexChar(s[j-1])
		secnd := fromHexChar(s[j])
		res[i] = (first * 16) + secnd
		i++
	}
	return res[:i]
}

func fromHexChar(ch byte) byte {
	switch {
	case '0' <= ch && ch <= '9':
		return ch - '0'
	case 'a' <= ch && ch <= 'f':
		return ch - 'a' + 10
	case 'A' <= ch && ch <= 'F':
		return ch - 'A' + 10
	}
	panic("invalid hex char")
}

func padLength(s string) string {
	for len(s)%3 != 0 {
		s += "="
	}
	return s
}

func encodeBase64(b []byte) []byte {
	const encode = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	buf := make([]byte, (len(b)+2)/3*4)

	j, i := 0, 0
	n := (len(b) / 3) * 3
	for i < n {
		// Calculate the binary value of three bytes.
		val := uint(b[i+0])<<16 + uint(b[i+1])<<8 + uint(b[i+2])

		// Convert the 24 bits into four 6-bit values.
		buf[j+0] = encode[val>>18&0x3F]
		buf[j+1] = encode[val>>12&0x3F]
		buf[j+2] = encode[val>>6&0x3F]
		buf[j+3] = encode[val&0x3F]

		i += 3
		j += 4
	}

	// If we've reached the end of the byte slice, return the buffer.
	remain := len(b) - i
	if remain == 0 {
		return buf
	}
	// Otherwise add the remaining first (and maybe second) remaining values.
	val := uint(b[i+0]) << 16
	if remain == 2 {
		val += uint(buf[j+1]) << 8
	}

	// And convert them into two (and maybe three) 6-bit values.
	buf[j+0] = encode[val>>18&0x3F]
	buf[j+1] = encode[val>>12&0x3F]
	if remain == 2 {
		buf[j+2] = encode[val>>6&0x3F]
	}

	return buf
}
