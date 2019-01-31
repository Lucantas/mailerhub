package jwt

type ErrInvalidToken struct {
	message string
}

type ErrEmptyToken struct {
	message string
}

// NewErrInvalidToken to create a new error struct of type Invalid Token
func NewErrInvalidToken(message string) *ErrInvalidToken {
	return &ErrInvalidToken{
		message: message,
	}
}

func (e *ErrInvalidToken) Error() string {
	return e.message
}

// NewErrEmptyToken to create a new error struct of type Empty Token
func NewErrEmptyToken(message string) *ErrEmptyToken {
	return &ErrEmptyToken{
		message: message,
	}
}

func (e *ErrEmptyToken) Error() string {
	return e.message
}
