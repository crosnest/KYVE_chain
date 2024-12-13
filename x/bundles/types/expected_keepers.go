package types

import (
	"cosmossdk.io/math"
	"github.com/KYVENetwork/chain/x/funders/types"
	pooltypes "github.com/KYVENetwork/chain/x/pool/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetModuleAddress(moduleName string) sdk.AccAddress
}

type PoolKeeper interface {
	AssertPoolExists(ctx sdk.Context, poolId uint64) error
	GetPoolWithError(ctx sdk.Context, poolId uint64) (pooltypes.Pool, error)
	GetPool(ctx sdk.Context, id uint64) (val pooltypes.Pool, found bool)

	IncrementBundleInformation(ctx sdk.Context, poolId uint64, currentHeight uint64, currentKey string, currentValue string)

	GetAllPools(ctx sdk.Context) (list []pooltypes.Pool)
	ChargeInflationPool(ctx sdk.Context, poolId uint64) (payout uint64, err error)
	GetProtocolInflationShare(ctx sdk.Context) (res math.LegacyDec)
	GetMaxVotingPowerPerPool(ctx sdk.Context) (res math.LegacyDec)
}

type StakerKeeper interface {
	GetAllStakerAddressesOfPool(ctx sdk.Context, poolId uint64) (stakers []string)
	AssertValaccountAuthorized(ctx sdk.Context, poolId uint64, stakerAddress string, valaddress string) error

	DoesValaccountExist(ctx sdk.Context, poolId uint64, stakerAddress string) bool

	LeavePool(ctx sdk.Context, staker string, poolId uint64)

	IncrementPoints(ctx sdk.Context, poolId uint64, stakerAddress string) (newPoints uint64)
	ResetPoints(ctx sdk.Context, poolId uint64, stakerAddress string) (previousPoints uint64)

	GetTotalAndHighestDelegationOfPool(ctx sdk.Context, poolId uint64) (totalDelegation, highestDelegation uint64)
	GetDelegationAmount(ctx sdk.Context, validator string) uint64
	GetValidator(ctx sdk.Context, staker string) (stakingtypes.Validator, bool)
	Slash(ctx sdk.Context, poolId uint64, staker string, slashFraction math.LegacyDec)
	PayoutRewards(ctx sdk.Context, staker string, amount sdk.Coins, payerModuleName string) error
	PayoutAdditionalCommissionRewards(ctx sdk.Context, validator string, payerModuleName string, amount sdk.Coins) error
}

type DelegationKeeper interface {
	GetTimeoutSlash(ctx sdk.Context) (res math.LegacyDec)
	GetUploadSlash(ctx sdk.Context) (res math.LegacyDec)
	GetVoteSlash(ctx sdk.Context) (res math.LegacyDec)
}

type FundersKeeper interface {
	GetCoinWhitelistMap(ctx sdk.Context) (whitelist map[string]types.WhitelistCoinEntry)
	ChargeFundersOfPool(ctx sdk.Context, poolId uint64, recipient string) (payout sdk.Coins, err error)
}

type TeamKeeper interface {
	GetTeamBlockProvision(ctx sdk.Context) int64
}
