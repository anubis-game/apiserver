package stream

import (
	"bytes"
	"encoding/hex"
	"strings"
	"time"

	"github.com/anubis-game/apiserver/pkg/contract/aggregator"
	"github.com/anubis-game/apiserver/pkg/contract/entrypoint"
	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xh3b4sd/tracer"
)

// verify is to process the dual-handshake protocol method, which requires all
// clients to allocate certain onchain and offchain resources, as well as
// generating several cryptographic signatures.
func (s *Stream) verify(hea []string) (string, error) {
	var err error

	{
		err = verHea(hea)
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// We expect a successfully mined transaction to be provided. This can frankly
	// be any kind of transaction at this point. So once we found a valid
	// transaction using the provided hash, we still have to validate its content
	// further below.
	var txn *types.Transaction
	{
		txn, err = s.reg.Search(common.HexToHash(hea[1]))
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// The second "connect" signature can theoretically be provided by anyone,
	// which is why this signature is not valid in itself. Only if this signature
	// recovers to the same signer as the first signature recovered from the
	// transaction above, only then do we have a verified dual-handshake.
	var sg2 []byte
	{
		sg2, err = hex.DecodeString(strings.TrimPrefix(hea[2], "0x"))
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	// If we end up with a list of valid user operations, then we already know
	// that the provided transaction was managed via ERC-4337 Account Abstraction
	// primitives. Below we need to check every transaction in every user
	// operation bundle, until we can confirm a valid dual handshake.
	var ops []entrypoint.UserOperation
	{
		ops, _, err = entrypoint.Decode(txn.Data())
		if err != nil {
			return "", tracer.Mask(err)
		}
	}

	for _, x := range ops {
		var agg []aggregator.Transaction
		{
			agg, err = aggregator.Decode(x.CallData)
			if err != nil {
				return "", tracer.Mask(err)
			}
		}

		for _, y := range agg {
			// The target of the aggregated transactions inside the user operations
			// represents the smart contract that the initiating user interacted with.
			// This target address must be the address of the Registry smart contract
			// that this Guardian is serving. So if the target address and the
			// Registry address don't match, we try the next transaction, if any.
			if !bytes.Equal(y.Target.Bytes(), s.reg.Address().Bytes()) {
				continue
			}

			// For any transaction made against the Registry smart contract, we decode
			// its calldata to see whether the transaction at hand does in fact
			// represent a contract write to the "Registry.request" method. Only if
			// this is the case, we will find all relevant input data in order to
			// replicate the first "request" and the second "connect" signature.
			var grd common.Address
			var tim time.Time
			var wal common.Address
			var sg1 []byte
			{
				grd, tim, wal, sg1, err = registry.Decode(y.CallData)
				if err != nil {
					return "", tracer.Mask(err)
				}
			}

			// One critical check that we need to conduct here is to ensure that the
			// unix timestamp is still within the valid threshold of +-60 seconds
			// relative to the current time.
			{
				err = verTim(time.Now().UTC(), tim)
				if err != nil {
					return "", tracer.Mask(err)
				}
			}

			// In the relevant ERC-4337 Account Abstraction sense, this is the Player
			// address representing the smart contract wallet enabling gas-sponsorship
			// for all users. This address is the sender of the user operation.
			var pla common.Address
			{
				pla = x.Sender
			}

			var si1 common.Address
			{
				si1, err = registry.Recover("request", tim, grd, pla, sg1)
				if err != nil {
					return "", tracer.Mask(err)
				}
			}

			var si2 common.Address
			{
				si2, err = registry.Recover("connect", tim, grd, pla, sg2)
				if err != nil {
					return "", tracer.Mask(err)
				}
			}

			// At this point we recovered the addresses from the first "request"
			// signature and the second "connect" signature. Only if those two
			// addresses are actually the same, only then have we proven that the
			// requesting user controls all involved wallets, and is therefore
			// considered valid.
			if !bytes.Equal(si1.Bytes(), si2.Bytes()) {
				return "", tracer.Maskf(signerAddressMatchError, "%s != %s", si1.Hex(), si2.Hex())
			}

			// TODO ensure that the Wallet address is not already cached
			//
			// if exi {
			// 	s.log.Log(
			// 		s.ctx,
			// 		"level", "error",
			// 		"message", fmt.Sprintf("Wallet %q added twice", add),
			// 	)
			// }

			// The dual-handshake is valid. We have proven that the returned address
			// represents the user's Wallet address.
			return wal.Hex(), nil
		}
	}

	return "", tracer.Mask(handshakeValidationFailedError)
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

	if len(hea[1]) != 66 {
		return tracer.Maskf(transactionHashLengthError, "%d", len(hea[1]))
	}

	if len(hea[2]) != 132 {
		return tracer.Maskf(signatureHashLengthError, "%d", len(hea[2]))
	}

	return nil
}

func verTim(now time.Time, mes time.Time) error {
	if absDur(now.Sub(mes)) > 60*time.Second {
		return tracer.Mask(signatureTimeInvalidError)
	}

	return nil
}

func absDur(d time.Duration) time.Duration {
	if d < 0 {
		return -d
	}

	return d
}
