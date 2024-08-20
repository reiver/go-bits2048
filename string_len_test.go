package bits2048

import (
	"testing"
)

func TestString_Len(t *testing.T) {

	tests := []struct{
		String String
		Expected int
	}{
		{
			String: String{[256]byte{}},
			Expected:0,
		},



		{
			String: String{[256]byte{0}},
			Expected:0,
		},
		{
			String: String{[256]byte{1}},
			Expected:1,
		},
		{
			String: String{[256]byte{2}},
			Expected:2,
		},
		{
			String: String{[256]byte{3}},
			Expected:3,
		},
		{
			String: String{[256]byte{4}},
			Expected:4,
		},



		{
			String: String{[256]byte{0,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:0,
		},
		{
			String: String{[256]byte{1,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:1,
		},
		{
			String: String{[256]byte{2,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:2,
		},
		{
			String: String{[256]byte{3,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:3,
		},
		{
			String: String{[256]byte{4,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:4,
		},
		{
			String: String{[256]byte{5,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:5,
		},
		{
			String: String{[256]byte{6,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:6,
		},
		{
			String: String{[256]byte{7,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:7,
		},
		{
			String: String{[256]byte{8,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:8,
		},
		{
			String: String{[256]byte{9,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:9,
		},
		{
			String: String{[256]byte{10,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:10,
		},
		{
			String: String{[256]byte{11,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:11,
		},
		{
			String: String{[256]byte{12,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:12,
		},
		{
			String: String{[256]byte{13,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:13,
		},
		{
			String: String{[256]byte{14,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:14,
		},
		{
			String: String{[256]byte{15,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:15,
		},
	}

	for testNumber, test := range tests {

		actual := test.String.Len()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'len' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %d", expected)
			t.Logf("ACTUAL:   %d", actual)
			continue
		}
	}
}
