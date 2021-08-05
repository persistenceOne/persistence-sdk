package types

type ListData interface {
	Data

	Add(...Data) ListData
	Remove(...Data) ListData

	IsPresent(Data) bool
}
