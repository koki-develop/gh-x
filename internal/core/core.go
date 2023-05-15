package core

import (
	"io"
	"os"

	"github.com/cli/go-gh/v2/pkg/api"
)

type X struct {
	config *Config
}

type Config struct {
	Output        io.Writer
	RestClient    *api.RESTClient
	GraphQLClient *api.GraphQLClient
}

func New(cfg *Config) *X {
	return &X{config: cfg}
}

func Default() (*X, error) {
	rest, err := api.DefaultRESTClient()
	if err != nil {
		return nil, err
	}

	gql, err := api.DefaultGraphQLClient()
	if err != nil {
		return nil, err
	}

	return New(&Config{
		Output:        os.Stdout,
		RestClient:    rest,
		GraphQLClient: gql,
	}), nil
}
