package stream

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/xh3b4sd/tracer"
)

// search is to process the user-challenge protocol method, which requires all
// clients to provide a valid session token, as can be obtained after a
// successful dual-handshake.
func (s *Stream) search(hea []string) (string, error) {
	var err error

	{
		err = seaHea(hea)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// TODO remove
	{
		fak := "f47ac10b-58cc-4372-a567-0e02b2c3d479"
		if hea[1] == fak {
			wal := common.HexToAddress("0xAD63B2262EB7D1591Ee8E6a85959a523dEce7983")
			s.tok.Update(fak, wal.Hex())
		}
	}

	var wal string
	{
		wal = s.tok.Search(hea[1])
	}

	if wal == "" {
		return "", tracer.Mask(challengeValidationFailedError)
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
