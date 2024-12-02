package stream

import (
	"fmt"
	"sync"

	"github.com/coder/websocket"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	Log logger.Interface
}

type Stream struct {
	cli map[string]Client
	log logger.Interface
	mut sync.RWMutex
}

func New(c Config) *Stream {
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}

	return &Stream{
		cli: map[string]Client{},
		log: c.Log,
		mut: sync.RWMutex{},
	}
}

func (s *Stream) Add(add string, cli Client) {
	s.mut.Lock()
	old, exi := s.cli[add]
	s.cli[add] = cli
	s.mut.Unlock()

	if exi {
		old.Close(true)
	}
}

func (s *Stream) Wri(add string, typ websocket.MessageType, byt []byte) {
	s.mut.RLock()
	for _, x := range s.cli {
		x.Write(typ, byt)
	}
	s.mut.RUnlock()
}

func (s *Stream) Rem(add string) {
	s.mut.Lock()
	old, exi := s.cli[add]
	delete(s.cli, add)
	s.mut.Unlock()

	if exi {
		old.Close(false)
	}
}
