# apiserver

```
./apiserver -h
```

```
Golang based apiserver.

Usage:
  apiserver [flags]
  apiserver [command]

Available Commands:
  daemon      Execute the long running process exposing RPC server handlers.
  version     Print version information of this command line tool.
  wallet      Generate Ethereum wallets.

Flags:
  -h, --help   help for apiserver

Use "apiserver [command] --help" for more information about a command.
```

In order to run the apiserver locally you need to add your own `env.local`.

```
APISERVER_CHAIN_REGISTRY_CONTRACT=0x0000000000000000000000000000000000000000
APISERVER_CHAIN_RPC_ENDPOINT=http://127.0.0.1:8545

APISERVER_CODE_REPOSITORY=https://github.com/anubis-game/apiserver

APISERVER_HTTP_HOST=127.0.0.1
APISERVER_HTTP_PORT=7777

APISERVER_LOG_LEVEL=debug

APISERVER_SIGNER_ADDRESS=0x0000000000000000000000000000000000000000
APISERVER_SIGNER_PRIVATE_KEY=0x0000000000000000000000000000000000000000000000000000000000000000
```

Running the apiserver locally with a proper `.env.local` in place.

```
./apiserver daemon
```

```
{ "time":"2024-11-27 20:13:06", "leve":"info", "mess":"server listening for calls", "addr":"127.0.0.1:7777", "call":"/Users/xh3b4sd/project/anubis-game/apiserver/pkg/server/server.go:98" }
```

Generating smart contract bindings using [abigen].

```
abigen --abi pkg/contract/registry/Registry.ABI.json --pkg registry --type RegistryBinding --out pkg/contract/registry/registry_binding.go
```

[abigen]: https://geth.ethereum.org/docs/tools/abigen
