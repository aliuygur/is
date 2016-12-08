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
		benchmarkCase{Alpha, "Empty", ""},
		benchmarkCase{Alpha, "Space", " "},
		benchmarkCase{Alpha, "False", "abc!!!"},
		benchmarkCase{Alpha, "True", "abc"},
		//UTFLetter
		benchmarkCase{UTFLetter, "Empty", ""},
		benchmarkCase{UTFLetter, "Space", " "},
		benchmarkCase{UTFLetter, "False", "\u0026"}, //UTF-8(ASCII): &
		benchmarkCase{UTFLetter, "True", "abc"},
		// Alphanumeric
		benchmarkCase{Alphanumeric, "Empty", ""},
		benchmarkCase{Alphanumeric, "Space", " "},
		benchmarkCase{Alphanumeric, "False", "\u0026"}, //UTF-8(ASCII): &
		benchmarkCase{Alphanumeric, "True", "abc123"},
		// UTFLetterNumeric
		benchmarkCase{UTFLetterNumeric, "Empty", ""},
		benchmarkCase{UTFLetterNumeric, "Space", " "},
		benchmarkCase{UTFLetterNumeric, "False", "abc!!!"},
		benchmarkCase{UTFLetterNumeric, "True", "abc1"},
		// Numeric
		benchmarkCase{Numeric, "Empty", ""},
		benchmarkCase{Numeric, "Space", " "},
		benchmarkCase{Numeric, "False", "abc!!!"},
		benchmarkCase{Numeric, "True", "0123"},
		// UTFNumeric
		benchmarkCase{UTFNumeric, "Empty", ""},
		benchmarkCase{UTFNumeric, "Space", " "},
		benchmarkCase{UTFNumeric, "False", "abc!!!"},
		benchmarkCase{UTFNumeric, "True", "\u0030"}, //UTF-8(ASCII): 0
		// UTFDigit
		benchmarkCase{UTFDigit, "Empty", ""},
		benchmarkCase{UTFDigit, "Space", " "},
		benchmarkCase{UTFDigit, "False", "abc!!!"},
		benchmarkCase{UTFDigit, "True", "\u0030"}, //UTF-8(ASCII): 0
		// LowerCase
		benchmarkCase{LowerCase, "Empty", ""},
		benchmarkCase{LowerCase, "Space", " "},
		benchmarkCase{LowerCase, "False", "ABC"},
		benchmarkCase{LowerCase, "True", "abc"},
		// UpperCase
		benchmarkCase{UpperCase, "Empty", ""},
		benchmarkCase{UpperCase, "Space", " "},
		benchmarkCase{UpperCase, "False", "abc"},
		benchmarkCase{UpperCase, "True", "ABC"},
		// Int
		benchmarkCase{Int, "Empty", ""},
		benchmarkCase{Int, "Space", " "},
		benchmarkCase{Int, "False", "abc"},
		benchmarkCase{Int, "True", "000"},
		// Email
		benchmarkCase{Email, "Empty", ""},
		benchmarkCase{Email, "Space", " "},
		benchmarkCase{Email, "False", "@invalid.com"},
		benchmarkCase{Email, "True", "foo@bar.com"},
		// URL
		benchmarkCase{URL, "Empty", ""},
		benchmarkCase{URL, "Space", " "},
		benchmarkCase{URL, "False", "./rel/test/dir"},
		benchmarkCase{URL, "True", "http://foobar.org/"},
		// RequestURL
		benchmarkCase{RequestURL, "Empty", ""},
		benchmarkCase{RequestURL, "Space", " "},
		benchmarkCase{RequestURL, "False", "invalid."},
		benchmarkCase{RequestURL, "True", "http://foobar.org/"},
		// RequestURI
		benchmarkCase{RequestURI, "Empty", ""},
		benchmarkCase{RequestURI, "Space", " "},
		benchmarkCase{RequestURI, "False", "invalid."},
		benchmarkCase{RequestURI, "True", "http://foobar.org/"},
		// Float
		benchmarkCase{Float, "Empty", ""},
		benchmarkCase{Float, "Space", " "},
		benchmarkCase{Float, "False", "+1f"},
		benchmarkCase{Float, "True", "123.123"},
		// Hexadecimal
		benchmarkCase{Hexadecimal, "Empty", ""},
		benchmarkCase{Hexadecimal, "Space", " "},
		benchmarkCase{Hexadecimal, "False", ".."},
		benchmarkCase{Hexadecimal, "True", "deadBEEF"},
		// Hexcolor
		benchmarkCase{Hexcolor, "Empty", ""},
		benchmarkCase{Hexcolor, "Space", " "},
		benchmarkCase{Hexcolor, "False", "#ff12FG"},
		benchmarkCase{Hexcolor, "True", "#f00"},
		// RGBcolor
		benchmarkCase{RGBcolor, "Empty", ""},
		benchmarkCase{RGBcolor, "Space", " "},
		benchmarkCase{RGBcolor, "False", "rgba(0,31,255)"},
		benchmarkCase{RGBcolor, "True", "rgb(0,31,255)"},
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
