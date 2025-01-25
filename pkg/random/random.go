package random

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Buf int
	Don <-chan struct{}
	Log logger.Interface
	Max byte
	Min byte
}

type Random struct {
	bck []byte
	don <-chan struct{}
	log logger.Interface
	max byte
	min byte
	que chan byte
	siz *big.Int
}

func New(c Config) *Random {
	if c.Buf == 0 {
		tracer.Panic(fmt.Errorf("%T.Buf must not be empty", c))
	}
	if c.Don == nil {
		tracer.Panic(fmt.Errorf("%T.Don must not be empty", c))
	}
	if c.Log == nil {
		tracer.Panic(fmt.Errorf("%T.Log must not be empty", c))
	}
	if c.Max == 0 {
		tracer.Panic(fmt.Errorf("%T.Max must not be empty", c))
	}

	// We need to cast min and max into integers before calculating the size,
	// because min=0 max=255 results in a size of 0 due to the MaxUint8 overflow
	// when doing +1 below. A size of 0 does then cause runtime panics for any
	// random distribution for the full uint8 spectrum. We have unit tests for
	// that case. See Test_Random_Random_uint8 and Test_Random_backup_uint8.
	var siz int
	{
		siz = int(c.Max) - int(c.Min) + 1
	}

	var bck []byte
	{
		bck = make([]byte, siz)
	}

	for i := range bck {
		bck[i] = c.Min + byte(i)
	}

	{
		musShf(bck)
	}

	return &Random{
		bck: bck,
		don: c.Don,
		log: c.Log,
		max: c.Max,
		min: c.Min,
		que: make(chan byte, c.Buf),
		siz: big.NewInt(int64(siz)),
	}
}

func (r *Random) Daemon() {
	for {
		select {
		case <-r.don:
			return
		case r.que <- r.random():
			// repeat
		}
	}
}

func (r *Random) Random() byte {
	return <-r.que
}

func (r *Random) random() byte {
	// Generate a cryptographically secure random number in the range [0, siz).
	b, err := rand.Int(rand.Reader, r.siz)
	if err != nil {
		r.log.Log(
			"level", "error",
			"message", err.Error(),
			"stack", tracer.Stack(err),
		)

		return r.backup()
	}

	// Add min to shift the range to [min, max].
	return byte(b.Int64()) + r.min
}

func (r *Random) backup() byte {
	bck := r.bck[0]
	r.bck = append(r.bck[1:], bck)
	return bck
}

func musShf(lis []byte) {
	for i := len(lis) - 1; i > 0; i-- {
		b, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			tracer.Panic(err)
		}

		j := int(b.Int64())

		{
			lis[i], lis[j] = lis[j], lis[i]
		}
	}
}
