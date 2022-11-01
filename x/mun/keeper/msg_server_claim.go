package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mun/x/mun/types"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgClaimResponse{}, nil
}
