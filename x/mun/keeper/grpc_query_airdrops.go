package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mun/x/mun/types"
)

func (k Keeper) Airdrops(goCtx context.Context, req *types.QueryAirdropsRequest) (*types.QueryAirdropsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryAirdropsResponse{}, nil
}
