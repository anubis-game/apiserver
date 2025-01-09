package registry

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type Config struct {
	// Add is the address of the deployed smart contract to interact with.
	Add string
	// Key is the private key signing transactions for contract writes.
	Key string
	// Log is a simple logger interface to print system relevant information.
	Log logger.Interface
	// RPC is the RPC endpoint for network connection.
	RPC string
}

type Registry struct {
	add common.Address
	bin *RegistryBinding
	cid *big.Int
	cli *ethclient.Client
	log logger.Interface
	opt *bind.TransactOpts
	sig types.Signer
}

func New(c Config) *Registry {
	if c.Add == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Add must not be empty", c)))
	}
	if c.Key == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Key must not be empty", c)))
	}
	if c.Log == nil {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.Log must not be empty", c)))
	}
	if c.RPC == "" {
		tracer.Panic(tracer.Mask(fmt.Errorf("%T.RPC must not be empty", c)))
	}

	var err error

	var add common.Address
	{
		add = common.HexToAddress(c.Add)
	}

	var cli *ethclient.Client
	{
		cli, err = ethclient.Dial(c.RPC)
		if err != nil {
			tracer.Panic(err)
		}
	}

	var cid *big.Int
	{
		cid, err = cli.ChainID(context.Background())
		if err != nil {
			tracer.Panic(err)
		}
	}

	var bin *RegistryBinding
	{
		bin, err = NewRegistryBinding(add, cli)
		if err != nil {
			tracer.Panic(err)
		}
	}

	var key *ecdsa.PrivateKey
	{
		key, err = crypto.HexToECDSA(strings.TrimPrefix(c.Key, "0x"))
		if err != nil {
			tracer.Panic(err)
		}
	}

	var opt *bind.TransactOpts
	{
		opt, err = bind.NewKeyedTransactorWithChainID(key, cid)
		if err != nil {
			tracer.Panic(err)
		}
	}

	var sig types.Signer
	{
		sig = types.NewCancunSigner(cid)
	}

	return &Registry{
		add: add,
		bin: bin,
		cid: cid,
		cli: cli,
		log: c.Log,
		opt: opt,
		sig: sig,
	}
}

func (r *Registry) writerOption() *bind.TransactOpts {
	return &bind.TransactOpts{
		From: r.opt.From,

		// Here we are trying to set some reasonable gas limits, specifically for
		// the EIP-1559 enabled minting transaction.
		//
		//     GasFeeCap is the max gas fee we are willing to pay
		//     GasTipCap is the max priority fee we are willing to pay
		//
		// Below is a testnet transaction providing some real world insight into
		// effective gas usage.
		//
		//     TODO
		//
		GasFeeCap: big.NewInt(5_000_000_000), // 5.00 gwei
		GasTipCap: big.NewInt(500_000_000),   // 0.50 gwei

		Signer: r.opt.Signer,
	}
}
