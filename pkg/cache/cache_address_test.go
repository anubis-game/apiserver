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
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			readMoreThanWrite(t, tc.c, addressKeys)
		})
	}
}

func Benchmark_Address_Cache_Create_And_Escape(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~160 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~130 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~260 ns/op
		{
			c: NewSync[common.Address, int](),
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

func Benchmark_Address_Cache_Lifecycle(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~390 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~300 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~625 ns/op
		{
			c: NewSync[common.Address, int](),
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

func Benchmark_Address_Cache_Ranger(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~885 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~1110 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~1110 ns/op
		{
			c: NewSync[common.Address, int](),
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

func Benchmark_Address_Cache_Read_More_Than_Write(b *testing.B) {
	testCases := []struct {
		c Interface[common.Address, int]
	}{
		// Case 000 ~370,000 ns/op
		{
			c: NewData[common.Address, int](),
		},
		// Case 001 ~69,000 ns/op
		{
			c: NewSxnc[common.Address, int](),
		},
		// Case 002 ~89,000 ns/op
		{
			c: NewSync[common.Address, int](),
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