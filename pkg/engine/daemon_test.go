package engine

import (
	"testing"
)

// ~294 ns/op
func Benchmark_daemon_dynamic_goroutines(b *testing.B) {
	c := make(chan struct{}, 8)

	for b.Loop() {
		go func() {
			c <- struct{}{}
		}()

		{
			<-c
		}
	}

	close(c)
}

// ~130 ns/op
func Benchmark_daemon_worker_pool_range(b *testing.B) {
	c := make(chan struct{}, 8)

	for range 8 {
		go func() {
			for range c {
			}
		}()
	}

	for b.Loop() {
		c <- struct{}{}
	}

	close(c)
}

// ~255 ns/op
func Benchmark_daemon_worker_pool_select_2(b *testing.B) {
	c := make(chan struct{}, 8)
	d := make(chan struct{})

	for range 8 {
		go func() {
			for {
				select {
				case <-d:
					return
				case <-c:
				}
			}
		}()
	}

	for b.Loop() {
		c <- struct{}{}
	}

	close(d)
}

// ~325 ns/op
func Benchmark_daemon_worker_pool_select_3(b *testing.B) {
	c := make(chan struct{}, 8)
	d := make(chan struct{})
	e := make(chan struct{})

	for range 8 {
		go func() {
			for {
				select {
				case <-e:
					return
				case <-d:
					return
				case <-c:
				}
			}
		}()
	}

	for b.Loop() {
		c <- struct{}{}
	}

	close(e)
}

// ~2 ns/op
func Benchmark_Engine_Daemon_without_goroutine(b *testing.B) {
	foo := func() int {
		return 5
	}

	for b.Loop() {
		foo()
	}
}

// ~225.00 ns/op, 1 allocs/op
func Benchmark_Engine_Daemon_with_goroutine(b *testing.B) {
	foo := func() int {
		return 5
	}

	for b.Loop() {
		go foo()
	}
}
