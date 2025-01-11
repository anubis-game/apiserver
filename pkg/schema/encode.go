package schema

func Encode(act Action, mes ...[]byte) []byte {
	var byt []byte
	{
		byt = []byte{byte(act)}
	}

	for _, x := range mes {
		byt = append(byt, x...)
	}

	return byt
}
