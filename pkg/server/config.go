package server

import "github.com/anubis-game/apiserver/pkg/runtime"

func (s *Server) config() {
	runtime.With("grd", s.env.SignerAddress)
	runtime.With("reg", s.env.ChainRegistryContract)
	runtime.With("src", s.env.CodeRepository)
}
