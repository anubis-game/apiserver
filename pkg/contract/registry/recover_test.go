package registry

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Test_Registry_Recover(t *testing.T) {
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
			grd: common.HexToAddress("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266"),
			tim: time.Unix(1736079926, 0).UTC(),
			pla: common.HexToAddress("0x90f79bf6eb2c4f870365e785982e1f101e93b906"),
			sgn: hexutil.MustDecode("0xc3849f2ade6d5a10924d98ebd9931194849cdd5199f4c0abd0b058d56cda71f44b891e445d8b1be99fdb22bf654220b540957a527a9552066a31b7f7a52c49631c"),
			sig: common.HexToAddress("0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc"),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%03d", i), func(t *testing.T) {
			var err error

			var sig common.Address
			{
				sig, err = Recover(tc.pre, tc.tim, tc.grd, tc.pla, tc.sgn)
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
