package unique

import (
	"fmt"
	"slices"
	"sync"
	"testing"
)

func Test_Unique_500(t *testing.T) {
	var uni *Unique[string, int16]
	{
		uni = New[string, int16]()
	}

	for i := range 384 {
		uni.Ensure(fmt.Sprintf("%03d", i))
	}

	var act int16
	{
		act = uni.Ensure(fmt.Sprintf("%d", 385))
	}

	if act != 384 {
		t.Fatalf("expected %#v got %#v", 384, act)
	}
}

func Test_Unique_Lifecycle(t *testing.T) {
	var uni *Unique[string, byte]
	{
		uni = New[string, byte]()
	}

	if uni.Length() != 0 {
		t.Fatalf("expected %#v got %#v", 0, uni.Length())
	}

	var au0 byte
	{
		au0 = uni.Ensure("000")
	}

	if uni.Length() != 1 {
		t.Fatalf("expected %#v got %#v", 1, uni.Length())
	}

	var au1 byte
	var au2 byte
	var au3 byte
	var au4 byte
	{
		au1 = uni.Ensure("001")
		au2 = uni.Ensure("002")
		au3 = uni.Ensure("003")
		au4 = uni.Ensure("004")
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

	var all []byte
	{
		all = []byte{eu0, eu1, eu2, eu3, eu4}
	}

	if uni.Length() != 5 {
		t.Fatalf("expected %#v got %#v", 5, uni.Length())
	}
	for x := range uni.Length() {
		if !slices.Contains(all, x) {
			t.Fatalf("expected %#v got %#v", eu0, au0)
		}
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
		au0 = uni.Ensure("000")
		au2 = uni.Ensure("002")
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
		au0 = uni.Ensure("000")
		au2 = uni.Ensure("002")
	}

	if au0 != eu0 {
		t.Fatalf("expected %#v got %#v", eu0, au0)
	}
	if au2 != eu2 {
		t.Fatalf("expected %#v got %#v", eu2, au2)
	}

	{
		uni.Delete("001") // 0x1 is freed
		uni.Delete("003") // 0x3 is freed
	}

	var au5 byte
	var au6 byte
	var au7 byte
	{
		au5 = uni.Ensure("005")
		au6 = uni.Ensure("006")
		au7 = uni.Ensure("007")
	}

	var eu5 byte
	var eu6 byte
	var eu7 byte
	{
		eu5 = 0x3 // "005" gets the freed ID of "003": 0x3
		eu6 = 0x1 // "006" gets the freed ID of "001": 0x1
		eu7 = 0x5 // "007" continues with 0x5
	}

	{
		all = []byte{eu0, eu2, eu4, eu5, eu6, eu7}
	}

	if uni.Length() != 6 {
		t.Fatalf("expected %#v got %#v", 6, uni.Length())
	}
	for x := range uni.Length() {
		if !slices.Contains(all, x) {
			t.Fatalf("expected %#v got %#v", eu0, au0)
		}
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
		uni.Delete("004") // 0x4 is freed
		uni.Delete("000") // 0x0 is freed
		uni.Delete("007") // 0x5 is freed
		uni.Delete("006") // 0x1 is freed again
	}

	var eu8 byte
	var eu9 byte
	var e10 byte
	var e11 byte
	{
		eu8 = 0x1 // "008" gets the freed ID of "006": 0x1
		eu9 = 0x5 // "009" gets the freed ID of "007": 0x5
		e10 = 0x0 // "010" gets the freed ID of "000": 0x0
		e11 = 0x4 // "011" gets the freed ID of "004": 0x4
	}

	var au8 byte
	var au9 byte
	var a10 byte
	var a11 byte
	{
		au8 = uni.Ensure("008")
		au9 = uni.Ensure("009")
		a10 = uni.Ensure("010")
		a11 = uni.Ensure("011")
	}

	{
		all = []byte{eu2, eu5, eu8, eu9, e10, e11}
	}

	if uni.Length() != 6 {
		t.Fatalf("expected %#v got %#v", 6, uni.Length())
	}
	for x := range uni.Length() {
		if !slices.Contains(all, x) {
			t.Fatalf("expected %#v got %#v", eu0, au0)
		}
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

	for i := range len(uni.lis) {
		uni.Ensure(fmt.Sprintf("flood-%03d", i))
	}

	var e12 byte
	var e13 byte
	{
		e12 = 0x0 // "012" is out of range
		e13 = 0x0 // "013" is out of range
	}

	var a12 byte
	var a13 byte
	{
		a12 = uni.Ensure("012")
		a13 = uni.Ensure("013")
	}

	if uni.Length() != 255 {
		t.Fatalf("expected %#v got %#v", 255, uni.Length())
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

	for i := range wrk {
		{
			wai.Add(1)
		}

		go func(i int) {
			for j := range ops {
				var key string
				{
					key = fmt.Sprintf("%03d", i)
				}

				var exi bool
				var uid int16
				{
					exi = uni.Exists(key)
					uid = uni.Ensure(key)
				}

				if !exi {
					mut.Lock()
					see[uid]++
					mut.Unlock()
				}

				if j%2 == 0 {
					mut.Lock()
					see[uid]--
					if see[uid] == 0 {
						delete(see, uid)
					}
					mut.Unlock()

					uni.Delete(key)
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

	if int(uni.Length()) != len(see) {
		t.Fatalf("expected %#v got %#v", len(see), uni.Length())
	}
	for _, v := range see {
		if v != 1 {
			t.Fatalf("expected %#v got %#v", 1, v)
		}
	}
	for x := range uni.Length() {
		_, e := see[x]
		if !e {
			t.Fatalf("expected %#v got %#v", true, false)
		}
	}
}

// ~0.11 ns/op
func Benchmark_Unique_concurrency(b *testing.B) {
	n := 1_000_000
	k := "key"
	u := New[string, int16]()

	w := sync.WaitGroup{}
	c := make(chan struct{})

	go func() {
		<-c

		for range n {
			u.Delete(k)
		}

		w.Done()
	}()

	go func() {
		<-c

		for range n {
			u.Ensure(k)
		}

		w.Done()
	}()

	go func() {
		<-c

		for range n {
			u.Exists(k)
		}

		w.Done()
	}()

	go func() {
		<-c

		for range n {
			u.Length()
		}

		w.Done()
	}()

	w.Add(4)
	close(c)
	w.Wait()
}
