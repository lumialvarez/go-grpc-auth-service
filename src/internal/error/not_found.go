package error

type NotFound struct {
	message string
}

func NewNotFound(message string) NotFound {
	return NotFound{message: message}
}

func (err NotFound) Error() string {
	return err.message
}
