// Copyright (c) 2016-2021 Shanghai Bianjie AI Technology Inc. (licensed under the Apache License, Version 2.0)
// Modifications Copyright (c) 2021-present Cronos.org (licensed under the Apache License, Version 2.0)
package types

// DONTCOVER

import (
	gogotypes "github.com/cosmos/gogoproto/types"
	"github.com/crypto-org-chain/chain-main/v4/x/nft/exported"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

var ModuleCdc = codec.NewLegacyAmino()

func init() {
	RegisterLegacyAminoCodec(ModuleCdc)
	cryptocodec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}

// RegisterLegacyAminoCodec concrete types on codec
// (Amino is still needed for Ledger at the moment)
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgIssueDenom{}, "chainmain/nft/MsgIssueDenom", nil)
	cdc.RegisterConcrete(&MsgTransferNFT{}, "chainmain/nft/MsgTransferNFT", nil)
	cdc.RegisterConcrete(&MsgEditNFT{}, "chainmain/nft/MsgEditNFT", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "chainmain/nft/MsgMintNFT", nil)
	cdc.RegisterConcrete(&MsgBurnNFT{}, "chainmain/nft/MsgBurnNFT", nil)

	cdc.RegisterInterface((*exported.NFT)(nil), nil)
	cdc.RegisterConcrete(&BaseNFT{}, "chainmain/nft/BaseNFT", nil)
}

func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgIssueDenom{},
		&MsgTransferNFT{},
		&MsgEditNFT{},
		&MsgMintNFT{},
		&MsgBurnNFT{},
	)

	registry.RegisterImplementations((*exported.NFT)(nil),
		&BaseNFT{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

// MustMarshalSupply return supply protobuf code
func MustMarshalSupply(cdc codec.Codec, supply uint64) []byte {
	supplyWrap := gogotypes.UInt64Value{Value: supply}
	return cdc.MustMarshal(&supplyWrap)
}

// MustUnMarshalSupply return the supply
func MustUnMarshalSupply(cdc codec.Codec, value []byte) uint64 {
	var supplyWrap gogotypes.UInt64Value
	cdc.MustUnmarshal(value, &supplyWrap)
	return supplyWrap.Value
}

// MustMarshalTokenID return the tokenID protobuf code
func MustMarshalTokenID(cdc codec.Codec, tokenID string) []byte {
	tokenIDWrap := gogotypes.StringValue{Value: tokenID}
	return cdc.MustMarshal(&tokenIDWrap)
}

// MustUnMarshalTokenID return the tokenID
func MustUnMarshalTokenID(cdc codec.Codec, value []byte) string {
	var tokenIDWrap gogotypes.StringValue
	cdc.MustUnmarshal(value, &tokenIDWrap)
	return tokenIDWrap.Value
}
