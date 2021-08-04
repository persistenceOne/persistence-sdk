package base

import (
	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/persistenceOne/persistenceSDK/schema/test_types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ test_types.Fact = (*fact)(nil)

func (fact fact) GetHashID() test_types.ID             { return fact.HashID }
func (fact fact) GetTypeID() test_types.ID             { return fact.TypeID }
func (fact fact) GetSignatures() test_types.Signatures { return fact.Signatures }
func (fact fact) IsMeta() bool {
	return false
}
func (fact fact) Sign(_ keyring.Keyring) test_types.Fact {
	clicont := client.Context{}
	sign, _, _ := clicont.Keyring.Sign(clicont.FromName, fact.HashID.Bytes())
	Signature := signature{
		ID:             ID{IdString: fact.HashID.String()},
		SignatureBytes: sign,
		ValidityHeight: height{clicont.Height},
	}
	fact.GetSignatures().Add(Signature)

	return fact
}

func NewFact(data types.Data) test_types.Fact {
	return fact{
		HashID:     data.GenerateHashID(),
		TypeID:     data.GetTypeID(),
		Signatures: signatures{},
	}
}

func ReadFact(metaFactString string) (test_types.Fact, error) {
	metaFact, Error := ReadMetaFact(metaFactString)
	if Error != nil {
		return nil, Error
	}

	return metaFact.RemoveData(), nil
}
