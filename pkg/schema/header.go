package schema

type Header string

const (
	// DualHandshake defines the verification challenge that must be overcome, for
	// every stream engine client wanting to participate in the realtime services
	// provided. In case the dual-handshake was successfully executed, onchain and
	// offchain resources may be allocated for the authenticated user.
	DualHandshake Header = "dual-handshake"
	// UserChallenge defines the verification method enabled after a successful
	// dual-handshake was executed and a valid session token was exchanged for the
	// authenticated user.
	UserChallenge Header = "user-challenge"
)
