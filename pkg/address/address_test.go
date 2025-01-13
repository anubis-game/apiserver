package address

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func Benchmark_Address_Empty_True(b *testing.B) {
	testCases := []struct {
		f func(common.Address) bool
	}{
		// Case 000 ~3.50 ns/op
		{
			f: func(a common.Address) bool {
				return bytes.Equal(a.Bytes(), zeroAddress.Bytes())
			},
		},
		// Case 001 ~3.39 ns/op
		{
			f: func(a common.Address) bool {
				return a.Cmp(zeroAddress) == 0
			},
		},
		// Case 002 ~3.10 ns/op
		{
			f: Empty,
		},
	}

	for i, tc := range testCases {
		var add common.Address
		{
			add = common.Address{}
		}

		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if !tc.f(add) {
					b.Fatal("expected", true, "got", false)
				}
			}
		})
	}
}

func Benchmark_Address_Empty_False(b *testing.B) {
	testCases := []struct {
		f func(common.Address) bool
	}{
		// Case 000 ~2.59 ns/op
		{
			f: func(a common.Address) bool {
				return bytes.Equal(a.Bytes(), zeroAddress.Bytes())
			},
		},
		// Case 001 ~3.08 ns/op
		{
			f: func(a common.Address) bool {
				return a.Cmp(zeroAddress) == 0
			},
		},
		// Case 002 ~2.30 ns/op
		{
			f: Empty,
		},
	}

	for i, tc := range testCases {
		var add common.Address
		{
			add = common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		}

		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if tc.f(add) {
					b.Fatal("expected", false, "got", true)
				}
			}
		})
	}
}

func Benchmark_Address_Equal_True(b *testing.B) {
	testCases := []struct {
		f func(common.Address, common.Address) bool
	}{
		// Case 000 ~3.47 ns/op
		{
			f: func(a common.Address, b common.Address) bool {
				return bytes.Equal(a.Bytes(), b.Bytes())
			},
		},
		// Case 001 ~3.30 ns/op
		{
			f: func(a common.Address, b common.Address) bool {
				return a.Cmp(b) == 0
			},
		},
		// Case 002 ~3.10 ns/op
		{
			f: Equal,
		},
	}

	for i, tc := range testCases {
		var one common.Address
		{
			one = common.HexToAddress("0x7413685672890ab09ab1235345c4d67c2f8edef1")
		}

		var two common.Address
		{
			two = common.HexToAddress("0x7413685672890ab09ab1235345c4d67c2f8edef1")
		}

		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if !tc.f(one, two) {
					b.Fatal("expected", true, "got", false)
				}
			}
		})
	}
}

func Benchmark_Address_Equal_False(b *testing.B) {
	testCases := []struct {
		f func(common.Address, common.Address) bool
	}{
		// Case 000 ~2.90 ns/op
		{
			f: func(a common.Address, b common.Address) bool {
				return bytes.Equal(a.Bytes(), b.Bytes())
			},
		},
		// Case 001 ~2.96 ns/op
		{
			f: func(a common.Address, b common.Address) bool {
				return a.Cmp(b) == 0
			},
		},
		// Case 002 ~2.60 ns/op
		{
			f: Equal,
		},
	}

	for i, tc := range testCases {
		var one common.Address
		{
			one = common.HexToAddress("0x7413685672890ab09ab1235345c4d67c2f8edef1")
		}

		var two common.Address
		{
			two = common.HexToAddress("0x1234567890abcdef1234567890abcdef12345678")
		}

		b.Run(fmt.Sprintf("%03d", i), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if tc.f(one, two) {
					b.Fatal("expected", false, "got", true)
				}
			}
		})
	}
}
