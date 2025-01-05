package registry

import (
	"crypto/ecdsa"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xh3b4sd/tracer"
)

func Recover(pre string, tim time.Time, grd common.Address, pla common.Address, sgn []byte) (common.Address, error) {
	var err error

	var msg string
	{
		msg = strings.Join([]string{
			pre,
			strconv.FormatInt(tim.Unix(), 10),
			strings.ToLower(grd.Hex()),
			strings.ToLower(pla.Hex()),
		}, "-")
	}

	var byt []byte
	{
		byt, err = crypto.Ecrecover(digHsh(msg), sgnHsh(sgn))
		if err != nil {
			return common.Address{}, tracer.Mask(signatureHashInvalidError)
		}
	}

	var pub *ecdsa.PublicKey
	{
		pub, err = crypto.UnmarshalPubkey(byt)
		if err != nil {
			return common.Address{}, tracer.Mask(err)
		}
	}

	return crypto.PubkeyToAddress(*pub), nil
}

func digHsh(mes string) []byte {
	return accounts.TextHash([]byte(mes))
}

func sgnHsh(sgn []byte) []byte {
	// Adjust the Ethereum-style recovery ID "V" (27 or 28) to standard (0 or 1).
	if sgn[64] >= 27 {
		sgn[64] -= 27
	}

	return sgn
}
