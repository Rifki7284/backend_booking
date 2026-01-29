package util

import (
	"fmt"

	"github.com/go-playground/validator"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	res := map[string]string{}
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			res[v.StructField()] = TranslateTag(v)
		}
	}

	return res
}
func TranslateTag(fd validator.FieldError) string {
	switch fd.Tag() {
	case "required":
		return fmt.Sprintf("field %s wajib diisi", fd.StructField())
	}
	return "validasi gagal"
}
