package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgClaimEth{}, "claim/ClaimEth", nil)
	cdc.RegisterConcrete(&MsgClaimArkeo{}, "claim/ClaimArkeo", nil)
	cdc.RegisterConcrete(&MsgTransferClaim{}, "claim/TransferClaim", nil)
	cdc.RegisterConcrete(&MsgAddClaim{}, "claim/AddClaim", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgTransferClaim{},
		&MsgAddClaim{},
		&MsgClaimEth{},
		&MsgClaimArkeo{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
