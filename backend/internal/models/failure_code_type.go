package models

//go:generate go run github.com/dmarkham/enumer -type=FailureCodeType -json -text -sql -transform=snake -trimprefix=FailureCode -linecomment
//swagger:enum FailureCode
type FailureCodeType int

const (
	FailureCodeNil FailureCodeType = iota // FailureCode
	FailureCodeParseRequest
	FailureCodeNotFound
	FailureCodeCreateResponse
	FailureCodeServiceFailed
	FailureCodeCreateUser
	FailureCodePaymentMethodNotFound
	FailureCodeInvalidPartnerConfig
)
