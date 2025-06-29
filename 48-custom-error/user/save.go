package user

type ValidationError struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func Save(id string, data interface{}) error {
	if id == "" {
		return &ValidationError{Message: "ID is empty."}
	}

	if id != "taka" {
		return &NotFoundError{Message: "Data not found."}
	}

	return nil
}
