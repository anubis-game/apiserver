package registry

import (
	"bytes"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_Registry_Verify_recover(t *testing.T) {
	testCases := []struct {
		pre string
		grd common.Address
		tim time.Time
		pla common.Address
		sgn []byte
		sig common.Address
	}{
		// Case 000
		{
			pre: "request",
			grd: common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
			tim: time.Unix(1733916326, 0).UTC(),
			pla: common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"),
			sgn: hexutil.MustDecode("0x0d1290c0d999b3ff7e829a9e6f863413da805ecd608aabc6ccb5b4f6e1e20df1263441548d9519f2ffc40185bd4a9933dc9c7c8c1c06760fcd85bcc421ef057e1c"),
			sig: common.HexToAddress("0x90F79bf6EB2c4f870365E785982E1f101E93b906"),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var sig common.Address
			{
				sig, err = recover(tc.pre, tc.grd, tc.tim, tc.pla, tc.sgn)
				if err != nil {
					t.Fatal(err)
				}
			}

			if !bytes.Equal(sig.Bytes(), tc.sig.Bytes()) {
				t.Fatal("expected", tc.sig.Hex(), "got", sig.Hex())
			}
		})
	}
}

func Test_Registry_Verify_verTim(t *testing.T) {
	testCases := []struct {
		now time.Time
		mes time.Time
		err error
	}{
		// Case 000
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			err: nil,
		},
		// Case 001
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 9, 59, 30, 0, time.UTC),
			err: nil,
		},
		// Case 002
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 9, 59, 0, 0, time.UTC),
			err: nil,
		},
		// Case 003
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 10, 0, 1, 0, time.UTC),
			err: signatureTimeFutureError,
		},
		// Case 004
		{
			now: time.Date(2024, 12, 11, 10, 0, 0, 0, time.UTC),
			mes: time.Date(2024, 12, 11, 9, 58, 59, 0, time.UTC),
			err: signatureTimeExpiredError,
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			err := verTim(tc.now, tc.mes)
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}
