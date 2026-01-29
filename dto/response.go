package dto

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func CreateResponseError(message string) Response[string] {
	return Response[string]{
		Success: false,
		Message: message,
		Data:    "",
	}
}
func CreateResponseSuccess[T any](data T) Response[T] {
	return Response[T]{
		Success: true,
		Message: "success",
		Data:    data,
	}
}
func CreateResponseErrorData(message string, data map[string]string) Response[map[string]string] {
	return Response[map[string]string]{
		Success: false,
		Message: message,
		Data:    data,
	}
}
