package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidCollection = sdkerrors.Register(ModuleName, 1, "invalid NFT collection")
	ErrUnknownCollection = sdkerrors.Register(ModuleName, 2, "unknown NFT collection")
	ErrInvalidNFT        = sdkerrors.Register(ModuleName, 3, "invalid NFT")
	ErrUnknownNFT        = sdkerrors.Register(ModuleName, 4, "unknown NFT")
	ErrNFTAlreadyExists  = sdkerrors.Register(ModuleName, 5, "NFT already exists")
	ErrEmptyProof        = sdkerrors.Register(ModuleName, 6, "NFT proof can't be empty")
	ErrEmptyMetadata     = sdkerrors.Register(ModuleName, 7, "Empty metadata")
)
