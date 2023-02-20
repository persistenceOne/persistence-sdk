package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	PersistenceDenom    string = "uxprt"
	PersistenceSymbol   string = "XPRT"
	PersistenceExponent        = uint32(6)
	BlocksPerMinute            = uint64(10)
	BlocksPerHour              = BlocksPerMinute * 60
	BlocksPerDay               = BlocksPerHour * 24
	BlocksPerWeek              = BlocksPerDay * 7
	BlocksPerMonth             = BlocksPerDay * 30
	BlocksPerYear              = BlocksPerDay * 365
	MicroUnit                  = int64(1e6)

	// TODO: discuss the denom values, took the denoms from: https://docs.osmosis.zone/osmosis-core/asset-info/
	// AtomDenom supported by oracle
	AtomDenom    string = "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
	AtomSymbol   string = "ATOM"
	AtomExponent        = uint32(6)

	// osmosisDenom supported by oracle
	osmosisDenom    string = "uosmo"
	osmosisSymbol   string = "OSMO"
	osmosisExponent        = uint32(6)

	// USDCDenom supported by oracle
	usdcDenom    string = "ibc/D189335C6E4A68B513C10AB227BF1C1D38C746766278BA3EEB4FB14124F1D858"
	usdcSymbol   string = "USDC"
	usdcExponent        = uint32(6)
)

type (
	// ExchangeRatePrevote defines a structure to store a validator's prevote on
	// the rate of USD in the denom asset.
	ExchangeRatePrevote struct {
		Hash        VoteHash       `json:"hash"`         // Vote hex hash to protect centralize data source problem
		Denom       string         `json:"denom"`        // Ticker symbol of denomination exchanged against USD
		Voter       sdk.ValAddress `json:"voter"`        // Voter validator address
		SubmitBlock int64          `json:"submit_block"` // Block height at submission
	}

	// ExchangeRateVote defines a structure to store a validator's vote on the
	// rate of USD in the denom asset.
	ExchangeRateVote struct {
		ExchangeRate sdk.Dec        `json:"exchange_rate"` // Exchange rate of a denomination against USD
		Denom        string         `json:"denom"`         // Ticker symbol of denomination exchanged against USD
		Voter        sdk.ValAddress `json:"voter"`         // Voter validator address
	}

	// VoteHash defines a hash value to hide vote exchange rate which is formatted
	// as a HEX string:
	// SHA256("{salt}:{symbol}:{exchangeRate},...,{symbol}:{exchangeRate}:{voter}")
	VoteHash []byte
)
