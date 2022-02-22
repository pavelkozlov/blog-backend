//go:generate easyjson -all ./utils.go
package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-playground/validator/v10"
)

func init() {
	v = validator.New()
}

var (
	v *validator.Validate
)

type ErrResponse struct {
	Error string `json:"error"`
}

type OkResponse struct {
	Data interface{} `json:"data"`
}

type UnmarshalJSON func(data []byte) error

func ParseAndValidateBody(input interface{}, unmarshal UnmarshalJSON, r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("can not ioutil.ReadAll: %w", err)
	}
	if err = unmarshal(body); err != nil {
		return fmt.Errorf("can not UnmarshalJSON: %w", err)
	}
	if err = v.Struct(input); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}
	return nil
}

func WriteErrResponse(err error, w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json")
	e, _ := ErrResponse{err.Error()}.MarshalJSON()
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	//nolint:errcheck
	fmt.Fprintln(w, string(e))
}

func WriteOkResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	d, _ := OkResponse{data}.MarshalJSON()
	//nolint:errcheck
	w.Write(d)
}
func ParseLimitAndOffset(r *http.Request) (int, int, error) {
	limit, offset := 10, 0
	u, err := url.Parse(r.URL.String())
	if err != nil {
		return limit, offset, fmt.Errorf("can not url.Parse: %w", err)
	}
	q, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return limit, offset, fmt.Errorf("can not url.ParseQuery: %w", err)
	}
	if q.Get("limit") != "" {
		incomeLimit, err := strconv.Atoi(q.Get("limit"))
		if err != nil {
			return limit, offset, fmt.Errorf("invalid value for limit: %w", err)
		}
		limit = incomeLimit
	}
	if q.Get("offset") != "" {
		incomeOffset, err := strconv.Atoi(q.Get("offset"))
		if err != nil {
			return limit, offset, fmt.Errorf("invalid value for offset: %w", err)
		}
		offset = incomeOffset
	}
	return limit, offset, nil
}
