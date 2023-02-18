package statuses

import "errors"

var (
	UnmarshallerError = errors.New("unmarshaller error")
	DataStoreError    = errors.New("data store error")
)
