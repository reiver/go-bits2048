package bits2048

import (
	"testing"
)

func TestString_Bytes(t *testing.T) {

	tests := []struct{
		String String
		Expected string
	}{
		{
			String:  String{[256]byte{}},
			Expected:"",
		},



		{
			String:  String{[256]byte{0}},
			Expected:"",
		},
		{
			String:  String{[256]byte{1}},
			Expected:"\x00",
		},
		{
			String:  String{[256]byte{2}},
			Expected:"\x00\x00",
		},
		{
			String:  String{[256]byte{3}},
			Expected:"\x00\x00\x00",
		},
		{
			String:  String{[256]byte{4}},
			Expected:"\x00\x00\x00\x00",
		},



		{
			String:  String{[256]byte{0,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"",
		},
		{
			String:  String{[256]byte{1,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"H",
		},
		{
			String:  String{[256]byte{2,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"He",
		},
		{
			String:  String{[256]byte{3,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hel",
		},
		{
			String:  String{[256]byte{4,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hell",
		},
		{
			String:  String{[256]byte{5,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello",
		},
		{
			String:  String{[256]byte{6,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello ",
		},
		{
			String:  String{[256]byte{7,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello w",
		},
		{
			String:  String{[256]byte{8,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello wo",
		},
		{
			String:  String{[256]byte{9,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello wor",
		},
		{
			String:  String{[256]byte{10,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello worl",
		},
		{
			String:  String{[256]byte{11,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello world",
		},
		{
			String:  String{[256]byte{12,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello world!",
		},
		{
			String:  String{[256]byte{13,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello world!\x00",
		},
		{
			String:  String{[256]byte{14,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello world!\x00\x00",
		},
		{
			String:  String{[256]byte{15,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:"Hello world!\x00\x00\x00",
		},
	}

	for testNumber, test := range tests {

		actualBytes := test.String.Bytes()

		actual := string(actualBytes)
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'String' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			continue
		}
	}
}
