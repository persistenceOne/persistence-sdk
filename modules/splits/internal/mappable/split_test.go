package mappable

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/splits/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Split_Methods(t *testing.T) {

	ownerID := base.NewID("ownerID")
	ownableID := base.NewID("ownableID")

	testSplitID := key.NewSplitID(ownerID, ownableID)
	splitDec := sdkTypes.NewDec(12)
	testSplit := NewSplit(testSplitID, splitDec).(split)

	require.Equal(t, split{ID: testSplitID, Split: splitDec}, testSplit)
	require.Equal(t, testSplitID, testSplit.GetID())
	require.Equal(t, ownerID, testSplit.GetOwnerID())
	require.Equal(t, ownableID, testSplit.GetOwnableID())
	require.Equal(t, splitDec, testSplit.GetSplit())
	require.Equal(t, NewSplit(testSplitID, sdkTypes.NewDec(11)).(split), testSplit.Send(sdkTypes.NewDec(1)))
	require.Equal(t, NewSplit(testSplitID, sdkTypes.NewDec(13)).(split), testSplit.Receive(sdkTypes.NewDec(1)))
	require.Equal(t, true, testSplit.CanSend(sdkTypes.NewDec(5)))
	require.Equal(t, false, testSplit.CanSend(sdkTypes.NewDec(15)))
	require.Equal(t, testSplitID, testSplit.GetKey())
}