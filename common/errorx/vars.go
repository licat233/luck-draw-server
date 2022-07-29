package errorx

import "errors"

var (
	ErrInvalidEmail = ExternalError(errors.New("invalid email"))
	ErrEmptyParam   = ExternalError(errors.New("Parameter cannot be empty"))
	ErrNotFound     = ExternalError(errors.New("not found"))
	ErrInvalidParam = ExternalError(errors.New("invalid parameter"))
	ErrInvalidUuid  = ExternalError(errors.New("uuid format error"))

	ErrAuth     = AuthError(errors.New("Authentication failed, identity exception"))
	ErrPlatform = AuthError(errors.New("Authentication failed, platform exception"))
	ErrIp       = AuthError(errors.New("Authentication failed, request ip exception"))
)
