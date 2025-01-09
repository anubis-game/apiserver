package release

import (
	"fmt"

	"github.com/anubis-game/apiserver/pkg/contract/registry"
	"github.com/anubis-game/apiserver/pkg/envvar"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cobra"
	"github.com/xh3b4sd/logger"
	"github.com/xh3b4sd/tracer"
)

type run struct {
	flag *flag
}

func (r *run) runE(cmd *cobra.Command, arg []string) error {
	var err error

	var env envvar.Env
	{
		env = envvar.Load(r.flag.Env)
	}

	var log logger.Interface
	{
		log = logger.New(logger.Config{
			Filter: logger.NewLevelFilter(env.LogLevel),
		})
	}

	var reg *registry.Registry
	{
		reg = registry.New(registry.Config{
			Add: env.ChainRegistryContract,
			Key: env.SignerPrivateKey,
			Log: log,
			RPC: env.ChainRpcEndpoint,
		})
	}

	var txn *types.Transaction
	{
		txn, err = reg.Release(common.HexToAddress("0xAD63B2262EB7D1591Ee8E6a85959a523dEce7983"))
		if err != nil {
			return tracer.Mask(err)
		}
	}

	fmt.Printf("%#v\n", txn.Hash().String())

	return nil
}
