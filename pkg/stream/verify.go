package stream

import (
	"encoding/hex"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/tracer"
)

func (s *Stream) verify(hea []string) (string, error) {
	var err error

	{
		err = verHea(hea)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var txn *types.Transaction
	{
		txn, err = s.reg.Search(common.HexToHash(hea[1]))
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var pla common.Address
	{
		pla, err = s.reg.Sender(txn)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var grd common.Address
	var tim time.Time
	var wal common.Address
	var sg1 []byte
	{
		grd, tim, wal, sg1, err = s.reg.Decode(txn.Data())
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	var sg2 []byte
	{
		sg2, err = hex.DecodeString(strings.TrimPrefix(hea[2], "0x"))
		if err != nil {
			tracer.Panic(err)
		}
	}

	{
		err = s.reg.Verify(grd, tim, pla, sg1, sg2)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	return wal.Hex(), nil
}

func verHea(hea []string) error {
	//
	//     hea[0] handshake method
	//     hea[1] transaction hash
	//     hea[2] signature hash
	//
	if len(hea) != 3 {
		return tracer.Maskf(handshakeHeaderInvalidError, "%d", len(hea))
	}

	if hea[0] != "dual-handshake" {
		return tracer.Mask(handshakeMethodInvalidError)
	}

	if len(hea[1]) != 66 {
		return tracer.Maskf(transactionHashLengthError, "%d", len(hea[1]))
	}

	if len(hea[2]) != 132 {
		return tracer.Maskf(signatureHashLengthError, "%d", len(hea[2]))
	}

	return nil
}
