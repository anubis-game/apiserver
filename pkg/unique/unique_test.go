package unique

import (
	"encoding/binary"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func Test_Unique_500(t *testing.T) {
	var len int
	{
		len = 500
	}

	var uni *Unique[string]
	{
		uni = New[string](len)
	}

	for i := 0; i < 384; i++ {
		uni.Ensure(fmt.Sprintf("%d", i))
	}

	var act [2]byte
	{
		act = uni.Ensure(fmt.Sprintf("%d", 385))
	}

	var exp [2]byte
	{
		exp = [2]byte{0x1, 0x80} // 256 + 128 = 384
	}

	if act != exp {
		t.Fatalf("expected %#v got %#v", exp, act)
	}

	var unt uint16
	{
		unt = binary.BigEndian.Uint16(exp[:])
	}

	if unt != 384 {
		t.Fatalf("expected %d got %d", 384, unt)
	}
}

func Test_Unique_Lifecycle(t *testing.T) {
	var len int
	{
		len = 6
	}

	var uni *Unique[string]
	{
		uni = New[string](len)
	}

	for i := 0; i < len; i++ {
		if uni.lis[i] != "" {
			t.Fatalf("expected %#v got %#v", "''", uni.lis[i])
		}
	}

	var au0 [2]byte
	var au1 [2]byte
	var au2 [2]byte
	var au3 [2]byte
	var au4 [2]byte
	{
		au0 = uni.Ensure("0")
		au1 = uni.Ensure("1")
		au2 = uni.Ensure("2")
		au3 = uni.Ensure("3")
		au4 = uni.Ensure("4")
	}

	var eu0 [2]byte
	var eu1 [2]byte
	var eu2 [2]byte
	var eu3 [2]byte
	var eu4 [2]byte
	{
		eu0 = [2]byte{0x0, 0x0}
		eu1 = [2]byte{0x0, 0x1}
		eu2 = [2]byte{0x0, 0x2}
		eu3 = [2]byte{0x0, 0x3}
		eu4 = [2]byte{0x0, 0x4}
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

	var au5 [2]byte
	var au6 [2]byte
	var au7 [2]byte
	var au8 [2]byte
	var au9 [2]byte
	{
		au5 = uni.Ensure("5")
		au6 = uni.Ensure("6")
		au7 = uni.Ensure("7")
		au8 = uni.Ensure("8")
		au9 = uni.Ensure("9")
	}

	var eu5 [2]byte
	var eu6 [2]byte
	var eu7 [2]byte
	var eu8 [2]byte
	var eu9 [2]byte
	{
		eu5 = [2]byte{0x0, 0x1} // "5" gets the freed ID of "1": 0x1
		eu6 = [2]byte{0x0, 0x3} // "6" gets the freed ID of "3": 0x3
		eu7 = [2]byte{0x0, 0x5} // "7" continues with 0x5
		eu8 = [2]byte{0x0, 0x0} // "8" is out of range
		eu9 = [2]byte{0x0, 0x0} // "9" is out of range
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
	if au8 != eu8 {
		t.Fatalf("expected %#v got %#v", eu8, au8)
	}
	if au9 != eu9 {
		t.Fatalf("expected %#v got %#v", eu9, au9)
	}

	{
		uni.Delete("4") // 0x4 is freed
		uni.Delete("0") // 0x0 is freed
		uni.Delete("7") // 0x5 is freed
		uni.Delete("6") // 0x3 is freed again
	}

	var a10 [2]byte
	var a11 [2]byte
	var a12 [2]byte
	var a13 [2]byte
	var a14 [2]byte
	{
		a10 = uni.Ensure("10")
		a11 = uni.Ensure("11")
		a12 = uni.Ensure("12")
		a13 = uni.Ensure("13")
		a14 = uni.Ensure("14")
	}

	var e10 [2]byte
	var e11 [2]byte
	var e12 [2]byte
	var e13 [2]byte
	var e14 [2]byte
	{
		e10 = [2]byte{0x0, 0x0} // "10" gets the freed ID of "0": 0x0
		e11 = [2]byte{0x0, 0x3} // "11" gets the freed ID of "3" and "6": 0x3
		e12 = [2]byte{0x0, 0x4} // "12" gets the freed ID of "4": 0x4
		e13 = [2]byte{0x0, 0x5} // "13" gets the freed ID of "7": 0x5
		e14 = [2]byte{0x0, 0x0} // "14" is out of range again
	}

	if a10 != e10 {
		t.Fatalf("expected %#v got %#v", e10, a10)
	}
	if a11 != e11 {
		t.Fatalf("expected %#v got %#v", e11, a11)
	}
	if a12 != e12 {
		t.Fatalf("expected %#v got %#v", e12, a12)
	}
	if a13 != e13 {
		t.Fatalf("expected %#v got %#v", e13, a13)
	}
	if a14 != e14 {
		t.Fatalf("expected %#v got %#v", e14, a14)
	}
}

func Test_Unique_Mutex(t *testing.T) {
	wrk := 100
	ops := 100
	cap := wrk * ops

	var uni *Unique[string]
	{
		uni = New[string](cap)
	}

	var wai sync.WaitGroup

	var mut sync.Mutex
	see := map[[2]byte]struct{}{}

	var len int32
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

				var uid [2]byte
				{
					uid = uni.Ensure(val)
				}

				mut.Lock()
				_, exi := see[uid]
				if !exi {
					see[uid] = struct{}{}
					atomic.AddInt32(&len, +1)
				}
				mut.Unlock()

				if j%2 == 0 {
					mut.Lock()
					delete(see, uid)
					mut.Unlock()

					uni.Delete(val)
					atomic.AddInt32(&len, -1)
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

	if uni.length() != int(len) {
		t.Fatalf("expected %#v got %#v", len, uni.length())
	}
}

func Benchmark_Unique(b *testing.B) {
	var uni *Unique[string]
	{
		uni = New[string](10)
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
		uni.Delete("2")
		uni.Delete("5")
		uni.Delete("6")
	}

	b.Run(fmt.Sprintf("%03d", 0), func(b *testing.B) {
		b.ResetTimer()
		for range b.N {
			// ~36.50 ns/op for both Unique.Ensure() and Unique.Delete()
			uni.Ensure("8")
			uni.Delete("8")
		}
	})
}
