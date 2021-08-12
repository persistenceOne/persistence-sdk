package send

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/identities/auxiliaries/verify"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/utilities"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
)

type msgServer struct {
	transactionKeeper
}

func (m msgServer) TransactService(goCtx context.Context, msg *message) (helpers.TransactionResponse, error) {
	message := messageFromInterface(msg)
	ctx := types.UnwrapSDKContext(goCtx)
	if auxiliaryResponse := m.transactionKeeper.verifyAuxiliary.GetKeeper().Help(ctx, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError()), nil
	}

	splits := m.transactionKeeper.mapper.NewCollection(ctx)

	if _, Error := utilities.SubtractSplits(splits, message.FromID, message.OwnableID, message.Value); Error != nil {
		return newTransactionResponse(Error), nil
	}

	if _, Error := utilities.AddSplits(splits, message.ToID, message.OwnableID, message.Value); Error != nil {
		return newTransactionResponse(Error), nil
	}

	return newTransactionResponse(nil), nil
}

func NewMsgServerImpl(keeper transactionKeeper) MsgServer {
	return &msgServer{keeper}
}

var _ MsgServer = msgServer{}
