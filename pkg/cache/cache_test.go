package cache

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func Test_Cache_Data_Create_And_Escape(t *testing.T) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000
		{
			c: NewData[int](),
		},
		// Case 001
		{
			c: NewSxnc[int](),
		},
		// Case 002
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			createAndEscape(t, tc.c)
		})
	}
}

func Test_Cache_Data_Lifecycle(t *testing.T) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000
		{
			c: NewData[int](),
		},
		// Case 001
		{
			c: NewSxnc[int](),
		},
		// Case 002
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lifecycle(t, tc.c)
		})
	}
}

func Test_Cache_Data_Ranger(t *testing.T) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000
		{
			c: NewData[int](),
		},
		// Case 001
		{
			c: NewSxnc[int](),
		},
		// Case 002
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ranger(t, tc.c)
		})
	}
}

func Test_Cache_Data_Read_More_Than_Write(t *testing.T) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000
		{
			c: NewData[int](),
		},
		// Case 001
		{
			c: NewSxnc[int](),
		},
		// Case 002
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			readMoreThanWrite(t, tc.c)
		})
	}
}

func Benchmark_Cache_Create_And_Escape(b *testing.B) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000 ~165 ns/op
		{
			c: NewData[int](),
		},
		// Case 001 ~120 ns/op
		{
			c: NewSxnc[int](),
		},
		// Case 002 ~265 ns/op
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				createAndEscape(b, tc.c)
			}
		})
	}
}

func Benchmark_Cache_Lifecycle(b *testing.B) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000 ~365 ns/op
		{
			c: NewData[int](),
		},
		// Case 001 ~295 ns/op
		{
			c: NewSxnc[int](),
		},
		// Case 002 ~650 ns/op
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lifecycle(b, tc.c)
			}
		})
	}
}

func Benchmark_Cache_Ranger(b *testing.B) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000 ~725 ns/op
		{
			c: NewData[int](),
		},
		// Case 001 ~1000 ns/op
		{
			c: NewSxnc[int](),
		},
		// Case 002 ~1000 ns/op
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ranger(b, tc.c)
			}
		})
	}
}

func Benchmark_Cache_Read_More_Than_Write(b *testing.B) {
	testCases := []struct {
		c Interface[int]
	}{
		// Case 000 ~340,000 ns/op
		{
			c: NewData[int](),
		},
		// Case 001 ~65,000 ns/op
		{
			c: NewSxnc[int](),
		},
		// Case 002 ~92,000 ns/op
		{
			c: NewSync[int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				readMoreThanWrite(b, tc.c)
			}
		})
	}
}

