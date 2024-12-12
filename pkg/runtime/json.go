package runtime

import (
	"encoding/json"

	"github.com/xh3b4sd/tracer"
)

func Json() []byte {
	return byt
}

func marshal() {
	var err error

	{
		byt, err = json.MarshalIndent(dic, "", "  ")
		if err != nil {
			tracer.Panic(tracer.Mask(err))
		}
	}
}
