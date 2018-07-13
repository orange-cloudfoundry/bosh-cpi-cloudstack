package action

type NotImplementedError string

func (e NotImplementedError) Error() string {
	return string(e)
}

func (e NotImplementedError) Type() string {
	return "Bosh::Clouds::NotImplemented"
}

func NewNotImplementedError(err error) NotImplementedError {
	return NotImplementedError(err.Error())
}
