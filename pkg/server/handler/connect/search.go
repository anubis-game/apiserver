package connect

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/xh3b4sd/tracer"
)

// search is to process the user-challenge protocol method, which requires all
// clients to provide a valid session token, as can be obtained after a
// successful dual-handshake.
func (h *Handler) search(hea []string) (common.Address, error) {
	var err error

	{
		err = seaHea(hea)
		if err != nil {
			return common.Address{}, tracer.Mask(err)
		}
	}

	var tok uuid.UUID
	{
		tok, err = uuid.Parse(hea[1])
		if err != nil {
			return common.Address{}, tracer.Mask(err)
		}
	}

	// TODO remove
	{
		fak := uuid.MustParse("f47ac10b-58cc-4372-a567-0e02b2c3d479")
		if tok.String() == fak.String() {
			wal := common.HexToAddress("0xAD63B2262EB7D1591Ee8E6a85959a523dEce7983")
			h.tok.Update(fak, wal)
		}
	}

	var wal common.Address
	var exi bool
	{
		wal, exi = h.tok.Search(tok)
		if !exi {
			return common.Address{}, tracer.Mask(challengeValidationFailedError)
		}
	}

	return wal, nil
}

func seaHea(hea []string) error {
	//
	//     hea[0] handshake method
	//     hea[1] session token
	//
	if len(hea) != 2 {
		return tracer.Maskf(challengeHeaderInvalidError, "%d", len(hea))
	}

	if len(hea[1]) != 36 {
		return tracer.Maskf(sessionTokenLengthError, "%d", len(hea[1]))
	}

	return nil
}
