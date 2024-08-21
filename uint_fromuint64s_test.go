package bits2048

import (
	"testing"
)

func TestUint_FromUint64s(t *testing.T) {

	tests := []struct{
		Values []uint64
		Expected Uint
	}{
		{
			Values:          []uint64{0},
			Expected: Uint{},
		},
		{
			Values:          []uint64{1},
			Expected: Uint{[32]uint64{1}},
		},
		{
			Values:          []uint64{2},
			Expected: Uint{[32]uint64{2}},
		},
		{
			Values:          []uint64{3},
			Expected: Uint{[32]uint64{3}},
		},
		{
			Values:          []uint64{4},
			Expected: Uint{[32]uint64{4}},
		},
		{
			Values:          []uint64{5},
			Expected: Uint{[32]uint64{5}},
		},
		{
			Values:          []uint64{6},
			Expected: Uint{[32]uint64{6}},
		},
		{
			Values:          []uint64{7},
			Expected: Uint{[32]uint64{7}},
		},
		{
			Values:          []uint64{8},
			Expected: Uint{[32]uint64{8}},
		},
		{
			Values:          []uint64{9},
			Expected: Uint{[32]uint64{9}},
		},
		{
			Values:          []uint64{10},
			Expected: Uint{[32]uint64{10}},
		},
		{
			Values:          []uint64{11},
			Expected: Uint{[32]uint64{11}},
		},
		{
			Values:          []uint64{12},
			Expected: Uint{[32]uint64{12}},
		},
		{
			Values:          []uint64{13},
			Expected: Uint{[32]uint64{13}},
		},
		{
			Values:          []uint64{14},
			Expected: Uint{[32]uint64{14}},
		},
		{
			Values:          []uint64{15},
			Expected: Uint{[32]uint64{15}},
		},



		{
			Values:          []uint64{0x1234},
			Expected: Uint{[32]uint64{0x1234}},
		},



		{
			Values:          []uint64{0xFFFF},
			Expected: Uint{[32]uint64{0xFFFF}},
		},



		{
			Values:          []uint64{0x123456},
			Expected: Uint{[32]uint64{0x123456}},
		},



		{
			Values:          []uint64{0xFFFFFF},
			Expected: Uint{[32]uint64{0xFFFFFF}},
		},



		{
			Values:          []uint64{0x12345678},
			Expected: Uint{[32]uint64{0x12345678}},
		},



		{
			Values:          []uint64{0xFFFFFFFF},
			Expected: Uint{[32]uint64{0xFFFFFFFF}},
		},



		{
			Values:          []uint64{0x123456789A},
			Expected: Uint{[32]uint64{0x123456789A}},
		},



		{
			Values:          []uint64{0xFFFFFFFFFF},
			Expected: Uint{[32]uint64{0xFFFFFFFFFF}},
		},



		{
			Values:          []uint64{0x123456789ABC},
			Expected: Uint{[32]uint64{0x123456789ABC}},
		},



		{
			Values:          []uint64{0xFFFFFFFFFFFF},
			Expected: Uint{[32]uint64{0xFFFFFFFFFFFF}},
		},



		{
			Values:          []uint64{0x123456789ABCDE},
			Expected: Uint{[32]uint64{0x123456789ABCDE}},
		},



		{
			Values:          []uint64{0xFFFFFFFFFFFFFF},
			Expected: Uint{[32]uint64{0xFFFFFFFFFFFFFF}},
		},



		{
			Values:          []uint64{0x123456789ABCDEF5},
			Expected: Uint{[32]uint64{0x123456789ABCDEF5}},
		},



		{
			Values:          []uint64{0xFFFFFFFFFFFFFFFF},
			Expected: Uint{[32]uint64{0xFFFFFFFFFFFFFFFF}},
		},



		{
			Values:          []uint64{0x8429BD64DE2194FC},
			Expected: Uint{[32]uint64{0x8429BD64DE2194FC}},
		},



		{
			Values:          []uint64{0x872b4dbc7030ba8d,0x980ff4cffa7d30a3,0x932f5456de614e1f,0x7065282edbc1e013,0x431e107ff3f1d733,0x5d19a3cf47a981f3,0xe3f5d16db149b369,0xf9b40d0eaa062abb,0xefba9470fb343c9c,0xfa5c0b5259bee26c,0x85e79ffabfca9904,0xc4c2672fc1329e38,0x89f6e18bc04e42c3,0xe4ca8f01c99078f0,0xfd68201ecfd1f709,0x17796e6b6828111e,0xebab18f21f6a6196,0x646ae5f59b92dad3,0x05aba942ae8d25eb,0x44973979252b96b2,0x42290dc908f936ef,0xb8cb0ac6360eadd0,0x0ba4fb7ee6e98003,0x78cf9acc7098228f,0x9d44de6e8f6160c1,0xf403ada528813b1c,0xc14d4dbfcef00d75,0x5bb9a62d2c886eef,0x7ea11a2b395434f1,0xdd59cd835f0d6055,0x363da2cb963e4099,0x42e5f10018b4c85b},
			Expected: Uint{[32]uint64{0x872b4dbc7030ba8d,0x980ff4cffa7d30a3,0x932f5456de614e1f,0x7065282edbc1e013,0x431e107ff3f1d733,0x5d19a3cf47a981f3,0xe3f5d16db149b369,0xf9b40d0eaa062abb,0xefba9470fb343c9c,0xfa5c0b5259bee26c,0x85e79ffabfca9904,0xc4c2672fc1329e38,0x89f6e18bc04e42c3,0xe4ca8f01c99078f0,0xfd68201ecfd1f709,0x17796e6b6828111e,0xebab18f21f6a6196,0x646ae5f59b92dad3,0x05aba942ae8d25eb,0x44973979252b96b2,0x42290dc908f936ef,0xb8cb0ac6360eadd0,0x0ba4fb7ee6e98003,0x78cf9acc7098228f,0x9d44de6e8f6160c1,0xf403ada528813b1c,0xc14d4dbfcef00d75,0x5bb9a62d2c886eef,0x7ea11a2b395434f1,0xdd59cd835f0d6055,0x363da2cb963e4099,0x42e5f10018b4c85b}},
		},
	}

	for testNumber, test := range tests {

		var actual Uint
		actual.FromUint64s(test.Values...)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual value is not what was expected.", testNumber)
			t.Logf("EXPECTED: %v", expected)
			t.Logf("ACTUAL  : %v", actual)
			continue
		}
	}
}
