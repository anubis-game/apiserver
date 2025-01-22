package random

import (
	"fmt"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/xh3b4sd/logger"
)

func Test_Random_Random(t *testing.T) {
	var ran *Random
	{
		ran = New(Config{
			Buf: 1000,
			Don: make(<-chan struct{}),
			Log: logger.Fake(),
			Max: matrix.Max,
			Min: matrix.Min,
		})
	}

	{
		go ran.Daemon()
	}

	dic := map[byte]int{}

	for i := 0; i < 1000; i++ {
		var b byte
		{
			b = ran.Random()
		}

		if b < matrix.Min {
			t.Fatal("expected", fmt.Sprintf(">= %d", matrix.Min), "got", int(b))
		}

		if b > matrix.Max {
			t.Fatal("expected", fmt.Sprintf("<= %d", matrix.Max), "got", int(b))
		}

		{
			dic[b]++
		}
	}

	if len(dic) != int(matrix.Siz) {
		t.Fatal("expected", matrix.Siz, "got", len(dic))
	}

	for k, v := range dic {
		if v < 2 {
			t.Fatalf("expected %d to appear at least twice, got %d", k, v)
		}
	}
}

func Test_Random_backup(t *testing.T) {
	var ran *Random
	{
		ran = New(Config{
			Buf: 1000,
			Don: make(<-chan struct{}),
			Log: logger.Fake(),
			Max: matrix.Max,
			Min: matrix.Min,
		})
	}

	{
		go ran.Daemon()
	}

	dic := map[byte]int{}

	for i := 0; i < 1000; i++ {
		var b byte
		{
			b = ran.backup()
		}

		if b < matrix.Min {
			t.Fatal("expected", fmt.Sprintf(">= %d", matrix.Min), "got", int(b))
		}

		if b > matrix.Max {
			t.Fatal("expected", fmt.Sprintf("<= %d", matrix.Max), "got", int(b))
		}

		{
			dic[b]++
		}
	}

	if len(dic) != int(matrix.Siz) {
		t.Fatal("expected", matrix.Siz, "got", len(dic))
	}

	for k, v := range dic {
		if v < 2 {
			t.Fatalf("expected %d to appear at least twice, got %d", k, v)
		}
	}
}
