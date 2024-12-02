package wallet

import (
	"regexp"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/xh3b4sd/tracer"
)

var (
	mesexp = regexp.MustCompile(`^signer-0x[0-9a-fA-F]{40}-[0-9]{10,}$`)
	uniexp = regexp.MustCompile(`[0-9]{10,}$`)
)

func Verify(met string, mes string, pub string, sig string, now time.Time) (string, error) {
	{
		if met != "personal_sign" {
			return "", tracer.Mask(signatureMethodInvalidError)
		}
	}

	{
		if len(mes) == 0 {
			return "", tracer.Mask(messageHashEmptyError)
		}
		if !mesexp.MatchString(mes) {
			return "", tracer.Mask(messageHashFormatError)
		}
	}

	{
		if len(pub) == 0 {
			return "", tracer.Mask(publicKeyEmptyError)
		}
		if len(pub) != 132 {
			return "", tracer.Maskf(publicKeyLengthError, "%d", len(pub))
		}
	}

	{
		if now.Sub(mesTim(mes)) > 1*time.Minute {
			return "", tracer.Mask(signatureHashExpiredError)
		}
	}

	{
		if len(sig) == 0 {
			return "", tracer.Mask(signatureHashEmptyError)
		}
		if len(sig) != 132 {
			return "", tracer.Maskf(signatureHashLengthError, "%d", len(sig))
		}
	}

	{
		if !crypto.VerifySignature(pubKey(pub), digHsh(mes), sigHsh(sig)) {
			return "", tracer.Mask(signatureHashInvalidError)
		}
	}

	return walAdd(pub).Hex(), nil
}

func digHsh(mes string) []byte {
	return accounts.TextHash([]byte(mes))
}

func mesTim(mes string) time.Time {
	var err error

	var sub []string
	{
		sub = uniexp.FindStringSubmatch(mes)
	}

	if len(sub) != 1 {
		return time.Time{}
	}

	var uni int64
	{
		uni, err = strconv.ParseInt(sub[0], 10, 64)
		if err != nil {
			return time.Time{}
		}
	}

	return time.Unix(uni, 0).UTC()
}

func pubKey(pub string) []byte {
	dec, err := hexutil.Decode(pub)
	if err != nil {
		return nil
	}

	return dec
}

func sigHsh(sig string) []byte {
	dec, err := hexutil.Decode(sig)
	if err != nil {
		return nil
	}

	return dec[:len(dec)-1]
}

func walAdd(pub string) common.Address {
	poi, err := crypto.UnmarshalPubkey(pubKey(pub))
	if err != nil {
		return common.BytesToAddress(nil)
	}

	return crypto.PubkeyToAddress(*poi)
}
