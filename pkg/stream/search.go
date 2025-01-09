package stream

import "github.com/xh3b4sd/tracer"

// search is to process the user-challenge protocol method, which requires all
// clients to provide a valid session token, as can be obtained after a
// successful dual-handshake.
func (s *Stream) search(hea []string) (string, error) {
	// TODO verify header format

	var add string
	{
		add = s.tok.Search(hea[1])
	}

	if add == "" {
		return "", tracer.Mask(challengeValidationFailedError)
	}

	return add, nil
}
