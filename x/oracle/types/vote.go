package types

import (
	"gopkg.in/yaml.v3"
)

// String implement stringify
func (v AggregateExchangeRatePrevote) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// String implement stringify
func (v AggregateExchangeRateVote) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// String implement stringify
func (v ExchangeRateTuple) String() string {
	out, _ := yaml.Marshal(v)
	return string(out)
}

// ExchangeRateTuples - array of ExchangeRateTuple
type ExchangeRateTuples []ExchangeRateTuple

// String implements fmt.Stringer interface
func (tuples ExchangeRateTuples) String() string {
	out, _ := yaml.Marshal(tuples)
	return string(out)
}
