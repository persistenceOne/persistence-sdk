package types_test

import (
	"strconv"
	"testing"

	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/persistenceOne/persistence-sdk/v4/simapp"
	"github.com/persistenceOne/persistence-sdk/v4/x/epochs/types"
)

type KeeperTestSuite struct {
	simapp.KeeperTestHelper

	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
}

func dummyAfterEpochEndEvent(epochIdentifier string, epochNumber int64) sdk.Event {
	return sdk.NewEvent(
		"afterEpochEnd",
		sdk.NewAttribute("epochIdentifier", epochIdentifier),
		sdk.NewAttribute("epochNumber", strconv.FormatInt(epochNumber, 10)),
	)
}

func dummyBeforeEpochStartEvent(epochIdentifier string, epochNumber int64) sdk.Event {
	return sdk.NewEvent(
		"beforeEpochStart",
		sdk.NewAttribute("epochIdentifier", epochIdentifier),
		sdk.NewAttribute("epochNumber", strconv.FormatInt(epochNumber, 10)),
	)
}

var errDummy = errors.New("9", 9, "dummyError")

// dummyEpochHook is a struct satisfying the epoch hook interface,
// that maintains a counter for how many times its been successfully called,
// and a boolean for whether it should panic during its execution.
type dummyEpochHook struct {
	successCounter int
	shouldPanic    bool
	shouldError    bool
}

func (hook *dummyEpochHook) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	if hook.shouldPanic {
		panic("dummyEpochHook is panicking")
	}

	if hook.shouldError {
		return errDummy
	}
	hook.successCounter++

	ctx.EventManager().EmitEvent(dummyAfterEpochEndEvent(epochIdentifier, epochNumber))

	return nil
}

func (hook *dummyEpochHook) BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64) error {
	if hook.shouldPanic {
		panic("dummyEpochHook is panicking")
	}

	if hook.shouldError {
		return errDummy
	}
	hook.successCounter++

	ctx.EventManager().EmitEvent(dummyBeforeEpochStartEvent(epochIdentifier, epochNumber))

	return nil
}

func (hook *dummyEpochHook) Clone() *dummyEpochHook {
	newHook := dummyEpochHook{shouldPanic: hook.shouldPanic, successCounter: hook.successCounter, shouldError: hook.shouldError}
	return &newHook
}

var _ types.EpochHooks = &dummyEpochHook{}

func (suite *KeeperTestSuite) TestHooksPanicRecovery() {
	panicHook := dummyEpochHook{shouldPanic: true}
	noPanicHook := dummyEpochHook{shouldPanic: false}
	errorHook := dummyEpochHook{shouldError: true}
	noErrorHook := dummyEpochHook{shouldError: false} // same as nopanic
	simpleHooks := []dummyEpochHook{panicHook, noPanicHook, errorHook, noErrorHook}

	tests := []struct {
		hooks                 []dummyEpochHook
		expectedCounterValues []int
		lenEvents             int
	}{
		{[]dummyEpochHook{noPanicHook}, []int{1}, 1},
		{[]dummyEpochHook{panicHook}, []int{0}, 0},
		{[]dummyEpochHook{errorHook}, []int{0}, 0},
		{simpleHooks, []int{0, 1, 0, 1}, 2},
	}

	for tcIndex, tc := range tests {
		for epochActionSelector := 0; epochActionSelector < 2; epochActionSelector++ {
			suite.SetupTest()

			hookRefs := []types.EpochHooks{}

			for _, hook := range tc.hooks {
				hookRefs = append(hookRefs, hook.Clone())
			}

			hooks := types.NewMultiEpochHooks(hookRefs...)

			events := func(epochID string, epochNumber int64, dummyEvent func(id string, number int64) sdk.Event) sdk.Events {
				//nolint:scopelint,testfile
				evts := make(sdk.Events, tc.lenEvents)
				//nolint:	scopelint,testfile
				for i := 0; i < tc.lenEvents; i++ {
					evts[i] = dummyEvent(epochID, epochNumber)
				}

				return evts
			}

			suite.Require().NotPanics(func() {
				//nolint:scopelint,testfile
				if epochActionSelector == 0 {
					err := hooks.BeforeEpochStart(suite.Ctx, "id", 0)
					suite.Require().NoError(err)
					suite.Require().Equal(events("id", 0, dummyBeforeEpochStartEvent), suite.Ctx.EventManager().Events(),
						"test case index %d, before epoch event check", tcIndex)
				} else if epochActionSelector == 1 {
					err := hooks.AfterEpochEnd(suite.Ctx, "id", 0)
					suite.Require().NoError(err)
					suite.Require().Equal(events("id", 0, dummyAfterEpochEndEvent), suite.Ctx.EventManager().Events(),
						"test case index %d, after epoch event check", tcIndex)

				}
			})

			for i := 0; i < len(hooks); i++ {
				epochHook := hookRefs[i].(*dummyEpochHook)
				suite.Require().Equal(tc.expectedCounterValues[i], epochHook.successCounter, "test case index %d", tcIndex)
			}
		}
	}
}
