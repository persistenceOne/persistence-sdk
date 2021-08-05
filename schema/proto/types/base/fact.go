package base

import (
	"github.com/99designs/keyring"
	"github.com/cosmos/cosmos-sdk/client"
	protoTypes "github.com/persistenceOne/persistenceSDK/schema/proto/types"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

var _ protoTypes.Fact = (*fact)(nil)

func (fact fact) GetHashID() protoTypes.ID             { return fact.HashID }
func (fact fact) GetTypeID() protoTypes.ID             { return fact.TypeID }
func (fact fact) GetSignatures() protoTypes.Signatures { return fact.Signatures }
func (fact fact) IsMeta() bool {
	return false
}
func (fact fact) Sign(_ keyring.Keyring) protoTypes.Fact {
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

func NewFact(data types.Data) protoTypes.Fact {
	return fact{
		HashID:     data.GenerateHashID(),
		TypeID:     data.GetTypeID(),
		Signatures: signatures{},
	}
}

func ReadFact(metaFactString string) (protoTypes.Fact, error) {
	metaFact, Error := ReadMetaFact(metaFactString)
	if Error != nil {
		return nil, Error
	}

	return metaFact.RemoveData(), nil
}
