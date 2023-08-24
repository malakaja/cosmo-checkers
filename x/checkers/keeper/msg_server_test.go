package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/malakaja/cosmo-checkers/checkers/testutil/keeper"
	"github.com/malakaja/cosmo-checkers/checkers/x/checkers"
	"github.com/malakaja/cosmo-checkers/checkers/x/checkers/keeper"
	"github.com/malakaja/cosmo-checkers/checkers/x/checkers/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, *keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), k, sdk.WrapSDKContext(ctx)
}
