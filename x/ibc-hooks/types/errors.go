package types

import (
	errorsmod "cosmossdk.io/errors"
)

var (
	ErrBadMetadataFormatMsg = "wasm metadata not properly formatted for: '%v'. %s"
	ErrBadExecutionMsg      = "cannot execute contract: %v"

	ErrMsgValidation = errorsmod.Register(ModuleName, 2, "error in wasmhook message validation")
	ErrMarshaling    = errorsmod.Register(ModuleName, 3, "cannot marshal the ICS20 packet")
	ErrInvalidPacket = errorsmod.Register(ModuleName, 4, "invalid packet data")
	ErrBadResponse   = errorsmod.Register(ModuleName, 5, "cannot create response")
	ErrWasmError     = errorsmod.Register(ModuleName, 6, "wasm error")
	ErrBadSender     = errorsmod.Register(ModuleName, 7, "bad sender")
)
