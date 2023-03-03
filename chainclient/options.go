package chainclient

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/credentials"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Options struct {
	GasPrices string
	TLSCert   credentials.TransportCredentials
}

type ClientOption func(opts *Options) error

func DefaultOptions() *Options {
	return &Options{}
}

func OptionGasPrices(gasPrices string) ClientOption {
	return func(opts *Options) error {
		_, err := sdk.ParseDecCoins(gasPrices)
		if err != nil {
			err = errors.Wrapf(err, "failed to ParseDecCoins %s", gasPrices)
			return err
		}

		opts.GasPrices = gasPrices
		return nil
	}
}

func OptionTLSCert(tlsCert credentials.TransportCredentials) ClientOption {
	return func(opts *Options) error {
		opts.TLSCert = tlsCert
		return nil
	}
}
