package errs

import (
	"errors"
	"fmt"
)

var (
	UnmarshalerError = errors.New("unmarshaler error")
	DataStoreError   = errors.New("data store error")
)

func WrapError(code any, err error) error {
	return fmt.Errorf("[code: %s] %s", code, err.Error())
}
