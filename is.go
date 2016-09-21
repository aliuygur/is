package is

import (
	"encoding/base64"
	"encoding/json"
	"math"
	"net"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// InRange returns true if value lies between left and right border
func InRange(value, left, right float64) bool {
	if left > right {
		left, right = right, left
	}
	return value >= left && value <= right
}

// Email is a constraint to do a simple validation for email addresses, it only check if the string contains "@"
// and that it is not in the first or last character of the string
// https://en.wikipedia.org/wiki/Email_address#Valid_email_addresses
func Email(s string) bool {
	if !strings.Contains(s, "@") || string(s[0]) == "@" || string(s[len(s)-1]) == "@" {
		return false
	}
	return true
}

// URL check if the string is an URL.
func URL(str string) bool {
	if str == "" || len(str) >= 2083 || len(str) <= 3 || strings.HasPrefix(str, ".") {
		return false
	}
	u, err := url.Parse(str)
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	return rxURL.MatchString(str)

}

// RequestURL check if the string rawurl, assuming
// it was recieved in an HTTP request, is a valid
// URL confirm to RFC 3986
func RequestURL(rawurl string) bool {
	url, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return false //Couldn't even parse the rawurl
	}
	if len(url.Scheme) == 0 {
		return false //No Scheme found
	}
	return true
}

// RequestURI check if the string rawurl, assuming
// it was recieved in an HTTP request, is an
// absolute URI or an absolute path.
func RequestURI(rawurl string) bool {
	_, err := url.ParseRequestURI(rawurl)
	return err == nil
}

// Alpha check if the string contains only letters (a-zA-Z). Empty string is valid.
func Alpha(s string) bool {
	for _, v := range s {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') {
			return false
		}
	}
	return true
}

//IsUTFLetter check if the string contains only unicode letter characters.
//Similar to IsAlpha but for all languages. Empty string is valid.
func UTFLetter(str string) bool {
	for _, v := range str {
		if !unicode.IsLetter(v) {
			return false
		}
	}
	return true

}

// Alphanumeric check if the string contains only letters and numbers. Empty string is valid.
func Alphanumeric(s string) bool {
	for _, v := range s {
		if ('Z' < v || v < 'A') && ('z' < v || v < 'a') && ('9' < v || v < '0') {
			return false
		}
	}
	return true
}

// UTFLetterNumeric check if the string contains only unicode letters and numbers. Empty string is valid.
func UTFLetterNumeric(s string) bool {
	for _, v := range s {
		if !unicode.IsLetter(v) && !unicode.IsNumber(v) { //letters && numbers are ok
			return false
		}
	}
	return true
}

// Numeric check if the string contains only numbers. Empty string is valid.
func Numeric(s string) bool {
	for _, v := range s {
		if '9' < v || v < '0' {
			return false
		}
	}
	return true
}

// UTFNumeric check if the string contains only unicode numbers of any kind.
// Numbers can be 0-9 but also Fractions ¾,Roman Ⅸ and Hangzhou 〩. Empty string is valid.
func UTFNumeric(s string) bool {
	for _, v := range s {
		if unicode.IsNumber(v) == false {
			return false
		}
	}
	return true
}

// Negative returns true if value < 0
func Negative(value float64) bool {
	return value < 0
}

// Positive returns true if value > 0
func Positive(value float64) bool {
	return value > 0
}

// NonNegative returns true if value >= 0
func NonNegative(value float64) bool {
	return value >= 0
}

// NonPositive returns true if value <= 0
func NonPositive(value float64) bool {
	return value <= 0
}

// Whole returns true if value is whole number
func Whole(value float64) bool {
	return math.Abs(math.Remainder(value, 1)) == 0
}

// Natural returns true if value is natural number (positive and whole)
func Natural(value float64) bool {
	return Whole(value) && Positive(value)
}

