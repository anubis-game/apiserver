package schema

import "bytes"

func Encode(act Action, mes ...[]byte) []byte {
	return bytes.Join(append([][]byte{{byte(act)}}, mes...), Comma)
}
