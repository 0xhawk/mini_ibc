package miniibc_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mini_ibc/testutil/keeper"
	"mini_ibc/testutil/nullify"
	"mini_ibc/x/miniibc"
	"mini_ibc/x/miniibc/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MiniibcKeeper(t)
	miniibc.InitGenesis(ctx, *k, genesisState)
	got := miniibc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
