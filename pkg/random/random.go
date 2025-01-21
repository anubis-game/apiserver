package random

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
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
	var bck []byte
	{
		bck = make([]byte, c.Max-c.Min+1)
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
		que: make(chan byte, 500),
		siz: big.NewInt(int64(c.Max - c.Min + 1)),
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
			context.Background(),
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
