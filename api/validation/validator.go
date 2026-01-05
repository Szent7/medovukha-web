package validation

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Szent7/medovukha-web/api/rest/v1/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleBindError(c *gin.Context, err error) {
	// Validation errors 422
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		fields := make(map[string]string)

		for _, fe := range ve {
			fields[fe.Field()] = fe.Tag()
		}

		c.IndentedJSON(http.StatusUnprocessableEntity,
			types.NewFailed[types.BaseMessage](types.APIError{
				Code:   types.ErrValidation,
				Fields: fields,
			}))
		return
	}

	// Parsing errors 400
	var syntaxErr *json.SyntaxError
	var typeErr *json.UnmarshalTypeError

	switch {
	case errors.As(err, &syntaxErr):
		c.IndentedJSON(http.StatusBadRequest,
			types.NewFailed[types.BaseMessage](types.APIError{
				Code:    types.ErrParse,
				Message: "invalid json format",
			}))
	case errors.As(err, &typeErr):
		c.IndentedJSON(http.StatusBadRequest,
			types.NewFailed[types.BaseMessage](types.APIError{
				Code:    types.ErrParse,
				Message: "invalid field type",
			}))
	default:
		// Common 400
		c.IndentedJSON(http.StatusBadRequest,
			types.NewFailed[types.BaseMessage](types.APIError{
				Code:    types.ErrParse,
				Message: "unknown error",
			}))
	}
}
