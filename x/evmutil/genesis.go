package evmutil

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/0glabs/0g-chain/x/evmutil/keeper"
	"github.com/0glabs/0g-chain/x/evmutil/types"
)

// InitGenesis initializes the store state from a genesis state.
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, gs *types.GenesisState, ak types.AccountKeeper) {
	if err := gs.Validate(); err != nil {
		panic(fmt.Sprintf("failed to validate %s genesis state: %s", types.ModuleName, err))
	}

	keeper.SetParams(ctx, gs.Params)

	// initialize module account
	if moduleAcc := ak.GetModuleAccount(ctx, types.ModuleName); moduleAcc == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	for _, account := range gs.Accounts {
		keeper.SetAccount(ctx, account)
	}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper keeper.Keeper) *types.GenesisState {
	accounts := keeper.GetAllAccounts(ctx)
	return types.NewGenesisState(accounts, keeper.GetParams(ctx))
}
