package error

type Inactive struct {
	message string
}

func NewInactive(message string) Inactive {
	return Inactive{message: message}
}

func (err Inactive) Error() string {
	return err.message
}
