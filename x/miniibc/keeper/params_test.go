package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mini_ibc/testutil/keeper"
	"mini_ibc/x/miniibc/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MiniibcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
