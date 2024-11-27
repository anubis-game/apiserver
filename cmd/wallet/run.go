package wallet

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/spf13/cobra"
	"github.com/tyler-smith/go-bip39"
	"github.com/xh3b4sd/tracer"
)

type run struct{}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	var ent []byte
	{
		ent, err = bip39.NewEntropy(256)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var mne string
	{
		mne, err = bip39.NewMnemonic(ent)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Println("mnemonic:")
	}

	for i, x := range strings.Fields(mne) {
		fmt.Printf("                %2d    %s\n", i, x)
	}

	var pat string
	{
		pat = hdwallet.DefaultBaseDerivationPath.String()
	}

	var wal *hdwallet.Wallet
	{
		wal, err = hdwallet.NewFromMnemonic(mne)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	var acc accounts.Account
	{
		acc, err = wal.Derive(hdwallet.MustParseDerivationPath(pat), false)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Println("public address:", acc.Address.Hex())
	}

	var key string
	{
		key, err = wal.PrivateKeyHex(acc)
		if err != nil {
			return tracer.Mask(err)
		}
	}

	{
		fmt.Println("private key:   ", key)
	}

	return nil
}
