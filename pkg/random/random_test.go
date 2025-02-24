package random

import (
	"fmt"
	"math"
	"testing"

	"github.com/xh3b4sd/logger"
)

func Test_Random_Random_coordinates(t *testing.T) {
	var ran *Random
	{
		ran = New(Config{
			Buf: 1000,
			Don: make(<-chan struct{}),
			Log: logger.Fake(),
			Max: 1100,
			Min: 1000,
		})
	}

	{
		go ran.Daemon()
	}

	dic := map[int]int{}

	for range 10000 {
		var b int
		{
			b = ran.Random()
		}

		if b < 1000 {
			t.Fatal("expected", fmt.Sprintf(">= %d", 1000), "got", int(b))
		}

		if b > 1100 {
			t.Fatal("expected", fmt.Sprintf("<= %d", 1100), "got", int(b))
		}

		{
			dic[b]++
		}
	}

	if len(dic) != 101 {
		t.Fatal("expected", 101, "got", len(dic))
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

	dic := map[int]int{}

	for range 10000 {
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
			Max: 1100,
			Min: 1000,
		})
	}

	{
		go ran.Daemon()
	}

	dic := map[int]int{}

	for range 10000 {
		var b int
		{
			b = ran.backup()
		}

		if b < 1000 {
			t.Fatal("expected", fmt.Sprintf(">= %d", 1000), "got", int(b))
		}

		if b > 1100 {
			t.Fatal("expected", fmt.Sprintf("<= %d", 1100), "got", int(b))
		}

		{
			dic[b]++
		}
	}

	if len(dic) != 101 {
		t.Fatal("expected", 101, "got", len(dic))
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

	dic := map[int]int{}

	for range 10000 {
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
