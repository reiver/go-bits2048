package bits2048

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrOverFlow = erorr.Error("overflow")
)

const (
	errNilReceiver    = erorr.Error("bits2048: nil receiver")
	errNilSourceOne   = erorr.Error("bits2048: nil source-one")
	errNilSourceTwo   = erorr.Error("bits2048: nil source-two")
	errStringTooLong  = erorr.Error("bits2048: string too long")
	errTooMany        = erorr.Error("bits2048: too many")
)
