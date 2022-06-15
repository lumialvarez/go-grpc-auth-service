package error

type AlreadyExists struct {
	message string
}

func NewAlreadyExists(message string) AlreadyExists {
	return AlreadyExists{message: message}
}

func (err AlreadyExists) Error() string {
	return err.message
}
