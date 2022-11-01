package claim

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"mun/testutil/sample"
	claimsimulation "mun/x/claim/simulation"
	
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = claimsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgClaim = "op_weight_msg_claim"
	// TODO: Determine the simulation weight value
	defaultWeightMsgClaim int = 100

	// this line is used by starport scaffolding # simapp/module/const
)




// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgClaim int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaim, &weightMsgClaim, nil,
		func(_ *rand.Rand) {
			weightMsgClaim = defaultWeightMsgClaim
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaim,
		claimsimulation.SimulateMsgClaim(am.accountKeeper, am.bankKeeper, am.distrKeeper, am.stakingKeeper, am.govKeeper, am.keeper),
		
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
