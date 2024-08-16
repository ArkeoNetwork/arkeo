package keeper

import (
	"testing"

	storemetrics "cosmossdk.io/store/metrics"
	arekoappParams "github.com/arkeonetwork/arkeo/app/params"
	"github.com/arkeonetwork/arkeo/common"
	"github.com/arkeonetwork/arkeo/common/cosmos"
	"github.com/arkeonetwork/arkeo/testutil/utils"
	"github.com/arkeonetwork/arkeo/x/arkeo/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	tmdb "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/codec"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	bech32Prefix = "arkeo"
)

var (
	// Bech32PrefixAccAddr
	Bech32PrefixAccAddr  = bech32Prefix
	Bech32PrefixAccPub   = bech32Prefix + sdk.PrefixPublic
	Bech32PrefixValAddr  = bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	Bech32PrefixValPub   = bech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	Bech32PrefixConsAddr = bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	Bech32PrefixConsPub  = bech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

func SetupKeeper(t testing.TB) (cosmos.Context, Keeper) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	keyAcc := cosmos.NewKVStoreKey(authtypes.StoreKey)
	keyBank := cosmos.NewKVStoreKey(banktypes.StoreKey)
	keyStake := cosmos.NewKVStoreKey(stakingtypes.StoreKey)
	keyParams := cosmos.NewKVStoreKey(typesparams.StoreKey)
	tkeyParams := cosmos.NewTransientStoreKey(typesparams.TStoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	cfg := sdk.GetConfig()

	cfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)

	logger := log.NewNopLogger()
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, logger, storemetrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyAcc, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyBank, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyParams, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeyParams, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	encodingConfig := arekoappParams.MakeEncodingConfig()
	types.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	cdc := utils.MakeTestMarshaler()
	amino := encodingConfig.Amino

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"ArkeoParams",
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	_ = paramskeeper.NewKeeper(cdc, amino, keyParams, tkeyParams)
	govModuleAddr := authtypes.NewModuleAddress(govtypes.ModuleName).String()
	ak := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(keyAcc),
		authtypes.ProtoBaseAccount,
		map[string][]string{
			stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
			stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
			types.ModuleName:               {authtypes.Minter, authtypes.Burner},
			types.ReserveName:              {},
			types.ProviderName:             {},
			types.ContractName:             {},
		},
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.Bech32PrefixAccAddr,
		govModuleAddr,
	)

	bk := bankkeeper.NewBaseKeeper(
		cdc,
		runtime.NewKVStoreService(keyBank),
		ak,
		nil,
		govModuleAddr,
		logger,
	)
	require.NoError(t, bk.SetParams(ctx, banktypes.DefaultParams()))

	sk := stakingkeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(keyStake),
		ak,
		bk,
		govModuleAddr,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)

	k := NewKVStore(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		bk,
		ak,
		*sk,
		govModuleAddr,
		logger,
	)
	k.SetVersion(ctx, common.GetCurrentVersion())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return ctx, *k
}

func SetupKeeperWithStaking(t testing.TB) (cosmos.Context, Keeper, stakingkeeper.Keeper) {
	storeKey := storetypes.NewKVStoreKey(types.StoreKey)
	keyAcc := cosmos.NewKVStoreKey(authtypes.StoreKey)
	keyBank := cosmos.NewKVStoreKey(banktypes.StoreKey)
	keyStake := cosmos.NewKVStoreKey(stakingtypes.StoreKey)
	keyParams := cosmos.NewKVStoreKey(typesparams.StoreKey)
	tkeyParams := cosmos.NewTransientStoreKey(typesparams.TStoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	cfg := sdk.GetConfig()

	cfg.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	cfg.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	cfg.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)

	logger := log.NewNopLogger()
	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db, logger, storemetrics.NewNoOpMetrics())
	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyAcc, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyBank, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyStake, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(keyParams, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(tkeyParams, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	encodingConfig := arekoappParams.MakeEncodingConfig()
	types.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	cdc := utils.MakeTestMarshaler()
	amino := encodingConfig.Amino

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"ArkeoParams",
	)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, log.NewNopLogger())

	govModuleAddr := "arkeo12jurap3ypfmwzghj9d0t4wanc9gne2cktxh2y2"
	_ = paramskeeper.NewKeeper(cdc, amino, keyParams, tkeyParams)
	ak := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(keyAcc),
		authtypes.ProtoBaseAccount,
		map[string][]string{
			stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
			stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
			types.ModuleName:               {authtypes.Minter, authtypes.Burner},
			types.ReserveName:              {},
			types.ProviderName:             {},
			types.ContractName:             {},
			govtypes.ModuleName:            {authtypes.Minter, authtypes.Burner},
		},
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32AccountAddrPrefix()),
		sdk.Bech32PrefixAccAddr,
		govModuleAddr,
	)

	bk := bankkeeper.NewBaseKeeper(
		cdc,
		runtime.NewKVStoreService(keyBank),
		ak,
		nil,
		govModuleAddr,
		logger,
	)
	_ = bk.SetParams(ctx, banktypes.DefaultParams())

	sk := stakingkeeper.NewKeeper(
		cdc,
		runtime.NewKVStoreService(keyStake),
		ak,
		bk,
		govModuleAddr,
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ValidatorAddrPrefix()),
		authcodec.NewBech32Codec(sdk.GetConfig().GetBech32ConsensusAddrPrefix()),
	)

	k := NewKVStore(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
		bk,
		ak,
		*sk,
		govModuleAddr,
		logger,
	)
	k.SetVersion(ctx, common.GetCurrentVersion())

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return ctx, *k, *sk
}
