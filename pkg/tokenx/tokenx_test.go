package tokenx

import (
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
)

func Test_TokenX_Expiry(t *testing.T) {
	var err error
	var zer uuid.UUID

	var tkx *TokenX[string]
	{
		tkx = New[string]()
	}

	{
		tkx.tic = 1 * time.Millisecond
		tkx.ttl = 100 * time.Millisecond
	}

	{
		go tkx.Daemon()
	}

	{
		time.Sleep(100 * time.Millisecond)
	}

	var tk1 uuid.UUID
	var tk2 uuid.UUID
	var tk3 uuid.UUID
	{
		tk1 = uuid.MustParse("11111111-1111-1111-1111-111111111111")
		tk2 = uuid.MustParse("22222222-2222-2222-2222-222222222222")
		tk3 = uuid.MustParse("33333333-3333-3333-3333-333333333333")
	}

	{
		key, exi := tkx.Search(tk1)
		if exi {
			t.Fatalf("expected %#v got %#v", "no key", key)
		}
	}

	{
		key, exi := tkx.Search(tk2)
		if exi {
			t.Fatalf("expected %#v got %#v", "no key", key)
		}
	}

	{
		key, exi := tkx.Search(tk3)
		if exi {
			t.Fatalf("expected %#v got %#v", "no key", key)
		}
	}

	var ky1 string
	var ky2 string
	var ky3 string
	{
		ky1 = "ky1"
		ky2 = "ky2"
		ky3 = "ky3"
	}

	{
		tk1, err = tkx.Create(ky1)
		if err != nil {
			t.Fatal(err)
		}
		if tk1 == zer {
			t.Fatalf("expected %#v got %#v", "random UUID", tk1)
		}
	}

	{
		time.Sleep(50 * time.Millisecond)
	}

	{
		tk2, err = tkx.Create(ky2)
		if err != nil {
			t.Fatal(err)
		}
		if tk2 == zer {
			t.Fatalf("expected %#v got %#v", "random UUID", tk2)
		}
	}

	{
		tk3, err = tkx.Create(ky3)
		if err != nil {
			t.Fatal(err)
		}
		if tk3 == zer {
			t.Fatalf("expected %#v got %#v", "random UUID", tk3)
		}
	}

	{
		time.Sleep(10 * time.Millisecond)
	}

	tkx.mut.Lock()
	if len(tkx.exp) != 3 {
		t.Fatalf("expected %#v got %#v", 3, len(tkx.exp))
	}
	if len(tkx.lis) != 3 {
		t.Fatalf("expected %#v got %#v", 3, len(tkx.lis))
	}
	if len(tkx.rev) != 3 {
		t.Fatalf("expected %#v got %#v", 3, len(tkx.rev))
	}
	tkx.mut.Unlock()

	{
		key, exi := tkx.Search(tk1)
		if !exi {
			t.Fatalf("expected %#v got %#v", ky1, key)
		}
		if key != ky1 {
			t.Fatalf("expected %#v got %#v", ky1, key)
		}
	}

	{
		key, exi := tkx.Search(tk2)
		if !exi {
			t.Fatalf("expected %#v got %#v", ky2, key)
		}
		if key != ky2 {
			t.Fatalf("expected %#v got %#v", ky2, key)
		}
	}

	{
		key, exi := tkx.Search(tk3)
		if !exi {
			t.Fatalf("expected %#v got %#v", ky3, key)
		}
		if key != ky3 {
			t.Fatalf("expected %#v got %#v", ky3, key)
		}
	}

	{
		time.Sleep(50 * time.Millisecond)
	}

	tkx.mut.Lock()
	if len(tkx.exp) != 2 {
		t.Fatalf("expected %#v got %#v", 2, len(tkx.exp))
	}
	if len(tkx.lis) != 2 {
		t.Fatalf("expected %#v got %#v", 2, len(tkx.lis))
	}
	if len(tkx.rev) != 2 {
		t.Fatalf("expected %#v got %#v", 2, len(tkx.rev))
	}
	tkx.mut.Unlock()

	{
		key, exi := tkx.Search(tk1)
		if exi {
			t.Fatalf("expected %#v got %#v", "no key", key)
		}
	}

	{
		key, exi := tkx.Search(tk2)
		if !exi {
			t.Fatalf("expected %#v got %#v", ky2, key)
		}
		if key != ky2 {
			t.Fatalf("expected %#v got %#v", ky2, key)
		}
	}

	{
		key, exi := tkx.Search(tk3)
		if !exi {
			t.Fatalf("expected %#v got %#v", ky3, key)
		}
		if key != ky3 {
			t.Fatalf("expected %#v got %#v", ky3, key)
		}
	}

	{
		time.Sleep(50 * time.Millisecond)
	}

	tkx.mut.Lock()
	if len(tkx.exp) != 0 {
		t.Fatalf("expected %#v got %#v", 0, len(tkx.exp))
	}
	if len(tkx.lis) != 0 {
		t.Fatalf("expected %#v got %#v", 0, len(tkx.lis))
	}
	if len(tkx.rev) != 0 {
		t.Fatalf("expected %#v got %#v", 0, len(tkx.rev))
	}
	tkx.mut.Unlock()
}

