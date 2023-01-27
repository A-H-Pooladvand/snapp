package res

import (
	"encoding/json"
	"net/http"
)

type Output struct {
	Writer http.ResponseWriter
}

func (o *Output) Ok(obj any) {
	o.Response(http.StatusOK, obj)
}

func (o *Output) Response(code int, obj any) {
	o.Writer.Header().Set("Content-Type", "application/json")
	o.Writer.WriteHeader(code)

	json.NewEncoder(o.Writer).Encode(obj)
}

func (o *Output) ServerError(obj any) {
	o.Error(http.StatusInternalServerError, obj)
}

func (o *Output) Error(status int, obj any) {
	o.Response(http.StatusInternalServerError, map[string]any{
		"error": obj,
	})
}

func (o *Output) BadRequest(obj any) {
	o.Error(http.StatusBadRequest, obj)
}

func Response(w http.ResponseWriter) *Output {
	return &Output{
		Writer: w,
	}
}