// UTFDigit check if the string contains only unicode radix-10 decimal digits. Empty string is valid.
func UTFDigit(s string) bool {
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

// Hexadecimal check if the string is a hexadecimal number.
func Hexadecimal(str string) bool {
	_, err := strconv.ParseInt(str, 16, 0)
	return err == nil
}

// Hexcolor check if the string is a hexadecimal color.
func Hexcolor(str string) bool {
	return rxHexcolor.MatchString(str)
}

// RGBcolor check if the string is a valid RGB color in form rgb(RRR, GGG, BBB).
func RGBcolor(str string) bool {
	return rxRGBcolor.MatchString(str)
}

// LowerCase check if the string is lowercase. Empty string is valid.
func LowerCase(str string) bool {
	if NullString(str) {
		return true
	}
	return str == strings.ToLower(str)
}

// UpperCase check if the string is uppercase. Empty string is valid.
func UpperCase(str string) bool {
	if NullString(str) {
		return true
	}
	return str == strings.ToUpper(str)
}

// Int check if the string is an integer. Empty string is valid.
func Int(str string) bool {
	if NullString(str) {
		return true
	}
	_, err := strconv.Atoi(str)

	return err == nil
}

// Float check if the string is a float.
func Float(str string) bool {
	_, err := strconv.ParseFloat(str, 0)
	return err == nil
}

// DivisibleBy check if the string is a number that's divisible by another.
// If second argument is not valid integer or zero, it's return false.
// Otherwise, if first argument is not valid integer or zero, it's return true (Invalid string converts to zero).
func DivisibleBy(str, num string) bool {
	f, _ := toFloat(str)
	p := int64(f)
	q, _ := toInt(num)
	if q == 0 {
		return false
	}
	return (p == 0) || (p%q == 0)
}

// NullString check if the string is null.
func NullString(str string) bool {
	return len(str) == 0
}

// Null check if the variable is null.
func Null(obj interface{}) bool {
	if obj != nil {
		return false
	}
	return true
}

// ByteLength check if the string's length (in bytes) falls in a range.
func ByteLength(str string, min, max int) bool {
	return len(str) >= min && len(str) <= max
}

// UUIDv3 check if the string is a UUID version 3.
func UUIDv3(str string) bool {
	return rxUUID3.MatchString(str)
}

// UUIDv4 check if the string is a UUID version 4.
func UUIDv4(str string) bool {
	return rxUUID4.MatchString(str)
}

// UUIDv5 check if the string is a UUID version 5.
func UUIDv5(str string) bool {
	return rxUUID5.MatchString(str)
}

// UUID check if the string is a UUID (version 3, 4 or 5).
func UUID(str string) bool {
	return rxUUID.MatchString(str)
}

// CreditCard check if the string is a credit card.
func CreditCard(str string) bool {
	r, _ := regexp.Compile("[^0-9]+")
	sanitized := r.ReplaceAll([]byte(str), []byte(""))
	if !rxCreditCard.MatchString(string(sanitized)) {
		return false
	}
	var sum int64
	var digit string
	var tmpNum int64
	var shouldDouble bool
	for i := len(sanitized) - 1; i >= 0; i-- {
		digit = string(sanitized[i:(i + 1)])
		tmpNum, _ = toInt(digit)
		if shouldDouble {
			tmpNum *= 2
			if tmpNum >= 10 {
				sum += ((tmpNum % 10) + 1)
			} else {
				sum += tmpNum
			}
		} else {
			sum += tmpNum
		}
		shouldDouble = !shouldDouble
	}

	if sum%10 == 0 {
		return true
	}
	return false
}

// ISBN10 check if the string is an ISBN version 10.
func ISBN10(str string) bool {
	return ISBN(str, 10)
}

// ISBN13 check if the string is an ISBN version 13.
func ISBN13(str string) bool {
	return ISBN(str, 13)
}

// ISBN check if the string is an ISBN (version 10 or 13).
// If version value is not equal to 10 or 13, it will be check both variants.
func ISBN(str string, version int) bool {
	r, _ := regexp.Compile("[\\s-]+")
	sanitized := r.ReplaceAll([]byte(str), []byte(""))
	var checksum int32
	var i int32
	if version == 10 {
		if !rxISBN10.MatchString(string(sanitized)) {
			return false
		}
		for i = 0; i < 9; i++ {
			checksum += (i + 1) * int32(sanitized[i]-'0')
		}
		if sanitized[9] == 'X' {
			checksum += 10 * 10
		} else {
			checksum += 10 * int32(sanitized[9]-'0')
		}
		if checksum%11 == 0 {
			return true
		}
		return false
	} else if version == 13 {
		if !rxISBN13.MatchString(string(sanitized)) {
			return false
		}
		factor := []int32{1, 3}
		for i = 0; i < 12; i++ {
			checksum += factor[i%2] * int32(sanitized[i]-'0')
		}
		if (int32(sanitized[12]-'0'))-((10-(checksum%10))%10) == 0 {
			return true
		}
		return false
	}
	return ISBN(str, 10) || ISBN(str, 13)
}

// JSON check if the string is valid JSON (note: uses json.Unmarshal).
func JSON(str string) bool {
	var js json.RawMessage
	return json.Unmarshal([]byte(str), &js) == nil
}

// Multibyte check if the string contains one or more multibyte chars. Empty string is valid.
func Multibyte(s string) bool {
	for _, v := range s {
		if v >= utf8.RuneSelf {
			return true
		}
	}

	return NullString(s)
}

// ASCII check if the string contains ASCII chars only. Empty string is valid.
func ASCII(s string) bool {
	for _, v := range s {
		if v >= utf8.RuneSelf {
			return false
		}
	}
	return true
}

// PrintableASCII check if the string contains printable ASCII chars only. Empty string is valid.
func PrintableASCII(s string) bool {
	for _, v := range s {
		if v < ' ' || v > '~' {
			return false
		}
	}
	return true
}

// FullWidth check if the string contains any full-width chars. Empty string is valid.
func FullWidth(str string) bool {
	if NullString(str) {
		return true
	}
	return rxFullWidth.MatchString(str)
}

// HalfWidth check if the string contains any half-width chars. Empty string is valid.
func HalfWidth(str string) bool {
	if NullString(str) {
		return true
	}
	return rxHalfWidth.MatchString(str)
}

// VariableWidth check if the string contains a mixture of full and half-width chars. Empty string is valid.
func VariableWidth(str string) bool {
	if NullString(str) {
		return true
	}
	return rxHalfWidth.MatchString(str) && rxFullWidth.MatchString(str)
}

// Base64 check if a string is base64 encoded.
func Base64(s string) bool {
	if NullString(s) {
		return false
	}
	_, err := base64.StdEncoding.DecodeString(s)

	return err == nil
}

// FilePath check is a string is Win or Unix file path and returns it's type.
func FilePath(str string) (bool, int) {
	if rxWinPath.MatchString(str) {
		//check windows path limit see:
		//  http://msdn.microsoft.com/en-us/library/aa365247(VS.85).aspx#maxpath
		if len(str[3:]) > 32767 {
			return false, Win
		}
		return true, Win
	} else if rxUnixPath.MatchString(str) {
		return true, Unix
	}
	return false, Unknown
}

// DataURI checks if a string is base64 encoded data URI such as an image
func DataURI(str string) bool {
	dataURI := strings.Split(str, ",")
	if !rxDataURI.MatchString(dataURI[0]) {
		return false
	}
	return Base64(dataURI[1])
}

// ISO3166Alpha2 checks if a string is valid two-letter country code
func ISO3166Alpha2(str string) bool {
	for _, entry := range ISO3166List {
		if str == entry.Alpha2Code {
			return true
		}
	}
	return false
}

// ISO3166Alpha3 checks if a string is valid three-letter country code
func ISO3166Alpha3(str string) bool {
	for _, entry := range ISO3166List {
		if str == entry.Alpha3Code {
			return true
		}
	}
	return false
}

// DNSName will validate the given string as a DNS name
func DNSName(str string) bool {
	if str == "" || len(strings.Replace(str, ".", "", -1)) > 255 {
		// constraints already violated
		return false
	}
	return rxDNSName.MatchString(str)
}

// DialString validates the given string for usage with the various Dial() functions
func DialString(str string) bool {

	if h, p, err := net.SplitHostPort(str); err == nil && h != "" && p != "" && (DNSName(h) || IP(h)) && Port(p) {
		return true
	}

	return false
}

// IP checks if a string is either IP version 4 or 6.
func IP(str string) bool {
	return net.ParseIP(str) != nil
}

// Port checks if a string represents a valid port
func Port(str string) bool {
	if i, err := strconv.Atoi(str); err == nil && i > 0 && i < 65536 {
		return true
	}
	return false
}

// IPv4 check if the string is an IP version 4.
func IPv4(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ".")
}

