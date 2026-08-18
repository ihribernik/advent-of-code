package inputs

import "errors"

// ErrEmptyInput indicates that an input file contains no lines.
var ErrEmptyInput = errors.New("input file is empty")
