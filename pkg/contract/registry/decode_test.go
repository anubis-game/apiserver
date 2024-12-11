package registry

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_Registry_Decode(t *testing.T) {
	testCases := []struct {
		byt []byte
		grd common.Address
		tim time.Time
		wal common.Address
		sgn []byte
	}{
		// Case 000
		{
			byt: hexutil.MustDecode("0x383cb92700000000000000000000000070997970c51812dc3a010c7d01b50e0d17dc79c800000000000000000000000000000000000000000000000000000000675976a60000000000000000000000003c44cdddb6a900fa2b585dd299e03d12fa4293bc000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000410d1290c0d999b3ff7e829a9e6f863413da805ecd608aabc6ccb5b4f6e1e20df1263441548d9519f2ffc40185bd4a9933dc9c7c8c1c06760fcd85bcc421ef057e1c00000000000000000000000000000000000000000000000000000000000000"),
			grd: common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8"),
			tim: time.Unix(1733916326, 0).UTC(),
			wal: common.HexToAddress("0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC"),
			sgn: hexutil.MustDecode("0x0d1290c0d999b3ff7e829a9e6f863413da805ecd608aabc6ccb5b4f6e1e20df1263441548d9519f2ffc40185bd4a9933dc9c7c8c1c06760fcd85bcc421ef057e1c"),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var reg *Registry
			{
				reg = &Registry{}
			}

			var grd common.Address
			var tim time.Time
			var wal common.Address
			var sgn []byte
			{
				grd, tim, wal, sgn, err = reg.Decode(tc.byt)
				if err != nil {
					t.Fatal(err)
				}
			}

			if !bytes.Equal(grd.Bytes(), tc.grd.Bytes()) {
				t.Fatal("expected", tc.grd.Hex(), "got", grd.Hex())
			}
			if !bytes.Equal(wal.Bytes(), tc.wal.Bytes()) {
				t.Fatal("expected", tc.wal.Hex(), "got", wal.Hex())
			}
			if !tim.Equal(tc.tim) {
				t.Fatal("expected", tc.tim.Unix(), "got", tim.Unix())
			}
			if !bytes.Equal(sgn, tc.sgn) {
				t.Fatal("expected", string(tc.sgn), "got", string(sgn))
			}
		})
	}
}
