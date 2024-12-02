package decode

import "strconv"

func Int64(byt []byte) (int64, error) {
	return strconv.ParseInt(string(byt), 10, 64)
}
