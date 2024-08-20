package bits2048

import (
	"strconv"
	"unsafe"
)

// String is string that can be 0 to 255 bytes long.
//
// String always takes up exactly 256-bytes of space
// (regarless of how long the string is).
//
// Internally String is a length-prefixed string.
// This is also sometimes called a “Pascal String”.
//
// Example usage:
//
//	var str bits2048.String
//	
//	str.FromString("Hello world! 🙂")
type String struct {
	value [2048/8]byte
}

// FromString returns a bits2048.String from a Go string.
//
// If the Go string is longer than 256-bytes then it returns an error.
func FromString(value string) (String, error) {
	var result String
	err := result.FromString(value)
	return result, err
}

// MustFromString returns a bits2048.String from a Go string.
//
// If the Go string is longer than 256-bytes then it panics.
func MustFromString(value string) String {
	result, err := FromString(value)
	if nil != err {
		panic(err)
	}
	return result
}

// Bytes returns bits2048.String as a Go []byte.
//
// (Note that this does NOT include the length-prefix, onl the string data.)
func (receiver *String) Bytes() []byte {
	if nil == receiver {
		panic(errNilReceiver)
	}

	return append([]byte(nil), receiver.bytes()...)
}

// String returns bits2048.String as a Go string.
func (receiver *String) bytes() []byte {
	if nil == receiver {
		return nil
	}

	var length int = receiver.Len()
	var p []byte = receiver.value[1:]

	p = (p[:length])

	return p
}

func (receiver String) GoString() string {
	var p []byte

	p = append(p, "bits2048.MustFromString("...)
	p = strconv.AppendQuote(p, receiver.String())
	p = append(p, ')')

	return string(p)
}

// Len returns the length of bits2048.String as the number of bytes.
func (receiver *String) Len() int {
	if nil == receiver {
		return 0
	}

	return int(receiver.value[0])
}

func (receiver *String) MarshalText() (text []byte, err error) {
	if nil == receiver {
		return nil, errNilReceiver
	}

	return receiver.Bytes(), nil
}

// FromString returns a bits2048.String from a Go string.
//
// If the Go string is longer than 256-bytes then it returns an error.
func (receiver *String) FromString(value string) error {
	if nil == receiver {
		return errNilReceiver
	}

	var length int = len(value)

	if len(receiver.value) < length {
		return errStringTooLong
	}

	receiver.value[0] = byte(length)

	for i:=0; i<length; i++ {
		receiver.value[i+1] = value[i]
	}

	return nil
}

// String returns bits2048.String as a Go string.
func (receiver *String) String() string {
	if nil == receiver {
		return ""
	}

	return string(receiver.bytes())
}

func (receiver *String) UnmarshalText(text []byte) error {
	if nil == receiver {
		return errNilReceiver
	}

	var length int = len(text)
	var ptr *byte = unsafe.SliceData(text)

	var str string = unsafe.String(ptr, length)

	return receiver.FromString(str)
}
