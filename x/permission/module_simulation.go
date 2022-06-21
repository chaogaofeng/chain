package permission

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/goldnet/chain/testutil/sample"
	permissionsimulation "github.com/goldnet/chain/x/permission/simulation"
	"github.com/goldnet/chain/x/permission/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = permissionsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgAssign = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAssign int = 100

	opWeightMsgUnassign = "op_weight_msg_create_chain"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUnassign int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	permissionGenesis := types.GenesisState{
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&permissionGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	permissionParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyEnabled), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(permissionParams.Enabled))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgAssign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAssign, &weightMsgAssign, nil,
		func(_ *rand.Rand) {
			weightMsgAssign = defaultWeightMsgAssign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAssign,
		permissionsimulation.SimulateMsgAssign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUnassign int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUnassign, &weightMsgUnassign, nil,
		func(_ *rand.Rand) {
			weightMsgUnassign = defaultWeightMsgUnassign
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUnassign,
		permissionsimulation.SimulateMsgUnassign(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
