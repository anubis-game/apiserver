package runtime

var (
	byt []byte = nil
)

var (
	dic = map[string]string{
		"arc": Arc(),
		"gos": Gos(),
		"sha": Sha(),
		"tag": Tag(),
		"ver": Ver(),
	}
)

func init() {
	marshal()
}
