package random

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Buf is the amount of random integers prepared in advance.
	Buf int
	// Don is the global done channel.
	Don <-chan struct{}
	// Log is used to print runtime information.
	Log logger.Interface
	// Max is the inclusive maximum integer randomly generated.
	Max int
	// Min is the inclusive minimum integer randomly generated.
	Min int
}

type Random struct {
	bck []int
	don <-chan struct{}
	log logger.Interface
	max int
	min int
	que chan int
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

	var siz int
	{
		siz = c.Max - c.Min + 1
	}

	var bck []int
	{
		bck = make([]int, siz)
	}

	for i := range bck {
		bck[i] = c.Min + i
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
		que: make(chan int, c.Buf),
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

func (r *Random) Random() int {
	return <-r.que
}

func (r *Random) random() int {
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
	return int(b.Int64()) + r.min
}

func (r *Random) backup() int {
	bck := r.bck[0]
	r.bck = append(r.bck[1:], bck)
	return bck
}

func musShf(lis []int) {
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
