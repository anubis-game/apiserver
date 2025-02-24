package cache

import (
	"fmt"
	"testing"
)

var (
	stringKeys = testKeys[string]{
		Foo: "0x0000000000000000000000000000000000000001",
		Bar: "0x0000000000000000000000000000000000000002",
		Baz: "0x0000000000000000000000000000000000000003",
		Zap: "0x0000000000000000000000000000000000000004",
		Pah: "0x0000000000000000000000000000000000000005",
	}
)

func Test_Cache_String_Create_And_Escape(t *testing.T) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000
		{
			c: NewData[string, int](),
		},
		// Case 001
		{
			c: NewSxnc[string, int](),
		},
		// Case 002
		{
			c: NewSync[string, int](),
		},
		// Case 003
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			createAndEscape(t, tc.c, stringKeys)
		})
	}
}

func Test_Cache_String_Lifecycle(t *testing.T) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000
		{
			c: NewData[string, int](),
		},
		// Case 001
		{
			c: NewSxnc[string, int](),
		},
		// Case 002
		{
			c: NewSync[string, int](),
		},
		// Case 003
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lifecycle(t, tc.c, stringKeys)
		})
	}
}

func Test_Cache_String_Ranger(t *testing.T) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000
		{
			c: NewData[string, int](),
		},
		// Case 001
		{
			c: NewSxnc[string, int](),
		},
		// Case 002
		{
			c: NewSync[string, int](),
		},
		// Case 003
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ranger(t, tc.c, stringKeys)
		})
	}
}

func Test_Cache_String_Read_More_Than_Write(t *testing.T) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000
		{
			c: NewData[string, int](),
		},
		// Case 001
		{
			c: NewSxnc[string, int](),
		},
		// Case 002
		{
			c: NewSync[string, int](),
		},
		// Case 003
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			readMoreThanWrite(t, tc.c, stringKeys)
		})
	}
}

func Test_Cache_String_Exists(t *testing.T) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000
		{
			c: NewData[string, int](),
		},
		// Case 001
		{
			c: NewSxnc[string, int](),
		},
		// Case 002
		{
			c: NewSync[string, int](),
		},
		// Case 003
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			exists(t, tc.c, stringKeys)
		})
	}
}

func Benchmark_Cache_String_Create_And_Escape(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~155 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~117 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~157 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~271 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				createAndEscape(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Lifecycle(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~412 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~312 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~400 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~790 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				lifecycle(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Ranger(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~11,600 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~11,300 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~11,400 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~11,000 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				ranger(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Read_More_Than_Write(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~398,000 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~91,000 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~96,000 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~1,764,000 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				readMoreThanWrite(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Exists(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~171 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~145 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~176 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~286 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				exists(b, tc.c, stringKeys)
			}
		})
	}
}
