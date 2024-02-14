package model

type HeaderKey string

const (
	HeaderKeyContentType   HeaderKey = "Content-Type"
	HeaderKeyAppJson       HeaderKey = "application/json"
	HeaderKeyAuthorization HeaderKey = "Authorization"
)
