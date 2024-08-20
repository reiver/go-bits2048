package bits2048

import (
	"testing"
)

func TestString_GoString(t *testing.T) {

	tests := []struct{
		String String
		Expected string
	}{
		{
			String:String{[256]byte{}},
			Expected:`bits2048.MustFromString("")`,
		},



		{
			String:String{[256]byte{0}},
			Expected:`bits2048.MustFromString("")`,
		},
		{
			String:String{[256]byte{1}},
			Expected:`bits2048.MustFromString("\x00")`,
		},
		{
			String:String{[256]byte{2}},
			Expected:`bits2048.MustFromString("\x00\x00")`,
		},
		{
			String:String{[256]byte{3}},
			Expected:`bits2048.MustFromString("\x00\x00\x00")`,
		},
		{
			String:String{[256]byte{4}},
			Expected:`bits2048.MustFromString("\x00\x00\x00\x00")`,
		},



		{
			String:String{[256]byte{0,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("")`,
		},
		{
			String:String{[256]byte{1,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("H")`,
		},
		{
			String:String{[256]byte{2,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("He")`,
		},
		{
			String:String{[256]byte{3,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hel")`,
		},
		{
			String:String{[256]byte{4,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hell")`,
		},
		{
			String:String{[256]byte{5,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello")`,
		},
		{
			String:String{[256]byte{6,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello ")`,
		},
		{
			String:String{[256]byte{7,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello w")`,
		},
		{
			String:String{[256]byte{8,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello wo")`,
		},
		{
			String:String{[256]byte{9,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello wor")`,
		},
		{
			String:String{[256]byte{10,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello worl")`,
		},
		{
			String:String{[256]byte{11,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello world")`,
		},
		{
			String:String{[256]byte{12,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello world!")`,
		},
		{
			String:String{[256]byte{13,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello world!\x00")`,
		},
		{
			String:String{[256]byte{14,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello world!\x00\x00")`,
		},
		{
			String:String{[256]byte{15,'H','e','l','l','o',' ','w','o','r','l','d','!'}},
			Expected:`bits2048.MustFromString("Hello world!\x00\x00\x00")`,
		},



		{
			String:            MustFromString("ONCE TWICE THRICE FOURCE"),
			Expected:`bits2048.MustFromString("ONCE TWICE THRICE FOURCE")`,
		},
		{
			String:            MustFromString("Hi! 🙂"),
			Expected:`bits2048.MustFromString("Hi! 🙂")`,
		},
		{
			String:            MustFromString("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"),
			Expected:`bits2048.MustFromString("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")`,
		},
		{
			String:            MustFromString("ODzCIi2TFDU3AVtJXzQt6BgVwrMu0Q6gxAyzvjqblfXMkL73aunbukkEdMPn9QHntlI0B3oMDKhcsEvQIDuzUXPev1aMRAPJCs6x5GVnHTuv8SOfbv07e2RU61q6R8OAxA02y3bjyGbLoMbzQjlsZAAWhgeB9KW8hIgeyo7eWnZrqTmpkQxXszjdqT6Zh7SEijpnuGfk28XmuSjlsBOyA4sfMYpwUZg48uliDpkXJWdT3aR3atqyMJQm3NwCR51"),
			Expected:`bits2048.MustFromString("ODzCIi2TFDU3AVtJXzQt6BgVwrMu0Q6gxAyzvjqblfXMkL73aunbukkEdMPn9QHntlI0B3oMDKhcsEvQIDuzUXPev1aMRAPJCs6x5GVnHTuv8SOfbv07e2RU61q6R8OAxA02y3bjyGbLoMbzQjlsZAAWhgeB9KW8hIgeyo7eWnZrqTmpkQxXszjdqT6Zh7SEijpnuGfk28XmuSjlsBOyA4sfMYpwUZg48uliDpkXJWdT3aR3atqyMJQm3NwCR51")`,
		},
	}

	for testNumber, test := range tests {

		actual := test.String.GoString()
		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual 'string' is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			continue
		}
	}
}
