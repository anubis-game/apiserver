package stream

import "github.com/xh3b4sd/tracer"

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
