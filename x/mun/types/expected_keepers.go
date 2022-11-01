package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
 )

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	// MintCoins is used only for simulation test codes
	MintCoins(ctx sdk.Context, name string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

// DistrKeeper is the keeper of the distribution store
type DistrKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type GovKeeper interface {
	IterateProposals(ctx sdk.Context, cb func(proposal govtypes.Proposal) (stop bool))
	GetVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) (vote govtypes.Vote, found bool)
	// Note that this function is only used before the UpgradeHeight defined in app/upgrades/v1.1.0
	IterateAllVotes(ctx sdk.Context, cb func(vote govtypes.Vote) (stop bool))
}

type StakingKeeper interface {
	GetParams(ctx sdk.Context) (params stakingtypes.Params)
}
