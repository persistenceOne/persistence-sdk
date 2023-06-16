package osmosisibctesting

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	govv1beta1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	transfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

var (
	oneAddress sdk.AccAddress = bytes.Repeat([]byte{0x1}, wasmtypes.ContractAddrLen)
)

func (chain *TestChain) StoreContractCode(suite *suite.Suite, path string) {
	simApp := chain.GetSimApp()

	wasmCode, err := os.ReadFile(path)
	suite.Require().NoError(err)

	govAuthority := simApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	src := wasmtypes.StoreCodeProposalFixture(func(p *wasmtypes.StoreCodeProposal) {
		p.RunAs = govAuthority.String()
		p.WASMByteCode = wasmCode
		checksum := sha256.Sum256(wasmCode)
		p.CodeHash = checksum[:]
	})

	msgServer := govkeeper.NewMsgServerImpl(simApp.GovKeeper)

	// ignore all submit events
	ctxIgnoreEvents := chain.GetContext().WithEventManager(sdk.NewEventManager())
	contentMsg, err := chain.submitLegacyProposal(
		suite,
		ctxIgnoreEvents,
		src,
		oneAddress,
		govAuthority.String(),
		msgServer,
	)
	suite.Require().NoError(err)

	content := v1.NewMsgExecLegacyContent(contentMsg.Content, govAuthority.String())
	_, err = msgServer.ExecLegacyContent(chain.GetContext(), content)
	suite.Require().NoError(err)
}

func (chain *TestChain) submitLegacyProposal(
	suite *suite.Suite,
	ctx sdk.Context,
	content govv1beta1.Content,
	myActorAddress sdk.AccAddress,
	govAuthority string,
	msgServer v1.MsgServer,
) (*v1.MsgExecLegacyContent, error) {
	contentMsg, err := v1.NewLegacyContent(content, govAuthority)
	suite.Require().NoError(err)

	proposal, err := v1.NewMsgSubmitProposal(
		[]sdk.Msg{contentMsg},
		sdk.Coins{},
		myActorAddress.String(),
		"",
		"my title",
		"my description",
	)
	suite.Require().NoError(err)

	_, err = msgServer.SubmitProposal(sdk.WrapSDKContext(ctx), proposal)
	return contentMsg, err
}

func (chain *TestChain) InstantiateRLContract(suite *suite.Suite, quotas string) sdk.AccAddress {
	simApp := chain.GetSimApp()
	transferModule := simApp.AccountKeeper.GetModuleAddress(transfertypes.ModuleName)
	govModule := simApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)

	initMsgBz := []byte(fmt.Sprintf(`{
           "gov_module":  "%s",
           "ibc_module":"%s",
           "paths": [%s]
        }`,
		govModule, transferModule, quotas))

	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(simApp.WasmKeeper)
	codeID := uint64(1)
	creator := simApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	addr, _, err := contractKeeper.Instantiate(chain.GetContext(), codeID, creator, creator, initMsgBz, "rate limiting contract", nil)
	suite.Require().NoError(err)
	return addr
}

func (chain *TestChain) StoreContractCodeDirect(suite *suite.Suite, path string) uint64 {
	simApp := chain.GetSimApp()
	govKeeper := wasmkeeper.NewGovPermissionKeeper(simApp.WasmKeeper)
	creator := simApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)

	wasmCode, err := os.ReadFile(path)
	suite.Require().NoError(err)
	accessEveryone := wasmtypes.AccessConfig{Permission: wasmtypes.AccessTypeEverybody}
	codeID, _, err := govKeeper.Create(chain.GetContext(), creator, wasmCode, &accessEveryone)
	suite.Require().NoError(err)
	return codeID
}

func (chain *TestChain) InstantiateContract(suite *suite.Suite, msg string, codeID uint64) sdk.AccAddress {
	simApp := chain.GetSimApp()
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(simApp.WasmKeeper)
	creator := simApp.AccountKeeper.GetModuleAddress(govtypes.ModuleName)
	addr, _, err := contractKeeper.Instantiate(chain.GetContext(), codeID, creator, creator, []byte(msg), "contract", nil)
	suite.Require().NoError(err)
	return addr
}

func (chain *TestChain) QueryContract(suite *suite.Suite, contract sdk.AccAddress, key []byte) string {
	simApp := chain.GetSimApp()
	state, err := simApp.WasmKeeper.QuerySmart(chain.GetContext(), contract, key)
	suite.Require().NoError(err)
	return string(state)
}

func (chain *TestChain) QueryContractJson(suite *suite.Suite, contract sdk.AccAddress, key []byte) gjson.Result {
	simApp := chain.GetSimApp()
	state, err := simApp.WasmKeeper.QuerySmart(chain.GetContext(), contract, key)
	suite.Require().NoError(err)
	suite.Require().True(gjson.Valid(string(state)))
	json := gjson.Parse(string(state))
	suite.Require().NoError(err)
	return json
}
