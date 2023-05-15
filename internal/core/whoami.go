package core

import "fmt"

func (x *X) Whoami() error {
	rslt, err := x.whoami()
	if err != nil {
		return err
	}

	if _, err := fmt.Fprintln(x.config.Output, rslt.Viewer.Login); err != nil {
		return err
	}

	return nil
}

type whoamiResult struct {
	Viewer struct {
		Login string
	}
}

func (x *X) whoami() (*whoamiResult, error) {
	var rslt whoamiResult
	if err := x.config.GraphQLClient.Query("whoami", &rslt, nil); err != nil {
		return nil, err
	}

	return &rslt, nil
}
