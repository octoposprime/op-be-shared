package model

import "errors"

var ERRORS []error = []error{
	ErrorNone,
	ErrorRedisConnection,
	ErrorDbConnection,
	ErrorUserNotFound,
	ErrorTokenExpired,
	ErrorDecodeJwtFailed,
	ErrorUserUnAuthorized,
	ErrorChannelNameNotValid,
}

const ErrSep string = "."
const (
	ErrBase        string = "error"
	ErrUser        string = "user"
	ErrJwt         string = "jwt"
	ErrDb          string = "db"
	ErrRedis       string = "redis"
	ErrConnection  string = "connection"
	ErrToken       string = "token"
	ErrDecode      string = "decode"
	ErrChannelName string = "channelname"
)

const (
	ErrFailed         string = "failed"
	ErrNotFound       string = "notfound"
	ErrNotValid       string = "notvalid"
	ErrExpired        string = "expired"
	ErrorUnAuthorized string = "unauthorized"
)

var (
	ErrorNone                error = nil
	ErrorRedisConnection           = errors.New(ErrBase + ErrSep + ErrRedis + ErrSep + ErrConnection + ErrSep + ErrFailed)
	ErrorDbConnection              = errors.New(ErrBase + ErrSep + ErrDb + ErrSep + ErrConnection + ErrSep + ErrFailed)
	ErrorUserNotFound              = errors.New(ErrBase + ErrSep + ErrUser + ErrSep + ErrNotFound)
	ErrorTokenExpired              = errors.New(ErrBase + ErrSep + ErrToken + ErrSep + ErrExpired)
	ErrorDecodeJwtFailed           = errors.New(ErrBase + ErrSep + ErrJwt + ErrSep + ErrDecode + ErrSep + ErrFailed)
	ErrorUserUnAuthorized          = errors.New(ErrBase + ErrSep + ErrUser + ErrSep + ErrorUnAuthorized)
	ErrorChannelNameNotValid       = errors.New(ErrBase + ErrSep + ErrChannelName + ErrSep + ErrNotValid)
)

func GetErrors() []error {
	return ERRORS
}
