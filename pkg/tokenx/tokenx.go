package tokenx

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

// TokenX generates numerical IDs in constant time within a given capacity. Note
// that we use a normal sync.Mutex for synchronization instead of a
// sync.RWMutex, because the synchronized code is many times faster than the
// additional overhead incurred by sync.RWMutex.
type TokenX[K comparable] struct {
	exp map[K]int64
	lis map[K]uuid.UUID
	mut sync.Mutex
	rev map[uuid.UUID]K
	tic time.Duration
	ttl time.Duration
}

func New[K comparable]() *TokenX[K] {
	return &TokenX[K]{
		exp: map[K]int64{},
		lis: map[K]uuid.UUID{},
		rev: map[uuid.UUID]K{},
		tic: time.Minute,
		ttl: time.Hour,
	}
}

func (t *TokenX[K]) Create(key K) (uuid.UUID, error) {
	// Try to create a new UUID and verify that it is in fact unique to our
	// internal state. If we cannot generate a valid UUID, then we return an
	// error.

	t.mut.Lock()

	tok, err := t.random()
	if err != nil {
		return uuid.UUID{}, tracer.Mask(err)
	}

	t.delete(key)
	t.create(key, tok)

	t.mut.Unlock()

	return tok, nil
}

func (t *TokenX[K]) Daemon() {
	for x := range time.Tick(t.tic) {
		var u int64
		{
			u = x.UnixNano()
		}

		t.mut.Lock()

		for k, v := range t.exp {
			if v < u {
				t.delete(k)
			}
		}

		t.mut.Unlock()
	}
}

func (t *TokenX[K]) Search(tok uuid.UUID) (K, bool) {
	t.mut.Lock()
	key, exi := t.rev[tok]
	t.mut.Unlock()
	return key, exi
}

func (t *TokenX[K]) create(key K, tok uuid.UUID) {
	t.exp[key] = time.Now().Add(t.ttl).UnixNano()
	t.lis[key] = tok
	t.rev[tok] = key
}

func (t *TokenX[K]) delete(key K) {
	old, exi := t.lis[key]
	if exi {
		delete(t.exp, key)
		delete(t.lis, key)
		delete(t.rev, old)
	}
}

func (t *TokenX[K]) random() (uuid.UUID, error) {
	for range 3 {
		tok, err := uuid.NewRandom()
		if err != nil {
			return uuid.UUID{}, tracer.Mask(err)
		}

		_, exi := t.rev[tok]
		if exi {
			continue
		}

		return tok, nil
	}

	return uuid.UUID{}, tracer.Mask(tokenAlreadyExistsError)
}
