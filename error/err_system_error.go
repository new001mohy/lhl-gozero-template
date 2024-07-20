package error

type ErrSystemError struct{}

func (e *ErrSystemError) Error() string {
	return ""
}
