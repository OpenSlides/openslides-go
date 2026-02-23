package auth

type authError struct {
	msg     string
	wrapped error
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
	return 403
}

// LogoutError indicates that a session was terminated due to logout.
//
// This error is used as the cause when cancelling a context due to backchannel
// logout.
type LogoutError struct{}

func (LogoutError) Type() string {
	return "auth-logout"
}

func (e LogoutError) Error() string {
	return "session logged out"
}
