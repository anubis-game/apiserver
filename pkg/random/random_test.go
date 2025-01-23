package random

import (
	"fmt"
	"math"
	"testing"

	"github.com/anubis-game/apiserver/pkg/matrix"
	"github.com/xh3b4sd/logger"
)

func Test_Random_Random_coordinates(t *testing.T) {
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

func Test_Random_Random_uint8(t *testing.T) {
	var ran *Random
	{
		ran = New(Config{
			Buf: 1000,
			Don: make(<-chan struct{}),
			Log: logger.Fake(),
			Max: math.MaxUint8,
		})
	}

	{
		go ran.Daemon()
	}

	dic := map[byte]int{}

	for i := 0; i < 10000; i++ {
		dic[ran.Random()]++
	}

	if len(dic) != 256 {
		t.Fatal("expected", 256, "got", len(dic))
	}

	for k, v := range dic {
		if v < 2 {
			t.Fatalf("expected %d to appear at least twice, got %d", k, v)
		}
	}
}

func Test_Random_backup_coordinates(t *testing.T) {
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

func Test_Random_backup_uint8(t *testing.T) {
	var ran *Random
	{
		ran = New(Config{
			Buf: 1000,
			Don: make(<-chan struct{}),
			Log: logger.Fake(),
			Max: math.MaxUint8,
		})
	}

	{
		go ran.Daemon()
	}

	dic := map[byte]int{}

	for i := 0; i < 1000; i++ {
		dic[ran.backup()]++
	}

	if len(dic) != 256 {
		t.Fatal("expected", 256, "got", len(dic))
	}

	for k, v := range dic {
		if v < 2 {
			t.Fatalf("expected %d to appear at least twice, got %d", k, v)
		}
	}
}
