package mailerdb

type ErrUserNotFound struct {
	message string
}

type ErrUserEmailConflict struct {
	message string
}

type ErrClientNotFound struct {
	message string
}

type ErrCampaignNotFound struct {
	message string
}

type ErrContactFormNotFound struct {
	message string
}

// NewErrUserNotFound to create a new error struct when the db isn't able
// to find a user based on some criteria
func newErrUserNotFound(message string) *ErrUserNotFound {
	return &ErrUserNotFound{
		message: message,
	}
}

func (e ErrUserNotFound) Error() string {
	return e.message
}

func IsErrUserNotFound(err error) bool {
	_, ok := err.(*ErrUserNotFound)
	return ok
}
