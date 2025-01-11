package cache

import (
	"reflect"
	"sync"
)

type testKeys[T comparable] struct {
	Foo T
	Bar T
	Baz T
	Zap T
	Pah T
}

func createAndEscape[K comparable, V int](t Testing, c Interface[K, V], k testKeys[K]) {
	{
		siz := c.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	{
		exi := c.Create(k.Foo, 42)
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		exi := c.Create(k.Foo, 99)
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		siz := c.Length()
		if siz != 1 {
			t.Fatal("expected", 1, "got", siz)
		}
	}

	{
		val, exi := c.Search(k.Foo)
		if val != 42 {
			t.Fatal("expected", 42, "got", val)
		}
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		val := c.Escape(k.Foo)
		if val != 42 {
			t.Fatal("expected", 42, "got", val)
		}
	}

	{
		_, exi := c.Search(k.Foo)
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		siz := c.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	{
		val := c.Escape(k.Baz)
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		siz := c.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}
}

func lifecycle[K comparable, V int](t Testing, c Interface[K, V], k testKeys[K]) {
	{
		siz := c.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	{
		c.Update(k.Foo, 33)
		c.Update(k.Bar, 47)
	}

	{
		siz := c.Length()
		if siz != 2 {
			t.Fatal("expected", 2, "got", siz)
		}
	}

	{
		exi := c.Exists(k.Foo)
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := c.Exists(k.Bar)
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := c.Exists(k.Baz)
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		val, _ := c.Search(k.Foo)
		if val != 33 {
			t.Fatal("expected", 33, "got", val)
		}
	}

	{
		val, _ := c.Search(k.Bar)
		if val != 47 {
			t.Fatal("expected", 47, "got", val)
		}
	}

	{
		val, _ := c.Search(k.Baz)
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		c.Update(k.Foo, 99)
	}

	{
		siz := c.Length()
		if siz != 2 {
			t.Fatal("expected", 2, "got", siz)
		}
	}

	{
		exi := c.Exists(k.Foo)
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := c.Exists(k.Bar)
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := c.Exists(k.Baz)
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		val, _ := c.Search(k.Foo)
		if val != 99 {
			t.Fatal("expected", 99, "got", val)
		}
	}

	{
		val, _ := c.Search(k.Bar)
		if val != 47 {
			t.Fatal("expected", 47, "got", val)
		}
	}

	{
		val, _ := c.Search(k.Baz)
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		c.Delete(k.Bar)
	}

	{
		siz := c.Length()
		if siz != 1 {
			t.Fatal("expected", 1, "got", siz)
		}
	}

	{
		exi := c.Exists(k.Foo)
		if !exi {
			t.Fatal("expected", true, "got", false)
		}
	}

	{
		exi := c.Exists(k.Bar)
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		exi := c.Exists(k.Baz)
		if exi {
			t.Fatal("expected", false, "got", true)
		}
	}

	{
		val, _ := c.Search(k.Foo)
		if val != 99 {
			t.Fatal("expected", 99, "got", val)
		}
	}

	{
		val, _ := c.Search(k.Bar)
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		val, _ := c.Search(k.Baz)
		if val != 0 {
			t.Fatal("expected", 0, "got", val)
		}
	}

	{
		siz := c.Length()
		if siz != 1 {
			t.Fatal("expected", 1, "got", siz)
		}
	}

	{
		c.Delete(k.Foo)
	}

	{
		siz := c.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}
}

func ranger[K comparable, V int](t Testing, c Interface[K, V], k testKeys[K]) {
	{
		siz := c.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}

	exp := map[K]V{
		k.Foo: 42,
		k.Bar: 99,
		k.Baz: 123,
	}

	for k, v := range exp {
		c.Update(k, v)
	}

	{
		siz := c.Length()
		if siz != 3 {
			t.Fatal("expected", 3, "got", siz)
		}
	}

	act := make(map[K]V)
	c.Ranger(func(key K, val V) {
		act[key] = val
	})

	if !reflect.DeepEqual(act, exp) {
		t.Fatal("expected", exp, "got", act)
	}

	for k := range exp {
		c.Delete(k)
	}

	{
		siz := c.Length()
		if siz != 0 {
			t.Fatal("expected", 0, "got", siz)
		}
	}
}

func readMoreThanWrite[K comparable, V int](t Testing, c Interface[K, V], k testKeys[K]) {
	var w sync.WaitGroup

	{
		c.Update(k.Foo, 33)
		c.Update(k.Bar, 47)
		c.Update(k.Baz, 99)
	}

	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 1000; i++ {
			{
				exi := c.Exists(k.Foo)
				if !exi {
					t.Fatal("expected", true, "got", false)
				}
			}

			{
				exi := c.Exists(k.Bar)
				if !exi {
					t.Fatal("expected", true, "got", false)
				}
			}

			{
				exi := c.Exists(k.Baz)
				if !exi {
					t.Fatal("expected", true, "got", false)
				}
			}

			{
				exi := c.Exists(k.Zap)
				if exi {
					t.Fatal("expected", false, "got", true)
				}
			}

			{
				exi := c.Exists(k.Pah)
				if exi {
					t.Fatal("expected", false, "got", true)
				}
			}
		}
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		c.Update(k.Foo, 33)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		c.Update(k.Bar, 47)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		c.Update(k.Baz, 99)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		for i := 0; i < 1000; i++ {
			{
				val, _ := c.Search(k.Foo)
				if val != 33 {
					t.Fatal("expected", 33, "got", val)
				}
			}

			{
				val, _ := c.Search(k.Bar)
				if val != 47 {
					t.Fatal("expected", 47, "got", val)
				}
			}

			{
				val, _ := c.Search(k.Baz)
				if val != 99 {
					t.Fatal("expected", 99, "got", val)
				}
			}

			{
				val, _ := c.Search(k.Zap)
				if val != 0 {
					t.Fatal("expected", 0, "got", val)
				}
			}

			{
				val, _ := c.Search(k.Pah)
				if val != 0 {
					t.Fatal("expected", 0, "got", val)
				}
			}
		}
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		c.Update(k.Foo, 33)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		c.Update(k.Bar, 47)
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		c.Update(k.Baz, 99)
	}()

	w.Wait()
}
