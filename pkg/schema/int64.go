package schema

import (
	"strconv"

	"github.com/xh3b4sd/tracer"
)

func Int64(byt []byte) (int64, error) {
	num, err := strconv.ParseInt(string(byt), 10, 64)
	if err != nil {
		return 0, tracer.Mask(err)
	}

	return num, nil
}
