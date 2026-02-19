package auth

type authError struct {
	msg     string
	wrapped error
	status  int
}

func (authError) Type() string {
	return "auth"
}

func (a authError) Error() string {
	return a.msg
}

func (a authError) Unwrap() error {
	return a.wrapped
}

func (a authError) StatusCode() int {
	if a.status != 0 {
		return a.status
	}
	return 403
}
