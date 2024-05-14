package domainio

type ShortCodeDomain struct {
	Url string
}

type ErrorResponse struct {
	Error   error
	ErrMsg  string
	ErrCode int
}
