package bits2048

import (
	"testing"
)

func TestUint_FromUint64(t *testing.T) {

	tests := []struct{
		Value uint64
		Expected Uint
	}{
		{
			Value: 0,
			Expected: Uint{},
		},
		{
			Value:                    1,
			Expected: Uint{[32]uint64{1}},
		},
		{
			Value:                    2,
			Expected: Uint{[32]uint64{2}},
		},
		{
			Value:                    3,
			Expected: Uint{[32]uint64{3}},
		},
		{
			Value:                    4,
			Expected: Uint{[32]uint64{4}},
		},
		{
			Value:                    5,
			Expected: Uint{[32]uint64{5}},
		},
		{
			Value:                    6,
			Expected: Uint{[32]uint64{6}},
		},
		{
			Value:                    7,
			Expected: Uint{[32]uint64{7}},
		},
		{
			Value:                    8,
			Expected: Uint{[32]uint64{8}},
		},
		{
			Value:                    9,
			Expected: Uint{[32]uint64{9}},
		},
		{
			Value:                    10,
			Expected: Uint{[32]uint64{10}},
		},
		{
			Value:                    11,
			Expected: Uint{[32]uint64{11}},
		},
		{
			Value:                    12,
			Expected: Uint{[32]uint64{12}},
		},
		{
			Value:                    13,
			Expected: Uint{[32]uint64{13}},
		},
		{
			Value:                    14,
			Expected: Uint{[32]uint64{14}},
		},
		{
			Value:                    15,
			Expected: Uint{[32]uint64{15}},
		},



		{
			Value:                    0x1234,
			Expected: Uint{[32]uint64{0x1234}},
		},



		{
			Value:                    0xFFFF,
			Expected: Uint{[32]uint64{0xFFFF}},
		},



		{
			Value:                    0x123456,
			Expected: Uint{[32]uint64{0x123456}},
		},



		{
			Value:                    0xFFFFFF,
			Expected: Uint{[32]uint64{0xFFFFFF}},
		},



		{
			Value:                    0x12345678,
			Expected: Uint{[32]uint64{0x12345678}},
		},



		{
			Value:                    0xFFFFFFFF,
			Expected: Uint{[32]uint64{0xFFFFFFFF}},
		},



		{
			Value:                    0x123456789A,
			Expected: Uint{[32]uint64{0x123456789A}},
		},



		{
			Value:                    0xFFFFFFFFFF,
			Expected: Uint{[32]uint64{0xFFFFFFFFFF}},
		},



		{
			Value:                    0x123456789ABC,
			Expected: Uint{[32]uint64{0x123456789ABC}},
		},



		{
			Value:                    0xFFFFFFFFFFFF,
			Expected: Uint{[32]uint64{0xFFFFFFFFFFFF}},
		},



		{
			Value:                    0x123456789ABCDE,
			Expected: Uint{[32]uint64{0x123456789ABCDE}},
		},



		{
			Value:                    0xFFFFFFFFFFFFFF,
			Expected: Uint{[32]uint64{0xFFFFFFFFFFFFFF}},
		},



		{
			Value:                    0x123456789ABCDEF5,
			Expected: Uint{[32]uint64{0x123456789ABCDEF5}},
		},



		{
			Value:                    0xFFFFFFFFFFFFFFFF,
			Expected: Uint{[32]uint64{0xFFFFFFFFFFFFFFFF}},
		},



		{
			Value:                    0x8429BD64DE2194FC,
			Expected: Uint{[32]uint64{0x8429BD64DE2194FC}},
		},
	}

	for testNumber, test := range tests {

		var actual Uint
		actual.FromUint64(test.Value)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", expected)
			t.Logf("ACTUAL  : %v", actual)
			continue
		}
	}
}
