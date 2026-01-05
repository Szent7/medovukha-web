package types

func NewSuccess[T any](data T) Response[T] {
	return Response[T]{
		Success: true,
		Error:   nil,
		Data:    &data,
	}
}

func NewFailed[T any](error APIError) Response[T] {
	return Response[T]{
		Success: false,
		Error:   &error,
		Data:    nil,
	}
}
