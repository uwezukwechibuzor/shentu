package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/shentufoundation/shentu/v2/x/bounty/types"
)

// Keeper - bounty keeper
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryCodec
	paramSpace paramtypes.Subspace

	//authKeeper authtypes.AccountKeeper
	certKeeper types.CertKeeper
}

// NewKeeper creates a new Keeper object
func NewKeeper(
	cdc codec.BinaryCodec, storeKey sdk.StoreKey, ck types.CertKeeper, paramSpace paramtypes.Subspace,
) Keeper {

	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
		certKeeper: ck,
	}
}
