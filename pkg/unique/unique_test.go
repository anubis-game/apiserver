package unique

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Unique_500(t *testing.T) {
	var uni *Unique[string, int16]
	{
		uni = New[string, int16]()
	}

	for i := 0; i < 384; i++ {
		uni.Ensure(fmt.Sprintf("%d", i))
	}

	var act int16
	{
		act = uni.Ensure(fmt.Sprintf("%d", 385))
	}

	var exp int16
	{
		exp = 384
	}

	if act != exp {
		t.Fatalf("expected %#v got %#v", exp, act)
	}
}

func Test_Unique_Lifecycle(t *testing.T) {
	var uni *Unique[string, byte]
	{
		uni = New[string, byte]()
	}

	for i := 0; i < len(uni.lis); i++ {
		if uni.lis[i] != "" {
			t.Fatalf("expected %#v got %#v", "''", uni.lis[i])
		}
	}

	var au0 byte
	var au1 byte
	var au2 byte
	var au3 byte
	var au4 byte
	{
		au0 = uni.Ensure("0")
		au1 = uni.Ensure("1")
		au2 = uni.Ensure("2")
		au3 = uni.Ensure("3")
		au4 = uni.Ensure("4")
	}

	var eu0 byte
	var eu1 byte
	var eu2 byte
	var eu3 byte
	var eu4 byte
	{
		eu0 = 0x0
		eu1 = 0x1
		eu2 = 0x2
		eu3 = 0x3
		eu4 = 0x4
	}

	if au0 != eu0 {
		t.Fatalf("expected %#v got %#v", eu0, au0)
	}
	if au1 != eu1 {
		t.Fatalf("expected %#v got %#v", eu1, au1)
	}
	if au2 != eu2 {
		t.Fatalf("expected %#v got %#v", eu2, au2)
	}
	if au3 != eu3 {
		t.Fatalf("expected %#v got %#v", eu3, au3)
	}
	if au4 != eu4 {
		t.Fatalf("expected %#v got %#v", eu4, au4)
	}

	// Once allocated, we can call Unique.Ensure() multiple times with the same
	// input and get the same output every time.

	{
		au0 = uni.Ensure("0")
		au2 = uni.Ensure("2")
	}

	if au0 != eu0 {
		t.Fatalf("expected %#v got %#v", eu0, au0)
	}
	if au2 != eu2 {
		t.Fatalf("expected %#v got %#v", eu2, au2)
	}

	// Once allocated, we can call Unique.Ensure() multiple times with the same
	// input and get the same output every time.

	{
		au0 = uni.Ensure("0")
		au2 = uni.Ensure("2")
	}

	if au0 != eu0 {
		t.Fatalf("expected %#v got %#v", eu0, au0)
	}
	if au2 != eu2 {
		t.Fatalf("expected %#v got %#v", eu2, au2)
	}

	{
		uni.Delete("1") // 0x1 is freed
		uni.Delete("3") // 0x3 is freed
	}

	var au5 byte
	var au6 byte
	var au7 byte
	{
		au5 = uni.Ensure("5")
		au6 = uni.Ensure("6")
		au7 = uni.Ensure("7")
	}

	var eu5 byte
	var eu6 byte
	var eu7 byte
	{
		eu5 = 0x1 // "5" gets the freed ID of "1": 0x1
		eu6 = 0x3 // "6" gets the freed ID of "3": 0x3
		eu7 = 0x5 // "7" continues with 0x5
	}

	if au5 != eu5 {
		t.Fatalf("expected %#v got %#v", eu5, au5)
	}
	if au6 != eu6 {
		t.Fatalf("expected %#v got %#v", eu6, au6)
	}
	if au7 != eu7 {
		t.Fatalf("expected %#v got %#v", eu7, au7)
	}

	{
		uni.Delete("4") // 0x4 is freed
		uni.Delete("0") // 0x0 is freed
		uni.Delete("7") // 0x5 is freed
		uni.Delete("6") // 0x3 is freed again
	}

	var au8 byte
	var au9 byte
	var a10 byte
	var a11 byte
	{
		au8 = uni.Ensure("8")
		au9 = uni.Ensure("9")
		a10 = uni.Ensure("10")
		a11 = uni.Ensure("11")
	}

	var eu8 byte
	var eu9 byte
	var e10 byte
	var e11 byte
	{
		eu8 = 0x0 // "8" gets the freed ID of "0": 0x0
		eu9 = 0x3 // "9" gets the freed ID of "3" and "6": 0x3
		e10 = 0x4 // "10" gets the freed ID of "4": 0x4
		e11 = 0x5 // "11" gets the freed ID of "7": 0x5
	}

	if au8 != eu8 {
		t.Fatalf("expected %#v got %#v", eu8, au8)
	}
	if au9 != eu9 {
		t.Fatalf("expected %#v got %#v", eu9, au9)
	}
	if a10 != e10 {
		t.Fatalf("expected %#v got %#v", e10, a10)
	}
	if a11 != e11 {
		t.Fatalf("expected %#v got %#v", e11, a11)
	}

	// Allocating all free UIDs allows us to test for out of range cases below.

	for i := 0; i < len(uni.lis); i++ {
		uni.Ensure(fmt.Sprintf("%03d", i))
	}

	var e12 byte
	var e13 byte
	{
		e12 = 0x0 // "12" is out of range
		e13 = 0x0 // "13" is out of range
	}

	var a12 byte
	var a13 byte
	{
		a12 = uni.Ensure("12")
		a13 = uni.Ensure("13")
	}

	if a12 != e12 {
		t.Fatalf("expected %#v got %#v", e12, a12)
	}
	if a13 != e13 {
		t.Fatalf("expected %#v got %#v", e13, a13)
	}
}

