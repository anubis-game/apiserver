package stream

import (
	"sync"
	"testing"
)

func Test_Stream_Worker(t *testing.T) {
	var wrk *Worker
	{
		wrk = NewWorker()
	}

	dic := map[int]int{}
	mut := sync.Mutex{}
	wai := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		{
			wai.Add(1)
		}

		go wrk.Worker(func() {
			mut.Lock()
			dic[i]++
			mut.Unlock()
			wai.Done()
		})
	}

	{
		wai.Wait()
	}

	if len(dic) != 100 {
		t.Fatal("expected", 100, "got", len(dic))
	}

	for k, v := range dic {
		if v < 1 {
			t.Fatalf("expected %d to appear once, got %d", k, v)
		}
	}
}
