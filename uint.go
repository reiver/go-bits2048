package bits2048

import (
	"strconv"
)

// Uint is a 2048-bit unsigned-number.
//
// Uint stores the data as an array of uint.
// And stores it in little-endian order.
// So that Uint[0] is the least-significant values,
// and Uint[31] is the most-significant.
type Uint struct {
	value [2048/64]uint64
}

// FromUint64 returns a bits2048.Uint from a Go uint64.
func FromUint64(value uint64) Uint {
	var result Uint
	result.FromUint64(value)
	return result
}

func FromUint64s(values ...uint64) (Uint, error) {
	var result Uint
	err := result.FromUint64s(values...)
	return result, err
}

func MustFromUint64s(values ...uint64) Uint {
	result, err := FromUint64s(values...)
	if nil != err {
		panic(err)
	}
	return result
}

// Cmp compared 'receiver' and 'other', and —
//
// It returns 0 if 'receiver' and 'other' are equal.
// It return -1 if 'receiver' is less-than 'other'.
// And it return 1 if 'receiver' is greater-than 'other'.
func (receiver *Uint) Cmp(other *Uint) int {
	if receiver == other {
		return 0
	}
	if *receiver == *other {
		return 0
	}

	var length int =len(receiver.value)
	for i:=(length-1); i>=0; i-- {
		n1 := receiver.value[i]
		n2 := other.value[i]

		switch {
		case n1 < n2:
			return -1
		case n1 > n2:
			return 1
		}
	}

	return 0
}

func (receiver *Uint) Add(src1 *Uint, src2 *Uint) error {
	if nil == receiver {
		return errNilReceiver
	}
	if nil == src1 {
		return errNilSourceOne
	}
	if nil == src2 {
		return errNilSourceTwo
	}

	var overflow bool
	for i, n1 := range src1.value {
		n2 := src2.value[i]

		var carry bool = overflow

		result := n1 + n2
		overflow = result < n1
		if carry {
			result++
		}

		receiver.value[i] = result
	}

	if overflow {
		return ErrOverFlow
	}

	return nil
}

func (receiver *Uint) FromUint64(value uint64) {
	if nil == receiver {
		panic(errNilReceiver)
	}

	receiver.value = [len(receiver.value)]uint64{}
	receiver.value[0] = value
}

func (receiver *Uint) FromUint64s(values ...uint64) error {
	if nil == receiver {
		return errNilReceiver
	}

	var lenvalues int = len(values)

	var length int =len(receiver.value)
	if length < lenvalues {
		return errTooMany
	}

	var empty Uint
	receiver.value = empty.value

	for i, value := range values {
		receiver.value[i] = value
	}

	return nil
}

func (receiver Uint) GoString() string {
	var p []byte

	p = append(p, "bits2048.MustFromUint64s("...)
	for i, part := range receiver.value {
		if 0 < i {
			p = append(p, ", "...)
		}

		p = append(p, "0x"...)
		p = strconv.AppendUint(p, part, 16)
	}
	p = append(p, ')')

	return string(p)
}
