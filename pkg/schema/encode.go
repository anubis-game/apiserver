package schema

func Encode(act byte, mes []byte) []byte {
	byt := make([]byte, 1+len(mes))

	{
		byt[0] = act
	}

	{
		copy(byt[1:], mes)
	}

	return byt
}
