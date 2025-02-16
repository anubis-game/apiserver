package schema

func Encode(act Action, mes []byte) []byte {
	byt := make([]byte, 1+len(mes))

	{
		byt[0] = byte(act)
	}

	{
		copy(byt[1:], mes)
	}

	return byt
}
