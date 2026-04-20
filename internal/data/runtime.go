package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

type Runtime int

func (r Runtime) MarshalJSON() ([]byte, error) {
	value := fmt.Sprintf("%d mins", r)
	quoted := strconv.Quote(value)

	return []byte(quoted), nil
}

func (r *Runtime) UnmarshalJSON(val []byte) error {
	unquoted, err := strconv.Unquote(string(val))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquoted, " ")
	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.Atoi(parts[0])
	if err != nil {
		return ErrInvalidRuntimeFormat
	}

	*r = Runtime(i)

	return nil
}