func Test_Unique_Mutex(t *testing.T) {
	// We define 30,000 calls which fits right into the capacity of an int16.

	wrk := 1000
	ops := 30

	var uni *Unique[string, int16]
	{
		uni = New[string, int16]()
	}

	var wai sync.WaitGroup

	var mut sync.Mutex
	see := map[int16]int{}

	for i := 0; i < wrk; i++ {
		{
			wai.Add(1)
		}

		go func(i int) {
			for j := 0; j < ops; j++ {
				var val string
				{
					val = fmt.Sprintf("%03d-%03d", i, j)
				}

				var uid int16
				{
					uid = uni.Ensure(val)
				}

				mut.Lock()
				see[uid]++
				mut.Unlock()

				if j%2 == 0 {
					mut.Lock()
					see[uid]--
					if see[uid] == 0 {
						delete(see, uid)
					}
					mut.Unlock()

					uni.Delete(val)
				}
			}

			{
				wai.Done()
			}
		}(i)
	}

	{
		wai.Wait()
	}

	if uni.length() != len(see) {
		t.Fatalf("expected %#v got %#v", len(see), uni.length())
	}
	for _, v := range see {
		if v != 1 {
			t.Fatalf("expected %#v got %#v", 1, v)
		}
	}
}

func Benchmark_Unique_Cache(b *testing.B) {
	var uni *Unique[string, int8]
	{
		uni = New[string, int8]()
	}

	{
		uni.Ensure("0")
		uni.Ensure("1")
		uni.Ensure("2")
		uni.Ensure("3")
		uni.Ensure("4")
		uni.Ensure("5")
		uni.Ensure("6")
		uni.Ensure("7")
	}

	{
		uni.Delete("4")
		uni.Delete("6")
		uni.Delete("7")
	}

	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			// ~72.30 ns/op
			uni.Ensure("8")
			uni.Ensure("9")

			uni.Delete("8")
			uni.Delete("9")
		}
	})
}

func Benchmark_Unique_Multi(b *testing.B) {
	var uni *Unique[string, int8]
	{
		uni = New[string, int8]()
	}

	{
		uni.Ensure("0")
		uni.Ensure("1")
		uni.Ensure("2")
		uni.Ensure("3")
		uni.Ensure("4")
		uni.Ensure("5")
		uni.Ensure("6")
		uni.Ensure("7")
	}

	{
		uni.Delete("4")
		uni.Delete("6")
		uni.Delete("7")
	}

	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			// ~73.50 ns/op
			uni.Ensure("8")
			uni.Delete("8")

			uni.Ensure("9")
			uni.Delete("9")
		}
	})
}