func Test_TokenX_Random(t *testing.T) {
	var tkx *TokenX[string]
	{
		tkx = New[string]()
	}

	if len(tkx.exp) != 0 {
		t.Fatalf("expected %#v got %#v", 0, len(tkx.exp))
	}
	if len(tkx.lis) != 0 {
		t.Fatalf("expected %#v got %#v", 0, len(tkx.lis))
	}
	if len(tkx.rev) != 0 {
		t.Fatalf("expected %#v got %#v", 0, len(tkx.rev))
	}

	prv := []uuid.UUID{}
	see := map[uuid.UUID]int{}
	for range 1000 {
		tok, err := tkx.Create("ky1")
		if err != nil {
			t.Fatal(err)
		}

		// Check that we do not accumulate any resources when we overwrite existing
		// tokens.

		if len(tkx.exp) != 1 {
			t.Fatalf("expected %#v got %#v", 1, len(tkx.exp))
		}
		if len(tkx.lis) != 1 {
			t.Fatalf("expected %#v got %#v", 1, len(tkx.lis))
		}
		if len(tkx.rev) != 1 {
			t.Fatalf("expected %#v got %#v", 1, len(tkx.rev))
		}

		{
			see[tok]++
		}

		// No overwritten token is allowed to exist anymore. So search for all
		// previous tokens and verify that none of them refers back to any key.

		for _, p := range prv {
			key, exi := tkx.Search(p)
			if exi {
				t.Fatalf("expected %#v got %#v", "no key", key)
			}
		}

		{
			prv = append(prv, tok)
		}
	}

	for _, v := range see {
		if v != 1 {
			t.Fatalf("expected %#v got %#v", "unique key", v)
		}
	}
}

// ~400.00 ns/op, 1 allocs/op
func Benchmark_TokenX_Create(b *testing.B) {
	var tkx *TokenX[string]
	{
		tkx = New[string]()
	}

	for b.Loop() {
		_, err := tkx.Create("ky1")
		if err != nil {
			b.Fatal(err)
		}
	}
}

// ~0.23 ns/op
func Benchmark_TokenX_concurrency(b *testing.B) {
	n := 250_000
	k := "key"
	t := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	x := New[string]()

	{
		x.tic = 1 * time.Millisecond
		x.ttl = 100 * time.Millisecond
	}

	{
		go x.Daemon()
	}

	{
		time.Sleep(100 * time.Millisecond)
	}

	w := sync.WaitGroup{}
	c := make(chan struct{})

	go func() {
		<-c

		for range n {
			_, err := x.Create(k)
			if err != nil {
				panic(err)
			}
		}

		w.Done()
	}()

	go func() {
		<-c

		for range n {
			x.Search(t)
		}

		w.Done()
	}()

	w.Add(2)
	close(c)
	w.Wait()
}
