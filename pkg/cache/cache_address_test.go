package cache

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

var (
	addressKeys = testKeys[common.Address]{
		Foo: common.HexToAddress("0x0000000000000000000000000000000000000001"),
		Bar: common.HexToAddress("0x0000000000000000000000000000000000000002"),
		Baz: common.HexToAddress("0x0000000000000000000000000000000000000003"),
		Zap: common.HexToAddress("0x0000000000000000000000000000000000000004"),
		Pah: common.HexToAddress("0x0000000000000000000000000000000000000005"),
	}
)

func Test_Cache_Address_Create_And_Escape(t *testing.T) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000
		{
			c: NewData[common.Address, int](),
		},
		// Case 001
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			createAndEscape(t, tc.c, addressKeys)
		})
	}
}

func Test_Cache_Address_Lifecycle(t *testing.T) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000
		{
			c: NewData[common.Address, int](),
		},
		// Case 001
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			lifecycle(t, tc.c, addressKeys)
		})
	}
}

func Test_Cache_Address_Ranger(t *testing.T) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000
		{
			c: NewData[common.Address, int](),
		},
		// Case 001
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			ranger(t, tc.c, addressKeys)
		})
	}
}

func Test_Cache_Address_Read_More_Than_Write(t *testing.T) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000
		{
			c: NewData[common.Address, int](),
		},
		// Case 001
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			readMoreThanWrite(t, tc.c, addressKeys)
		})
	}
}

func Test_Cache_Address_Exists(t *testing.T) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000
		{
			c: NewData[common.Address, int](),
		},
		// Case 001
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			exists(t, tc.c, addressKeys)
		})
	}
}

func Benchmark_Cache_Address_Create_And_Escape(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~152 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~114 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~263 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003 ~334 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004 ~116 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				createAndEscape(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Lifecycle(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~399 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~285 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~649 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003 ~942 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004 ~285 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				lifecycle(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Ranger(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~11,550 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~12,080 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~11,800 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003 ~11,780 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004 ~9,320 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ranger(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Read_More_Than_Write(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~375,000 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~68,000 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~89,000 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003 ~1,651,000 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004 ~67,000 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				readMoreThanWrite(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Exists(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~159 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~140 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~209 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003 ~346 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004 ~142 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				exists(b, tc.c, addressKeys)
			}
		})
	}
}
