package domainio

type ShortCodeDomain struct {
	Url string
}

type ErrorResponse struct {
	Error   error
	ErrMsg  string
	ErrCode int
}

type RedirectUrl struct {
	Code      int
	ShortCode string
	Url       string
}
