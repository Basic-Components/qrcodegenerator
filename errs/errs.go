package errs

import (
	"errors"
)

var ConfigDecodeError error = errors.New("unable to decode config into struct")
var ConfigVerifyError error = errors.New("config not satisfied the schema")

var DataDifferentError error = errors.New("data different error")
