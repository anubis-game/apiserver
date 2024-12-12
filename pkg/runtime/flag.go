package runtime

import "runtime"

var (
	sha = "n/a"
	tag = "n/a"
)

func Arc() string {
	return runtime.GOARCH
}

func Gos() string {
	return runtime.GOOS
}

func Sha() string {
	return sha
}

func Tag() string {
	return tag
}

func Ver() string {
	return runtime.Version()
}
