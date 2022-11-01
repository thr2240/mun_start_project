package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"mun/x/claim/types"
)

func (k msgServer) Claim(goCtx context.Context, msg *types.MsgClaim) (*types.MsgClaimResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := k.Keeper.Claim(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgClaimResponse{}, nil
}
