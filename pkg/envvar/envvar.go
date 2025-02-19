package envvar

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Env struct {
	ChainRegistryContract string `split_words:"true" required:"true"`
	ChainRpcEndpoint      string `split_words:"true" required:"true"`
	CodeRepository        string `split_words:"true" required:"true"`
	ConnectionTimeout     string `split_words:"true" required:"true"`
	EngineCapacity        int    `split_words:"true" required:"true"`
	HttpHost              string `split_words:"true" required:"true"`
	HttpPort              string `split_words:"true" required:"true"`
	LogLevel              string `split_words:"true" required:"true"`
	SignerAddress         string `split_words:"true" required:"true"`
	SignerPrivateKey      string `split_words:"true" required:"true"`
}

func Create(kin string) Env {
	var err error

	var fil string
	{
		fil = fmt.Sprintf(".env.%s", kin)
	}

	for {
		{
			err = godotenv.Load(fil)
			if err != nil {
				fmt.Printf("Error: could not load %s (%s)\n", fil, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		var env Env
		{
			err = envconfig.Process("APISERVER", &env)
			if err != nil {
				fmt.Printf("Error: could not process env file %s (%s)\n", fil, err)
				time.Sleep(5 * time.Second)
				continue
			}
		}

		return env
	}
}
