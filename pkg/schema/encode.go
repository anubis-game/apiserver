package schema

func Encode(act Action, mes ...[]byte) []byte {
	off, siz := 1, 1

	for _, x := range mes {
		siz += len(x)
	}

	byt := make([]byte, siz)
	byt[0] = byte(act)

	for _, x := range mes {
		off += copy(byt[off:off+len(x)], x)
	}

	return byt
}
