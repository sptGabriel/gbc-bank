package transfers

import "errors"

var ErrSELFTransfer = errors.New("origin account cannot be the same as the destination account")
