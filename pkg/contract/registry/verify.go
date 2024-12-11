package registry

import (
	"bytes"
	"crypto/ecdsa"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xh3b4sd/tracer"
)

func (r *Registry) Verify(grd common.Address, tim time.Time, pla common.Address, sg1 []byte, sg2 []byte) error {
	var err error

	{
		if !bytes.Equal(grd.Bytes(), r.add.Bytes()) {
			return tracer.Maskf(guardianAddressMatchError, "%s != %s", grd.Hex(), r.add.Hex())
		}
	}

	{
		err = verTim(time.Now().UTC(), tim)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var si1 common.Address
	{
		si1, err = recover("request", grd, tim, pla, sg1)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var si2 common.Address
	{
		si2, err = recover("connect", grd, tim, pla, sg2)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		if !bytes.Equal(si1.Bytes(), si2.Bytes()) {
			return tracer.Maskf(signerAddressMatchError, "%s != %s", si1.Hex(), si2.Hex())
		}
	}

	return nil
}

func recover(pre string, grd common.Address, tim time.Time, pla common.Address, sgn []byte) (common.Address, error) {
	var err error

	var msg string
	{
		msg = strings.Join([]string{
			pre,
			strings.ToLower(grd.Hex()),
			strconv.FormatInt(tim.Unix(), 10),
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

func verTim(now time.Time, mes time.Time) error {
	if now.Sub(mes) > 60*time.Second {
		return tracer.Mask(signatureTimeExpiredError)
	}

	if mes.After(now) {
		return tracer.Mask(signatureTimeFutureError)
	}

	return nil
}
