package is

import (
	"fmt"
	"strings"
	"testing"
)

func TestAlpha(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", true},
		{"소주", false},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", false}, //UTF-8(ASCII): 0
		{"123", false},
		{"0123", false},
		{"-00123", false},
		{"0", false},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := Alpha(test.param)
		if actual != test.expected {
			t.Errorf("Expected Alpha(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestUTFLetter(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"", true},
		{"abc", true},
		{"소주", true},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", true},
		{"소", true},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", false}, //UTF-8(ASCII): 0
		{"123", false},
		{"0123", false},
		{"-00123", false},
		{"0", false},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", true},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := UTFLetter(test.param)
		if actual != test.expected {
			t.Errorf("Expected UTFLetter(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestAlphanumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc123", true},
		{"ABC111", true},
		{"abc1", true},
		{"abc〩", false},
		{"abc", true},
		{"소주", false},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"0", true},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", true},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := Alphanumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected Alphanumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestUTFLetterNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", true},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", true},
		{"abc〩", true},
		{"abc", true},
		{"소주", true},
		{"ABC", true},
		{"FoObAr", true},
		{"소aBC", true},
		{"소", true},
		{"달기&Co.", false},
		{"〩Hours", true},
		{"\ufff0", false},
		{"\u0070", true},  //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"0", true},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"-1¾", false},
		{"1¾", true},
		{"〥〩", true},
		{"모자", true},
		{"ix", true},
		{"۳۵۶۰", true},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := UTFLetterNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected UTFLetterNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", false},
		{"소주", false},
		{"ABC", false},
		{"FoObAr", false},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", false}, //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"+00123", false},
		{"0", true},
		{"-0", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"12𐅪3", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", false},
		{"۳۵۶۰", false},
		{"1--", false},
		{"1-1", false},
		{"-", false},
		{"--", false},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := Numeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected Numeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestUTFNumeric(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", true},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", false},
		{"소주", false},
		{"ABC", false},
		{"FoObAr", false},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", false}, //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"0", true},
		{"-0", false},
		{"--0", false},
		{"-0-", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"12𐅪3", true},
		{"-1¾", false},
		{"1¾", true},
		{"〥〩", true},
		{"모자", false},
		{"ix", false},
		{"۳۵۶۰", true},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := UTFNumeric(test.param)
		if actual != test.expected {
			t.Errorf("Expected UTFNumeric(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestUTFDigit(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"\n", false},
		{"\r", false},
		{"Ⅸ", false},
		{"", true},
		{"   fooo   ", false},
		{"abc!!!", false},
		{"abc1", false},
		{"abc〩", false},
		{"abc", false},
		{"소주", false},
		{"ABC", false},
		{"FoObAr", false},
		{"소aBC", false},
		{"소", false},
		{"달기&Co.", false},
		{"〩Hours", false},
		{"\ufff0", false},
		{"\u0070", false}, //UTF-8(ASCII): p
		{"\u0026", false}, //UTF-8(ASCII): &
		{"\u0030", true},  //UTF-8(ASCII): 0
		{"123", true},
		{"0123", true},
		{"-00123", false},
		{"0", true},
		{"-0", false},
		{"--0", false},
		{"-0-", false},
		{"123.123", false},
		{" ", false},
		{".", false},
		{"12𐅪3", false},
		{"1483920", true},
		{"", true},
		{"۳۵۶۰", true},
		{"-29", false},
		{"-1¾", false},
		{"1¾", false},
		{"〥〩", false},
		{"모자", false},
		{"ix", false},
		{"۳۵۶۰", true},
		{"1++", false},
		{"1+1", false},
		{"+", false},
		{"++", false},
		{"+1", false},
	}
	for _, test := range tests {
		actual := UTFDigit(test.param)
		if actual != test.expected {
			t.Errorf("Expected UTFDigit(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestLowerCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc123", true},
		{"abc", true},
		{"a b c", true},
		{"abcß", true},
		{"abcẞ", false},
		{"ABCẞ", false},
		{"tr竪s 端ber", true},
		{"fooBar", false},
		{"123ABC", false},
		{"ABC123", false},
		{"ABC", false},
		{"S T R", false},
		{"fooBar", false},
		{"abacaba123", true},
	}
	for _, test := range tests {
		actual := LowerCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected LowerCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestUpperCase(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc123", false},
		{"abc", false},
		{"a b c", false},
		{"abcß", false},
		{"abcẞ", false},
		{"ABCẞ", true},
		{"tr竪s 端ber", false},
		{"fooBar", false},
		{"123ABC", true},
		{"ABC123", true},
		{"ABC", true},
		{"S T R", true},
		{"fooBar", false},
		{"abacaba123", false},
	}
	for _, test := range tests {
		actual := UpperCase(test.param)
		if actual != test.expected {
			t.Errorf("Expected UpperCase(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestInt(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"-2147483648", true},           //Signed 32 Bit Min Int
		{"2147483647", true},            //Signed 32 Bit Max Int
		{"-2147483649", true},           //Signed 32 Bit Min Int - 1
		{"2147483648", true},            //Signed 32 Bit Max Int + 1
		{"4294967295", true},            //Unsigned 32 Bit Max Int
		{"4294967296", true},            //Unsigned 32 Bit Max Int + 1
		{"-9223372036854775808", true},  //Signed 64 Bit Min Int
		{"9223372036854775807", true},   //Signed 64 Bit Max Int
		{"-9223372036854775809", false}, //Signed 64 Bit Min Int - 1
		{"9223372036854775808", false},  //Signed 64 Bit Max Int + 1
		{"18446744073709551615", false}, //Unsigned 64 Bit Max Int
		{"18446744073709551616", false}, //Unsigned 64 Bit Max Int + 1
		{"", true},
		{"123", true},
		{"0", true},
		{"-0", true},
		{"+0", true},
		{"01", true},
		{"123.123", false},
		{" ", false},
		{"000", true},
	}
	for _, test := range tests {
		actual := Int(test.param)
		if actual != test.expected {
			t.Errorf("Expected Int(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestEmail(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{``, false},
		{`foo@bar.com`, true},
		{`x@x.x`, true},
		{`foo@bar.com.au`, true},
		{`foo+bar@bar.com`, true},
		{`foo@bar.coffee`, true},
		{`foo@bar.中文网`, true},
		{`invalidemail@`, false},
		{`invalid.com`, false},
		{`@invalid.com`, false},
		{`test|123@m端ller.com`, true},
		{`hans@m端ller.com`, true},
		{`hans.m端ller@test.com`, true},
		{`NathAn.daVIeS@DomaIn.cOM`, true},
		{`NATHAN.DAVIES@DOMAIN.CO.UK`, true},
		{`very.(),:;<>[]".VERY."very@\ "very".unusual@strange.example.com`, true},
	}
	for _, test := range tests {
		actual := Email(test.param)
		if actual != test.expected {
			t.Errorf("Expected Email(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func ExampleEmail() {
	fmt.Println(Email("jhon@example.com"))
	fmt.Println(Email("invalid.com"))
	fmt.Println(Email(`very.(),:;<>[]".VERY."very@\ "very".unusual@strange.example.com`))
	// Output:
	// true
	// false
	// true
}

func TestURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"http://foo.bar#com", true},
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", true},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"ftp.foo.bar", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://user:pass@www.foobar.com/path/file", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"http://foobar.com/a-", true},
		{"http://foobar.پاکستان/", true},
		{"http://foobar.c_o_m", false},
		{"", false},
		{"xyz://foobar.com", false},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", false},
		{"http://www.foo_bar.com/", false},
		{"http://localhost:3000/", true},
		{"http://foobar.com#baz=qux", true},
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", false},
		{"http://www.foo---bar.com/", false},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", false},
		{"irc://#channel@network", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
		{"http://foo^bar.org", false},
		{"http://foo&*bar.org", false},
		{"http://foo&bar.org", false},
		{"http://foo bar.org", false},
		{"http://foo.bar.org", true},
		{"http://www.foo.bar.org", true},
		{"http://www.foo.co.uk", true},
		{"foo", false},
		{"http://.foo.com", false},
		{"http://,foo.com", false},
		{",foo.com", false},
		// according to issues #62 #66
		{"https://pbs.twimg.com/profile_images/560826135676588032/j8fWrmYY_normal.jpeg", true},
		{"http://me.example.com", true},
		{"http://www.me.example.com", true},
		{"https://farm6.static.flickr.com", true},
		{"https://zh.wikipedia.org/wiki/Wikipedia:%E9%A6%96%E9%A1%B5", true},
		{"google", false},
		// According to #87
		{"http://hyphenated-host-name.example.co.in", true},
		{"http://cant-end-with-hyphen-.example.com", false},
		{"http://-cant-start-with-hyphen.example.com", false},
		{"http://www.domain-can-have-dashes.com", true},
	}
	for _, test := range tests {
		actual := URL(test.param)
		if actual != test.expected {
			t.Errorf("Expected URL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestRequestURL(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		// {"http://foo.bar#com", false}, goversion < 1.6: true, goversion > 1.5: false
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", false},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"", false},
		{"xyz://foobar.com", true},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		// {"http://foobar.com#baz=qux", false}, goversion < 1.6: true, goversion > 1.5: false
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", true},
		{"irc://#channel@network", true},
		{"/abs/test/dir", false},
		{"./rel/test/dir", false},
	}
	for _, test := range tests {
		actual := RequestURL(test.param)
		if actual != test.expected {
			t.Errorf("Expected RequestURL(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestRequestURI(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		// {"http://foo.bar#com", false}, goversion < 1.6: true, goversion > 1.5: false
		{"http://foobar.com", true},
		{"https://foobar.com", true},
		{"foobar.com", false},
		{"http://foobar.coffee/", true},
		{"http://foobar.中文网/", true},
		{"http://foobar.org/", true},
		{"http://foobar.org:8080/", true},
		{"ftp://foobar.ru/", true},
		{"http://user:pass@www.foobar.com/", true},
		{"http://127.0.0.1/", true},
		{"http://duckduckgo.com/?q=%2F", true},
		{"http://localhost:3000/", true},
		{"http://foobar.com/?foo=bar#baz=qux", true},
		{"http://foobar.com?foo=bar", true},
		{"http://www.xn--froschgrn-x9a.net/", true},
		{"xyz://foobar.com", true},
		{"invalid.", false},
		{".com", false},
		{"rtmp://foobar.com", true},
		{"http://www.foo_bar.com/", true},
		{"http://localhost:3000/", true},
		// {"http://foobar.com#baz=qux", false}, goversion < 1.6: true, goversion > 1.5: false
		{"http://foobar.com/t$-_.+!*\\'(),", true},
		{"http://www.foobar.com/~foobar", true},
		{"http://www.-foobar.com/", true},
		{"http://www.foo---bar.com/", true},
		{"mailto:someone@example.com", true},
		{"irc://irc.server.org/channel", true},
		{"irc://#channel@network", true},
		{"/abs/test/dir", true},
		{"./rel/test/dir", false},
	}
	for _, test := range tests {
		actual := RequestURI(test.param)
		if actual != test.expected {
			t.Errorf("Expected RequestURI(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestFloat(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"  ", false},
		{"-.123", true},
		{"abacaba", false},
		{"1f", false},
		{"-1f", false},
		{"+1f", false},
		{"123", true},
		{"123.", true},
		{"123.123", true},
		{"-123.123", true},
		{"+123.123", true},
		{"0.123", true},
		{"-0.123", true},
		{"+0.123", true},
		{".0", true},
		{"01.123", true},
		{"-0.22250738585072011e-307", true},
		{"+0.22250738585072011e-307", true},
	}
	for _, test := range tests {
		actual := Float(test.param)
		if actual != test.expected {
			t.Errorf("Expected Float(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestHexadecimal(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abcdefg", false},
		{"", false},
		{"..", false},
		{"deadBEEF", true},
		{"ff0044", true},
	}
	for _, test := range tests {
		actual := Hexadecimal(test.param)
		if actual != test.expected {
			t.Errorf("Expected Hexadecimal(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestHexcolor(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"#ff", false},
		{"fff0", false},
		{"#ff12FG", false},
		{"CCccCC", true},
		{"fff", true},
		{"#f00", true},
	}
	for _, test := range tests {
		actual := Hexcolor(test.param)
		if actual != test.expected {
			t.Errorf("Expected Hexcolor(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestRGBcolor(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"rgb(0,31,255)", true},
		{"rgb(1,349,275)", false},
		{"rgb(01,31,255)", false},
		{"rgb(0.6,31,255)", false},
		{"rgba(0,31,255)", false},
		{"rgb(0,  31, 255)", true},
	}
	for _, test := range tests {
		actual := RGBcolor(test.param)
		if actual != test.expected {
			t.Errorf("Expected RGBcolor(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestNull(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abacaba", false},
		{"", true},
	}
	for _, test := range tests {
		actual := NullString(test.param)
		if actual != test.expected {
			t.Errorf("Expected NullString(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestDivibleBy(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param1   string
		param2   string
		expected bool
	}{
		{"4", "2", true},
		{"100", "10", true},
		{"", "1", true},
		{"123", "foo", false},
		{"123", "0", false},
	}
	for _, test := range tests {
		actual := DivisibleBy(test.param1, test.param2)
		if actual != test.expected {
			t.Errorf("Expected DivisibleBy(%q, %q) to be %v, got %v", test.param1, test.param2, test.expected, actual)
		}
	}
}

// This small example illustrate how to work with IsDivisibleBy function.
func ExampleDivisibleBy() {
	println("1024  divisible by 64: ", DivisibleBy("1024", "64"))
}

func TestByteLength(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param1   string
		param2   int
		param3   int
		expected bool
	}{
		{"abacaba", 100, -1, false},
		{"abacaba", 1, 3, false},
		{"abacaba", 1, 7, true},
		{"abacaba", 0, 8, true},
		{"\ufff0", 1, 1, false},
	}
	for _, test := range tests {
		actual := ByteLength(test.param1, test.param2, test.param3)
		if actual != test.expected {
			t.Errorf("Expected ByteLength(%q, %q, %q) to be %v, got %v", test.param1, test.param2, test.param3, test.expected, actual)
		}
	}
}

func TestJSON(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"145", true},
		{"asdf", false},
		{"123:f00", false},
		{"{\"Name\":\"Alice\",\"Body\":\"Hello\",\"Time\":1294706395881547000}", true},
		{"{}", true},
		{"{\"Key\":{\"Key\":{\"Key\":123}}}", true},
		{"[]", true},
		{"null", true},
	}
	for _, test := range tests {
		actual := JSON(test.param)
		if actual != test.expected {
			t.Errorf("Expected JSON(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestMultibyte(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"abc", false},
		{"123", false},
		{"<>@;.-=", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"あいうえお foobar", true},
		{"test＠example.com", true},
		{"test＠example.com", true},
		{"1234abcDEｘｙｚ", true},
		{"ｶﾀｶﾅ", true},
		{"~", false},
	}
	for _, test := range tests {
		actual := Multibyte(test.param)
		if actual != test.expected {
			t.Errorf("Expected Multibyte(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestASCII(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"", true},
	}
	for _, test := range tests {
		actual := ASCII(test.param)
		if actual != test.expected {
			t.Errorf("Expected ASCII(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestPrintableASCII(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ｆｏｏbar", false},
		{"ｘｙｚ０９８", false},
		{"１２３456", false},
		{"ｶﾀｶﾅ", false},
		{"foobar", true},
		{"0987654321", true},
		{"test@example.com", true},
		{"1234abcDEF", true},
		{"newline\n", false},
		{"\x19test\x7F", false},
	}
	for _, test := range tests {
		actual := PrintableASCII(test.param)
		if actual != test.expected {
			t.Errorf("Expected PrintableASCII(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestFullWidth(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"abc", false},
		{"abc123", false},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", false},
		{"ひらがな・カタカナ、．漢字", true},
		{"３ー０　ａ＠ｃｏｍ", true},
		{"Ｆｶﾀｶﾅﾞﾬ", true},
		{"Good＝Parts", true},
		{"", true},
	}
	for _, test := range tests {
		actual := FullWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected FullWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestHalfWidth(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"あいうえお", false},
		{"００１１", false},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", true},
		{"l-btn_02--active", true},
		{"abc123い", true},
		{"ｶﾀｶﾅﾞﾬ￩", true},
		{"", true},
	}
	for _, test := range tests {
		actual := HalfWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected HalfWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestVariableWidth(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", true},
		{"ひらがなカタカナ漢字ABCDE", true},
		{"３ー０123", true},
		{"Ｆｶﾀｶﾅﾞﾬ", true},
		{"", true},
		{"Good＝Parts", true},
		{"abc", false},
		{"abc123", false},
		{"!\"#$%&()<>/+=-_? ~^|.,@`{}[]", false},
		{"ひらがな・カタカナ、．漢字", false},
		{"１２３４５６", false},
		{"ｶﾀｶﾅﾞﾬ", false},
	}
	for _, test := range tests {
		actual := VariableWidth(test.param)
		if actual != test.expected {
			t.Errorf("Expected VariableWidth(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestUUID(t *testing.T) {
	t.Parallel()

	// Tests without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3xxx", false},
		{"a987fbc94bed3078cf079141ba07c9f3", false},
		{"934859", false},
		{"987fbc9-4bed-3078-cf07a-9141ba07c9f3", false},
		{"aaaaaaaa-1111-1111-aaag-111111111111", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}
	for _, test := range tests {
		actual := UUID(test.param)
		if actual != test.expected {
			t.Errorf("Expected UUID(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// UUID ver. 3
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"412452646", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-4078-8f07-9141ba07c9f3", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", true},
	}
	for _, test := range tests {
		actual := UUIDv3(test.param)
		if actual != test.expected {
			t.Errorf("Expected UUIDv3(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// UUID ver. 4
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"a987fbc9-4bed-5078-af07-9141ba07c9f3", false},
		{"934859", false},
		{"57b73598-8764-4ad0-a76a-679bb6640eb1", true},
		{"625e63f3-58f5-40b7-83a1-a72ad31acffb", true},
	}
	for _, test := range tests {
		actual := UUIDv4(test.param)
		if actual != test.expected {
			t.Errorf("Expected UUIDv4(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// UUID ver. 5
	tests = []struct {
		param    string
		expected bool
	}{

		{"", false},
		{"xxxa987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"9c858901-8a57-4791-81fe-4c455b099bc9", false},
		{"a987fbc9-4bed-3078-cf07-9141ba07c9f3", false},
		{"987fbc97-4bed-5078-af07-9141ba07c9f3", true},
		{"987fbc97-4bed-5078-9f07-9141ba07c9f3", true},
	}
	for _, test := range tests {
		actual := UUIDv5(test.param)
		if actual != test.expected {
			t.Errorf("Expected UUIDv5(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestCreditCard(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"5398228707871528", false},
		{"375556917985515", true},
		{"36050234196908", true},
		{"4716461583322103", true},
		{"4716-2210-5188-5662", true},
		{"4929 7226 5379 7141", true},
		{"5398228707871527", true},
	}
	for _, test := range tests {
		actual := CreditCard(test.param)
		if actual != test.expected {
			t.Errorf("Expected CreditCard(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestBN(t *testing.T) {
	t.Parallel()

	// Without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3836221195", true},
		{"1-61729-085-8", true},
		{"3 423 21412 0", true},
		{"3 401 01319 X", true},
		{"9784873113685", true},
		{"978-4-87311-368-5", true},
		{"978 3401013190", true},
		{"978-3-8362-2119-1", true},
	}
	for _, test := range tests {
		actual := ISBN(test.param, -1)
		if actual != test.expected {
			t.Errorf("Expected ISBN(%q, -1) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// ISBN 10
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3423214121", false},
		{"978-3836221191", false},
		{"3-423-21412-1", false},
		{"3 423 21412 1", false},
		{"3836221195", true},
		{"1-61729-085-8", true},
		{"3 423 21412 0", true},
		{"3 401 01319 X", true},
	}
	for _, test := range tests {
		actual := ISBN10(test.param)
		if actual != test.expected {
			t.Errorf("Expected ISBN10(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// ISBN 13
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"foo", false},
		{"3-8362-2119-5", false},
		{"01234567890ab", false},
		{"978 3 8362 2119 0", false},
		{"9784873113685", true},
		{"978-4-87311-368-5", true},
		{"978 3401013190", true},
		{"978-3-8362-2119-1", true},
	}
	for _, test := range tests {
		actual := ISBN13(test.param)
		if actual != test.expected {
			t.Errorf("Expected ISBN13(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestDataURI(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"data:image/png;base64,TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"data:text/plain;base64,Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"image/gif;base64,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
		{"data:image/gif;base64,MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"data:image/png;base64,12345", false},
		{"", false},
		{"data:text,:;base85,U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", false},
	}
	for _, test := range tests {
		actual := DataURI(test.param)
		if actual != test.expected {
			t.Errorf("Expected DataURI(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestBase64(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"TG9yZW0gaXBzdW0gZG9sb3Igc2l0IGFtZXQsIGNvbnNlY3RldHVyIGFkaXBpc2NpbmcgZWxpdC4=", true},
		{"Vml2YW11cyBmZXJtZW50dW0gc2VtcGVyIHBvcnRhLg==", true},
		{"U3VzcGVuZGlzc2UgbGVjdHVzIGxlbw==", true},
		{"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAuMPNS1Ufof9EW/M98FNw" +
			"UAKrwflsqVxaxQjBQnHQmiI7Vac40t8x7pIb8gLGV6wL7sBTJiPovJ0V7y7oc0Ye" +
			"rhKh0Rm4skP2z/jHwwZICgGzBvA0rH8xlhUiTvcwDCJ0kc+fh35hNt8srZQM4619" +
			"FTgB66Xmp4EtVyhpQV+t02g6NzK72oZI0vnAvqhpkxLeLiMCyrI416wHm5Tkukhx" +
			"QmcL2a6hNOyu0ixX/x2kSFXApEnVrJ+/IxGyfyw8kf4N2IZpW5nEP847lpfj0SZZ" +
			"Fwrd1mnfnDbYohX2zRptLy2ZUn06Qo9pkG5ntvFEPo9bfZeULtjYzIl6K8gJ2uGZ" + "HQIDAQAB", true},
		{"12345", false},
		{"", false},
		{"Vml2YW11cyBmZXJtZtesting123", false},
	}
	for _, test := range tests {
		actual := Base64(test.param)
		if actual != test.expected {
			t.Errorf("Expected Base64(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestO3166Alpha2(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"ABCD", false},
		{"A", false},
		{"AC", false},
		{"AP", false},
		{"GER", false},
		{"NU", true},
		{"DE", true},
		{"JP", true},
		{"JPN", false},
		{"ZWE", false},
		{"GER", false},
		{"DEU", false},
	}
	for _, test := range tests {
		actual := ISO3166Alpha2(test.param)
		if actual != test.expected {
			t.Errorf("Expected ISO3166Alpha2(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestO3166Alpha3(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"ABCD", false},
		{"A", false},
		{"AC", false},
		{"AP", false},
		{"NU", false},
		{"DE", false},
		{"JP", false},
		{"ZWE", true},
		{"JPN", true},
		{"GER", false},
		{"DEU", true},
	}
	for _, test := range tests {
		actual := ISO3166Alpha3(test.param)
		if actual != test.expected {
			t.Errorf("Expected ISO3166Alpha3(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestIP(t *testing.T) {
	t.Parallel()

	// Without version
	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IP(test.param)
		if actual != test.expected {
			t.Errorf("Expected IP(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv4
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"1.2.3.4", true},
		{"::1", false},
		{"2001:db8:0000:1:1:1:1:1", false},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IPv4(test.param)
		if actual != test.expected {
			t.Errorf("Expected IPv4(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}

	// IPv6
	tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"127.0.0.1", false},
		{"0.0.0.0", false},
		{"255.255.255.255", false},
		{"1.2.3.4", false},
		{"::1", true},
		{"2001:db8:0000:1:1:1:1:1", true},
		{"300.0.0.0", false},
	}
	for _, test := range tests {
		actual := IPv6(test.param)
		if actual != test.expected {
			t.Errorf("Expected IPv6(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestPort(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"1", true},
		{"65535", true},
		{"0", false},
		{"65536", false},
		{"65538", false},
	}

	for _, test := range tests {
		actual := Port(test.param)
		if actual != test.expected {
			t.Errorf("Expected Port(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestDNSName(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"localhost", true},
		{"localhost.local", true},
		{"localhost.localdomain.intern", true},
		{"-localhost", false},
		{"localhost.-localdomain", false},
		{"localhost.localdomain.-int", false},
		{"_localhost", false},
		{"localhost._localdomain", false},
		{"localhost.localdomain._int", false},
		{"lÖcalhost", false},
		{"localhost.lÖcaldomain", false},
		{"localhost.localdomain.üntern", false},
		{"漢字汉字", false},
		{"www.jubfvq1v3p38i51622y0dvmdk1mymowjyeu26gbtw9andgynj1gg8z3msb1kl5z6906k846pj3sulm4kiyk82ln5teqj9nsht59opr0cs5ssltx78lfyvml19lfq1wp4usbl0o36cmiykch1vywbttcus1p9yu0669h8fj4ll7a6bmop505908s1m83q2ec2qr9nbvql2589adma3xsq2o38os2z3dmfh2tth4is4ixyfasasasefqwe4t2ub2fz1rme.de", false},
	}

	for _, test := range tests {
		actual := DNSName(test.param)
		if actual != test.expected {
			t.Errorf("Expected DNS(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestDialString(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"localhost.local:1", true},
		{"localhost.localdomain:9090", true},
		{"localhost.localdomain.intern:65535", true},
		{"127.0.0.1:30000", true},
		{"[::1]:80", true},
		{"[1200::AB00:1234::2552:7777:1313]:22", false},
		{"-localhost:1", false},
		{"localhost.-localdomain:9090", false},
		{"localhost.localdomain.-int:65535", false},
		{"localhost.loc:100000", false},
		{"漢字汉字:2", false},
		{"www.jubfvq1v3p38i51622y0dvmdk1mymowjyeu26gbtw9andgynj1gg8z3msb1kl5z6906k846pj3sulm4kiyk82ln5teqj9nsht59opr0cs5ssltx78lfyvml19lfq1wp4usbl0o36cmiykch1vywbttcus1p9yu0669h8fj4ll7a6bmop505908s1m83q2ec2qr9nbvql2589adma3xsq2o38os2z3dmfh2tth4is4ixyfasasasefqwe4t2ub2fz1rme.de:20000", false},
	}

	for _, test := range tests {
		actual := DialString(test.param)
		if actual != test.expected {
			t.Errorf("Expected DialString(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestMAC(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"3D:F2:C9:A6:B3:4F", true},
		{"3D-F2-C9-A6-B3:4F", false},
		{"123", false},
		{"", false},
		{"abacaba", false},
	}
	for _, test := range tests {
		actual := MAC(test.param)
		if actual != test.expected {
			t.Errorf("Expected MAC(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestFilePath(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
		osType   int
	}{
		{"c:\\" + strings.Repeat("a", 32767), true, Win}, //See http://msdn.microsoft.com/en-us/library/aa365247(VS.85).aspx#maxpath
		{"c:\\" + strings.Repeat("a", 32768), false, Win},
		{"c:\\path\\file (x86)\bar", true, Win},
		{"c:\\path\\file", true, Win},
		{"c:\\path\\file:exe", false, Unknown},
		{"C:\\", true, Win},
		{"c:\\path\\file\\", true, Win},
		{"c:/path/file/", false, Unknown},
		{"/path/file/", true, Unix},
		{"/path/file:SAMPLE/", true, Unix},
		{"/path/file:/.txt", true, Unix},
		{"/path", true, Unix},
	}
	for _, test := range tests {
		actual, osType := FilePath(test.param)
		if actual != test.expected || osType != test.osType {
			t.Errorf("Expected FilePath(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestLatitude(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"-90.000", true},
		{"+90", true},
		{"47.1231231", true},
		{"+99.9", false},
		{"108", false},
	}
	for _, test := range tests {
		actual := Latitude(test.param)
		if actual != test.expected {
			t.Errorf("Expected Latitude(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestLongitude(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"-180.000", true},
		{"180.1", false},
		{"+73.234", true},
		{"+382.3811", false},
		{"23.11111111", true},
	}
	for _, test := range tests {
		actual := Longitude(test.param)
		if actual != test.expected {
			t.Errorf("Expected Longitude(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestSSN(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"", false},
		{"00-90-8787", false},
		{"66690-76", false},
		{"191 60 2869", true},
		{"191-60-2869", true},
	}
	for _, test := range tests {
		actual := SSN(test.param)
		if actual != test.expected {
			t.Errorf("Expected SSN(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestMongoID(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    string
		expected bool
	}{
		{"507f1f77bcf86cd799439011", true},
		{"507f1f77bcf86cd7994390", false},
		{"507f1f77bcf86cd79943901z", false},
		{"507f1f77bcf86cd799439011 ", false},
		{"", false},
	}
	for _, test := range tests {
		actual := MongoID(test.param)
		if actual != test.expected {
			t.Errorf("Expected MongoID(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestSemver(t *testing.T) {
	t.Parallel()
	var tests = []struct {
		param    string
		expected bool
	}{
		{"v1.0.0", true},
		{"1.0.0", true},
		{"1.1.01", false},
		{"1.01.0", false},
		{"01.1.0", false},
		{"v1.1.01", false},
		{"v1.01.0", false},
		{"v01.1.0", false},
		{"1.0.0-alpha", true},
		{"1.0.0-alpha.1", true},
		{"1.0.0-0.3.7", true},
		{"1.0.0-0.03.7", false},
		{"1.0.0-00.3.7", false},
		{"1.0.0-x.7.z.92", true},
		{"1.0.0-alpha+001", true},
		{"1.0.0+20130313144700", true},
		{"1.0.0-beta+exp.sha.5114f85", true},
		{"1.0.0-beta+exp.sha.05114f85", true},
		{"1.0.0-+beta", false},
		{"1.0.0-b+-9+eta", false},
		{"v+1.8.0-b+-9+eta", false},
	}
	for _, test := range tests {
		actual := Semver(test.param)
		if actual != test.expected {
			t.Errorf("Expected Semver(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestNegative(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, false},
		{-1, true},
		{10, false},
		{3.14, false},
		{-96, true},
		{-10e-12, true},
	}
	for _, test := range tests {
		actual := Negative(test.param)
		if actual != test.expected {
			t.Errorf("Expected Negative(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestNonNegative(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, true},
		{-1, false},
		{10, true},
		{3.14, true},
		{-96, false},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := NonNegative(test.param)
		if actual != test.expected {
			t.Errorf("Expected NonNegative(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestPositive(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, false},
		{-1, false},
		{10, true},
		{3.14, true},
		{-96, false},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := Positive(test.param)
		if actual != test.expected {
			t.Errorf("Expected Positive(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestNonPositive(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, true},
		{-1, true},
		{10, false},
		{3.14, false},
		{-96, true},
		{-10e-12, true},
	}
	for _, test := range tests {
		actual := NonPositive(test.param)
		if actual != test.expected {
			t.Errorf("Expected NonPositive(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestWhole(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, true},
		{-1, true},
		{10, true},
		{3.14, false},
		{-96, true},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := Whole(test.param)
		if actual != test.expected {
			t.Errorf("Expected Whole(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestNatural(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		expected bool
	}{
		{0, false},
		{-1, false},
		{10, true},
		{3.14, false},
		{96, true},
		{-10e-12, false},
	}
	for _, test := range tests {
		actual := Natural(test.param)
		if actual != test.expected {
			t.Errorf("Expected Natural(%q) to be %v, got %v", test.param, test.expected, actual)
		}
	}
}

func TestInRange(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param    float64
		left     float64
		right    float64
		expected bool
	}{
		{0, 0, 0, true},
		{1, 0, 0, false},
		{-1, 0, 0, false},
		{0, -1, 1, true},
		{0, 0, 1, true},
		{0, -1, 0, true},
		{0, 0, -1, true},
		{0, 10, 5, false},
	}
	for _, test := range tests {
		actual := InRange(test.param, test.left, test.right)
		if actual != test.expected {
			t.Errorf("Expected InRange(%q, %q, %q) to be %v, got %v", test.param, test.left, test.right, test.expected, actual)
		}
	}
}

func TestMatches(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		param1   string
		param2   string
		expected bool
	}{
		{"123456789", "[0-9]+", true},
		{"abacada", "cab$", false},
		{"111222333", "((111|222|333)+)+", true},
		{"abacaba", "((123+]", false},
	}
	for _, test := range tests {
		actual := Matches(test.param1, test.param2)
		if actual != test.expected {
			t.Errorf("Expected Matches(%q,%q) to be %v, got %v", test.param1, test.param2, test.expected, actual)
		}
	}
}

func TestStringLength(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		value    string
		min      string
		max      string
		expected bool
	}{
		{"123456", "0", "100", true},
		{"1239999", "0", "0", false},
		{"1239asdfasf99", "100", "200", false},
		{"1239999asdff29", "10", "30", true},
		{"あいうえお", "0", "5", true},
		{"あいうえおか", "0", "5", false},
		{"あいうえお", "0", "0", false},
		{"あいうえ", "5", "10", false},
	}
	for _, test := range tests {
		actual := StringLength(test.value, test.min, test.max)
		if actual != test.expected {
			t.Errorf("Expected StringLength(%s, %s, %s) to be %v, got %v", test.value, test.min, test.max, test.expected, actual)
		}
	}
}
