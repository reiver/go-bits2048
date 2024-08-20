package bits2048

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilReceiver    = erorr.Error("bits2048: nil receiver")
	errStringTooLong  = erorr.Error("bits2048: string too long")
)
