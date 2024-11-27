# apiserver

In order to run the apiserver locally you need to add your own `env.local`.

```
APISERVER_CHAIN_REGISTRY_CONTRACT=TODO
APISERVER_CHAIN_RPC_ENDPOINT=http://127.0.0.1:8545

APISERVER_HTTP_HOST=127.0.0.1
APISERVER_HTTP_PORT=7777

APISERVER_LOG_LEVEL=debug

APISERVER_SIGNER_PRIVATE_KEY=TODO
```

```
./apiserver daemon
```

```
{ "time":"2024-11-27 20:13:06", "leve":"info", "mess":"server listening for calls", "addr":"127.0.0.1:7777", "call":"/Users/xh3b4sd/project/anubis-game/apiserver/pkg/server/server.go:98" }
```
