// +build go1.7

package is

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

type benchmarkCase struct {
	function func(string) bool
	caseName string
	param    string
}

var (
	benchmarkFuncs = []benchmarkCase{
		// Alpha
		{Alpha, "Empty", ""},
		{Alpha, "Space", " "},
		{Alpha, "False", "abc!!!"},
		{Alpha, "True", "abc"},
		//UTFLetter
		{UTFLetter, "Empty", ""},
		{UTFLetter, "Space", " "},
		{UTFLetter, "False", "\u0026"}, //UTF-8(ASCII): &
		{UTFLetter, "True", "abc"},
		// Alphanumeric
		{Alphanumeric, "Empty", ""},
		{Alphanumeric, "Space", " "},
		{Alphanumeric, "False", "\u0026"}, //UTF-8(ASCII): &
		{Alphanumeric, "True", "abc123"},
		// UTFLetterNumeric
		{UTFLetterNumeric, "Empty", ""},
		{UTFLetterNumeric, "Space", " "},
		{UTFLetterNumeric, "False", "abc!!!"},
		{UTFLetterNumeric, "True", "abc1"},
		// Numeric
		{Numeric, "Empty", ""},
		{Numeric, "Space", " "},
		{Numeric, "False", "abc!!!"},
		{Numeric, "True", "0123"},
		// UTFNumeric
		{UTFNumeric, "Empty", ""},
		{UTFNumeric, "Space", " "},
		{UTFNumeric, "False", "abc!!!"},
		{UTFNumeric, "True", "\u0030"}, //UTF-8(ASCII): 0
		// UTFDigit
		{UTFDigit, "Empty", ""},
		{UTFDigit, "Space", " "},
		{UTFDigit, "False", "abc!!!"},
		{UTFDigit, "True", "\u0030"}, //UTF-8(ASCII): 0
		// LowerCase
		{LowerCase, "Empty", ""},
		{LowerCase, "Space", " "},
		{LowerCase, "False", "ABC"},
		{LowerCase, "True", "abc"},
		// UpperCase
		{UpperCase, "Empty", ""},
		{UpperCase, "Space", " "},
		{UpperCase, "False", "abc"},
		{UpperCase, "True", "ABC"},
		// Int
		{Int, "Empty", ""},
		{Int, "Space", " "},
		{Int, "False", "abc"},
		{Int, "True", "000"},
		// Email
		{Email, "Empty", ""},
		{Email, "Space", " "},
		{Email, "False", "@invalid.com"},
		{Email, "True", "foo@bar.com"},
		// URL
		{URL, "Empty", ""},
		{URL, "Space", " "},
		{URL, "False", "./rel/test/dir"},
		{URL, "True", "http://foobar.org/"},
		// RequestURL
		{RequestURL, "Empty", ""},
		{RequestURL, "Space", " "},
		{RequestURL, "False", "invalid."},
		{RequestURL, "True", "http://foobar.org/"},
		// RequestURI
		{RequestURI, "Empty", ""},
		{RequestURI, "Space", " "},
		{RequestURI, "False", "invalid."},
		{RequestURI, "True", "http://foobar.org/"},
		// Float
		{Float, "Empty", ""},
		{Float, "Space", " "},
		{Float, "False", "+1f"},
		{Float, "True", "123.123"},
		// Hexadecimal
		{Hexadecimal, "Empty", ""},
		{Hexadecimal, "Space", " "},
		{Hexadecimal, "False", ".."},
		{Hexadecimal, "True", "deadBEEF"},
		// Hexcolor
		{Hexcolor, "Empty", ""},
		{Hexcolor, "Space", " "},
		{Hexcolor, "False", "#ff12FG"},
		{Hexcolor, "True", "#f00"},
		// RGBcolor
		{RGBcolor, "Empty", ""},
		{RGBcolor, "Space", " "},
		{RGBcolor, "False", "rgba(0,31,255)"},
		{RGBcolor, "True", "rgb(0,31,255)"},
	}
)

func getFuncName(f interface{}) string {
	n := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	return filepath.Base(n)
}

func Benchmark(b *testing.B) {
	for _, bencCase := range benchmarkFuncs {
		name := getFuncName(bencCase.function) + "." + bencCase.caseName
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bencCase.function(bencCase.param)
			}
		})
	}
}
