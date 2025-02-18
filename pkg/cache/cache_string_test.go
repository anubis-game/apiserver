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
		// Case 000, ~170 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~135 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~270 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~332 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				createAndEscape(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Lifecycle(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~365 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~330 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~675 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~918 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lifecycle(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Ranger(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~750 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~1000 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~1000 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~950 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ranger(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Read_More_Than_Write(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~360,000 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~70,000 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~95,000 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~1,522,000 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				readMoreThanWrite(b, tc.c, stringKeys)
			}
		})
	}
}

func Benchmark_Cache_String_Exists(b *testing.B) {
	testCases := []struct {
		c Interface[string, int]
	}{
		// Case 000, ~155 ns/op
		{
			c: NewData[string, int](),
		},
		// Case 001, ~144 ns/op
		{
			c: NewSxnc[string, int](),
		},
		// Case 002, ~229 ns/op
		{
			c: NewSync[string, int](),
		},
		// Case 003, ~335 ns/op
		{
			c: NewCmap[string, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				exists(b, tc.c, stringKeys)
			}
		})
	}
}
