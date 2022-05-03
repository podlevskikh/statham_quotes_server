package api

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
)

type Specs struct {
	TCPPort string `envconfig:"TCP_PORT" default:"3333"`
}

var specs Specs

func init() {
	err := godotenv.Overload()
	if err != nil {
		fmt.Printf("failed to load .env file %s\n", err)
	}
	err = envconfig.Process("", &specs)
	if err != nil {
		zerolog.DefaultContextLogger.Fatal().Err(err).Msg("missing envs")
	}
}
