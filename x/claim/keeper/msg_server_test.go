package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arkeonetwork/arkeo/testutil/keeper"
	"github.com/arkeonetwork/arkeo/x/claim/keeper"
	"github.com/arkeonetwork/arkeo/x/claim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, keepertest.TestKeepers, context.Context) {
	keepers, ctx := keepertest.CreateTestClaimKeepers(t)
	return keeper.NewMsgServerImpl(keepers.ClaimKeeper), keepers, sdk.WrapSDKContext(ctx) //nolint:staticcheck
}
