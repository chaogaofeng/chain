package identity

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/goldnet/chain/testutil/sample"
	identitysimulation "github.com/goldnet/chain/x/identity/simulation"
	"github.com/goldnet/chain/x/identity/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = identitysimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreate int = 100

	opWeightMsgUpdate = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdate int = 100

	opWeightMsgRemove = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRemove int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	identityGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&identityGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	identityParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyEnabled), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(identityParams.Enabled))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreate, &weightMsgCreate, nil,
		func(_ *rand.Rand) {
			weightMsgCreate = defaultWeightMsgCreate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreate,
		identitysimulation.SimulateMsgCreate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdate int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdate, &weightMsgUpdate, nil,
		func(_ *rand.Rand) {
			weightMsgUpdate = defaultWeightMsgUpdate
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdate,
		identitysimulation.SimulateMsgUpdate(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRemove int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRemove, &weightMsgRemove, nil,
		func(_ *rand.Rand) {
			weightMsgRemove = defaultWeightMsgRemove
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRemove,
		identitysimulation.SimulateMsgRemove(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
