package error

type InvalidPassword struct {
	message string
}

func NewInvalidPassword(message string) InvalidPassword {
	return InvalidPassword{message: message}
}

func (err InvalidPassword) Error() string {
	return err.message
}
