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
		// Case 000, ~152 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001, ~104 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002, ~147 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003, ~276 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004, ~117 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				createAndEscape(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Lifecycle(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000, ~420 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001, ~276 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002, ~488 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003, ~786 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004, ~281 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				lifecycle(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Ranger(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000, ~11,600 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001, ~11,200 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002, ~11,400 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003, ~11,200 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004, ~9,000 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				ranger(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Read_More_Than_Write(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000, ~372,000 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001, ~63,000 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002, ~65,000 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003, ~1,656,000 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004, ~62,000 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				readMoreThanWrite(b, tc.c, addressKeys)
			}
		})
	}
}

func Benchmark_Cache_Address_Exists(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000, ~171 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001, ~132 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002, ~164 ns/op
		{
			c: NewSync[common.Address, int](),
		},
		// Case 003, ~296 ns/op
		{
			c: NewCmap[common.Address, int](),
		},
		// Case 004, ~132 ns/op
		{
			c: NewPool[common.Address, int](),
		},
	}

	for i, tc := range testCases {
		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			for b.Loop() {
				exists(b, tc.c, addressKeys)
			}
		})
	}
}