// IPv6 check if the string is an IP version 6.
func IPv6(str string) bool {
	ip := net.ParseIP(str)
	return ip != nil && strings.Contains(str, ":")
}

// MAC check if a string is valid MAC address.
// Possible MAC formats:
// 01:23:45:67:89:ab
// 01:23:45:67:89:ab:cd:ef
// 01-23-45-67-89-ab
// 01-23-45-67-89-ab-cd-ef
// 0123.4567.89ab
// 0123.4567.89ab.cdef
func MAC(str string) bool {
	_, err := net.ParseMAC(str)
	return err == nil
}

// MongoID check if the string is a valid hex-encoded representation of a MongoDB ObjectId.
func MongoID(str string) bool {
	return rxMongoID.MatchString(str)
}

// Latitude check if a string is valid latitude.
func Latitude(str string) bool {
	return rxLatitude.MatchString(str)
}

// Longitude check if a string is valid longitude.
func Longitude(str string) bool {
	return rxLongitude.MatchString(str)
}

// SSN will validate the given string as a U.S. Social Security Number
func SSN(str string) bool {
	if str == "" || len(str) != 11 {
		return false
	}
	return rxSSN.MatchString(str)
}

// Semver check if string is valid semantic version
func Semver(str string) bool {
	return rxSemver.MatchString(str)
}

// Matches check if string matches the pattern (pattern is regular expression)
// In case of error return false
func Matches(str, pattern string) bool {
	match, _ := regexp.MatchString(pattern, str)
	return match
}

// StringMatches checks if a string matches a given pattern.
func StringMatches(s string, params ...string) bool {
	if len(params) == 1 {
		pattern := params[0]
		return Matches(s, pattern)
	}
	return false
}

// StringLength check string's length (including multi byte strings)
func StringLength(str string, params ...string) bool {

	if len(params) == 2 {
		strLength := utf8.RuneCountInString(str)
		min, _ := toInt(params[0])
		max, _ := toInt(params[1])
		return strLength >= int(min) && strLength <= int(max)
	}

	return false
}

//Exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// toFloat convert the input string to a float, or 0.0 if the input is not a float.
func toFloat(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		res = 0.0
	}
	return res, err
}

// toInt convert the input string to an integer, or 0 if the input is not an integer.
func toInt(str string) (int64, error) {
	res, err := strconv.ParseInt(str, 0, 64)
	if err != nil {
		res = 0
	}
	return res, err
}
