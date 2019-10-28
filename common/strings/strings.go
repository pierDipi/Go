// Package strings provide additional functions to manipulate
// utf-8 encoded strings.
// Source code: https://https://github.com/pierDipi/go/tree/master/common/strings
package strings // import "github.com/pierdipi/go/common/strings"

import (
	stdstrings "strings"
	"unicode"
	"unicode/utf8"
)

// CommonPrefix returns the longest string prefix such that
// `a` and `b` start with prefix.
// If `a` and `b` have no common prefix returns an empty string.
func CommonPrefix(a, b string) string {
	if AnyIsEmpty(a, b) {
		return ""
	}
	endIndex := 0
	for i, w := 0, 0; i < len(a); i += w {
		if i >= len(b) {
			break
		}
		runeValueA, width := utf8.DecodeRuneInString(a[i:])
		runeValueB, _ := utf8.DecodeRuneInString(b[i:])
		if runeValueA == runeValueB {
			endIndex = i + width
		} else {
			break
		}
		w = width
	}
	return a[:endIndex]
}

// CommonSuffix returns the longest string suffix such that
// `a` and `b` end with suffix.
// If `a` and `b` have no common suffix returns an empty string.
func CommonSuffix(a, b string) string {
	if AnyIsEmpty(a, b) {
		return ""
	}
	longest := Longest(a, b)
	shortest := Shortest(a, b)
	diff := len(*longest) - len(*shortest)
	i := len(*longest)
	startIndex := i
	for w := 1; i > 0 && w > 0; i -= w {
		if i - diff <= 0 {
			break
		}
		runeValueA, width := utf8.DecodeLastRuneInString((*longest)[:i])
		runeValueB, _ := utf8.DecodeLastRuneInString((*shortest)[:i - diff])
		if runeValueA == runeValueB {
			startIndex = i - width
		} else {
			break
		}
		w = width
	}
	return (*longest)[startIndex:]
}

// Longest returns the pointer to the string `a` if the len(a) >= len(b).
func Longest(a, b string) *string  {
	if len(a) >= len(b) {
		return &a
	}
	return &b
}

// Shortest returns the pointer to the string `a` if the len(a) < len(b).
func Shortest(a, b string) *string {
	if len(a) < len(b) {
		return &a
	}
	return &b
}

// AnyIsEmpty returns true if at least one string is an empty string.
func AnyIsEmpty(s ... string) bool {
	for i := range s {
		if IsEmpty(s[i]) {
			return true
		}
	}
	return false
}

// IsEmpty returns true if the given string is empty.
func IsEmpty(s string) bool {
	return s == ""
}

// IsNotEmpty returns true if the given string is not empty.
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsNilOrEmpty returns true if the given string pointer is `nil` or empty.
func IsNilOrEmpty(s *string) bool {
	return s == nil || IsEmpty(*s)
}

// IsNotNilOrEmpty returns true if the given pointer is not `nil` and not empty.
func IsNotNilOrEmpty(s *string) bool {
	return !IsNilOrEmpty(s)
}

// PaddingPrefix returns a string, of length at least `minLength`,
// consisting of a string prefixed with as many copies of `character`
// as necessary to reach that length.
func PaddingPrefix(s string, minLength uint, character rune) string {
	return padding(s, minLength, string(character)) + s
}

// PaddingSuffix returns a string, of length at least `minLength`,
// consisting of a string suffixed with as many copies of `character`
// as necessary to reach that length.
func PaddingSuffix(s string, minLength uint, character rune) string {
	return s + padding(s, minLength, string(character))
}

func padding(s string, minLength uint, padding string) string {
	if len(s) >= int(minLength) {
		return ""
	}
	return stdstrings.Repeat(padding, int(minLength) - len(s))
}

// EqualsIgnoreCase returns true if the given strings are equals
// ignoring the characters case.
func EqualsIgnoreCase(a, b string) bool {
	return stdstrings.EqualFold(a, b)
}

// IsLowerCase returns true if the given string
// contains only lower case characters.
func IsLowerCase(a string) bool {
	runes := []rune(a)
	for i := range runes {
		if !unicode.IsLower(runes[i]) {
			return false
		}
	}
	return true
}

// IsUpperCase returns true if the given string
// contains only upper case characters.
func IsUpperCase(a string) bool {
	if a == "" {
		return true
	}
	return !IsLowerCase(a)
}

// TruncateSuffix truncates the given string to the given maxLength.
// If the length of the string is greater than maxLength the returned string
// will be exactly maxLength long and will end with the given truncationIndicator.
// Otherwise, the string will be returned as a string with no changes to the content of s.
// If maxLength < len(truncationIndicator) panics.
func TruncateSuffix(s string, maxLength int, truncationIndicator string) string {
	if maxLength < len(truncationIndicator) {
		panic("maxLength must be greater than the length of truncationIndicator")
	}
	if len(s) <= maxLength {
		return s
	}
	return s[0: maxLength - len(truncationIndicator)] + truncationIndicator
}

// TruncatePrefix truncates the given string to the given maxLength.
// If the length of the string is greater than maxLength the returned string
// will be exactly maxLength long and will start with the given truncationIndicator.
// Otherwise, the string will be returned as a string with no changes to the content of s.
// If maxLength < len(truncationIndicator) panics.
func TruncatePrefix(s string, maxLength int, truncationIndicator string) string {
	if maxLength < len(truncationIndicator) {
		panic("maxLength must be greater than the length of truncationIndicator")
	}
	if len(s) <= maxLength {
		return s
	}
	return truncationIndicator + s[len(s) - maxLength + len(truncationIndicator):]
}