func createAndEscape(t Testing, cac Interface[int]) {
	{
		siz := cac.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	{
		exi := cac.Create("foo", 42)
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		exi := cac.Create("foo", 99)
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		siz := cac.Length()
		if siz != 1 {
			t.Fatal("expected", 1, "got", siz)
		}
	}

	{
		val := cac.Search("foo")
		if val != 42 {
			t.Fatal("expected", 42, "got", val)
		}
	}

	{
		val := cac.Escape("foo")
		if val != 42 {
			t.Fatal("expected", 42, "got", val)
		}
	}

	{
		exi := cac.Exists("foo")
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		val := cac.Search("foo")
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		siz := cac.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	{
		val := cac.Escape("baz")
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		siz := cac.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}
}

func lifecycle(t Testing, cac Interface[int]) {
	{
		siz := cac.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	{
		cac.Update("foo", 33)
		cac.Update("bar", 47)
	}

	{
		siz := cac.Length()
		if siz != 2 {
			t.Fatal("expected", 2, "got", siz)
		}
	}

	{
		exi := cac.Exists("foo")
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := cac.Exists("bar")
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := cac.Exists("baz")
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		val := cac.Search("foo")
		if val != 33 {
			t.Fatal("expected", 33, "got", val)
		}
	}

	{
		val := cac.Search("bar")
		if val != 47 {
			t.Fatal("expected", 47, "got", val)
		}
	}

	{
		val := cac.Search("baz")
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		cac.Update("foo", 99)
	}

	{
		siz := cac.Length()
		if siz != 2 {
			t.Fatal("expected", 2, "got", siz)
		}
	}

	{
		exi := cac.Exists("foo")
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := cac.Exists("bar")
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := cac.Exists("baz")
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		val := cac.Search("foo")
		if val != 99 {
			t.Fatal("expected", 99, "got", val)
		}
	}

	{
		val := cac.Search("bar")
		if val != 47 {
			t.Fatal("expected", 47, "got", val)
		}
	}

	{
		val := cac.Search("baz")
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		cac.Delete("bar")
	}

	{
		siz := cac.Length()
		if siz != 1 {
			t.Fatal("expected", 1, "got", siz)
		}
	}

	{
		exi := cac.Exists("foo")
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := cac.Exists("bar")
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		exi := cac.Exists("baz")
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		val := cac.Search("foo")
		if val != 99 {
			t.Fatal("expected", 99, "got", val)
		}
	}

	{
		val := cac.Search("bar")
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		val := cac.Search("baz")
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		siz := cac.Length()
		if siz != 1 {
			t.Fatal("expected", 1, "got", siz)
		}
	}

	{
		cac.Delete("foo")
	}

	{
		siz := cac.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}
}

func ranger(t Testing, cac Interface[int]) {
	{
		siz := cac.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	exp := map[string]int{
		"foo": 42,
		"bar": 99,
		"baz": 123,
	}

	for k, v := range exp {
		cac.Update(k, v)
	}

	{
		siz := cac.Length()
		if siz != 3 {
			t.Fatal("expected", 3, "got", siz)
		}
	}

	act := make(map[string]int)
	cac.Ranger(func(key string, val int) {
		act[key] = val
	})

	if !reflect.DeepEqual(act, exp) {
		t.Fatal("expected", exp, "got", act)
	}

	for k := range exp {
		cac.Delete(k)
	}

	{
		siz := cac.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}
}

func readMoreThanWrite(t Testing, cac Interface[int]) {
	var w sync.WaitGroup

	{
		cac.Update("foo", 33)
		cac.Update("bar", 47)
		cac.Update("baz", 99)
	}

	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 1000; i++ {
			{
				exi := cac.Exists("foo")
				if !exi {
					t.Fatal("expected", true, "got", false)
				}
			}

			{
				exi := cac.Exists("bar")
				if !exi {
					t.Fatal("expected", true, "got", false)
				}
			}

			{
				exi := cac.Exists("baz")
				if !exi {
					t.Fatal("expected", true, "got", false)
				}
			}

			{
				exi := cac.Exists("zap")
				if exi {
					t.Fatal("expected", false, "got", true)
				}
			}

			{
				exi := cac.Exists("pah")
				if exi {
					t.Fatal("expected", false, "got", true)
				}
			}
		}
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		cac.Update("foo", 33)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		cac.Update("bar", 47)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		cac.Update("baz", 99)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 1000; i++ {
			{
				val := cac.Search("foo")
				if val != 33 {
					t.Fatal("expected", 33, "got", val)
				}
			}

			{
				val := cac.Search("bar")
				if val != 47 {
					t.Fatal("expected", 47, "got", val)
				}
			}

			{
				val := cac.Search("baz")
				if val != 99 {
					t.Fatal("expected", 99, "got", val)
				}
			}

			{
				val := cac.Search("zap")
				if val != 0 {
					t.Fatal("expected", 0, "got", val)
				}
			}

			{
				val := cac.Search("pah")
				if val != 0 {
					t.Fatal("expected", 0, "got", val)
				}
			}
		}
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		cac.Update("foo", 33)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		cac.Update("bar", 47)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		cac.Update("baz", 99)
	}()

	w.Wait()
}
