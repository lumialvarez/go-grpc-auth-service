package error

type InvalidCredentials struct {
	message string
}

func NewInvalidCredentials(message string) InvalidCredentials {
	return InvalidCredentials{message: message}
}

func (err InvalidCredentials) Error() string {
	return err.message
}
