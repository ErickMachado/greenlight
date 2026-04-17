package data

import (
	"fmt"
	"strconv"
)

type Runtime int

func (r Runtime) MarshalJSON() ([]byte, error) {
	value := fmt.Sprintf("%d mins", r)
	quoted := strconv.Quote(value)

	return []byte(quoted), nil
}
