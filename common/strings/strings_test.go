package strings

import (
	"github.com/pierdipi/go/common/assert"
	"testing"
)

func TestCommonPrefixString(t *testing.T) {
	tt := []struct {
		a, b, expected string;
	}{
		{a: "", b: "", expected: ""},
		{a: "a", b: "", expected: ""},
		{a: "a", b: "a", expected: "a"},
		{a: "abc", b: "a", expected: "a"},
		{a: "ab", b: "abcd", expected: "ab"},
		{a: "ac", b: "ab", expected: "a"},
		{a: "òè+èà", b: "òè+èà", expected: "òè+èà"},
		{a: "传/傳", b: "传/傳", expected: "传/傳"},
		{a: "传/傳", b: "传/", expected: "传/"},
	}
	assertor := assert.New(t)
	var commonPrefix string
	for r := range tt {
		commonPrefix = CommonPrefix(tt[r].a, tt[r].b)
		assertor.AssertThatString(commonPrefix).IsEqual(tt[r].expected)
	}
}

func TestCommonSuffixString(t *testing.T) {
	tt := []struct{
		a, b, expected string
	}{
		{a: "", b: "", expected: ""},
		{a: "a", b: "", expected: ""},
		{a: "a", b: "a", expected: "a"},
		{a: "abc", b: "c", expected: "c"},
		{a: "cd", b: "abcd", expected: "cd"},
		{a: "ca", b: "ba", expected: "a"},
		{a: "òè+èà", b: "òè+èà", expected: "òè+èà"},
		{a: "传/傳", b: "传/傳", expected: "传/傳"},
		{a: "传/傳", b: "傳", expected: "傳"},
		{a: "abcdefgh", b: "bcdefgh", expected: "bcdefgh"},
		{a: "傳", b: "abcdgrgf传/傳", expected: "傳"},
	}
	assertor := assert.New(t)
	var commonSuffix string
	for r := range tt {
		commonSuffix = CommonSuffix(tt[r].a, tt[r].b)
		assertor.AssertThatString(commonSuffix).IsEqual(tt[r].expected)
	}
}

func TestIsNotEmpty(t *testing.T) {
	tt := []struct{
		s string
		expected bool
	}{
		{s: "", expected: false},
		{s: "a", expected: true},
		{s: "ò", expected: true},
	}
	assertor := assert.New(t)
	for r := range tt {
		assertor.AssertThatBool(IsNotEmpty(tt[r].s)).IsEqual(tt[r].expected)
	}
}

func TestIsNotNilOrEmpty(t *testing.T) {

	getRef := func(s string) *string { return &s }

	tt := []struct{
		s *string
		expected bool
	}{
		{s: nil, expected: false},
		{s: getRef(""), expected: false},
		{s: getRef("a"), expected: true},
		{s: getRef("ò"), expected: true},
	}
	assertor := assert.New(t)
	for r := range tt {
		assertor.AssertThatBool(IsNotNilOrEmpty(tt[r].s)).IsEqual(tt[r].expected)
	}
}

func TestPaddingPrefix(t *testing.T) {
	tt := []struct{
		s string
		minLength uint
		char rune
		expected string
	}{
		{s: "9", minLength: 3, char: '0', expected: "009"},
		{s: "90", minLength: 3, char: '0', expected: "090"},
		{s: "", minLength: 1, char: '0', expected: "0"},
		{s: "", minLength: 0, char: '0', expected: ""},
	}
	assertor := assert.New(t)
	var stringWithPadding string
	for r := range tt {
		stringWithPadding = PaddingPrefix(tt[r].s, tt[r].minLength, tt[r].char)
		assertor.AssertThatString(stringWithPadding).IsEqual(tt[r].expected)
	}
}

func TestPaddingSuffix(t *testing.T) {
	tt := []struct{
		s string
		minLength uint
		char rune
		expected string
	}{
		{s: "9", minLength: 3, char: '0', expected: "900"},
		{s: "90", minLength: 3, char: '0', expected: "900"},
		{s: "", minLength: 1, char: '0', expected: "0"},
		{s: "", minLength: 0, char: '0', expected: ""},
	}
	assertor := assert.New(t)
	var stringWithPadding string
	for r := range tt {
		stringWithPadding = PaddingSuffix(tt[r].s, tt[r].minLength, tt[r].char)
		assertor.AssertThatString(stringWithPadding).IsEqual(tt[r].expected)
	}
}

func TestIsUpperCase(t *testing.T) {
	tt := []struct{
		a string
		expected bool
	}{
		{a: "ABC", expected: true},
		{a: "a", expected: false},
		{a: "", expected: true},
	}
	assertor := assert.New(t)
	for r := range tt {
		assertor.AssertThatBool(IsUpperCase(tt[r].a)).IsEqual(tt[r].expected)
	}
}

func TestEqualsIgnoreCase(t *testing.T) {
	tt := []struct{
		a, b string
		expected bool
	}{
		{a: "ABC", b: "abc", expected: true},
		{a: "a", b: "a", expected: true},
		{a: "", b: "",expected: true},
		{a: "", b: "a",expected: false},
	}
	assertor := assert.New(t)
	for r := range tt {
		assertor.AssertThatBool(EqualsIgnoreCase(tt[r].a, tt[r].b)).IsEqual(tt[r].expected)
	}
}

func TestTruncateSuffix(t *testing.T) {
	tt := []struct{
		s string
		maxLength int
		truncationIndicator string
		expected string
		shouldPanic bool
	}{
		{s: "foo", maxLength: 3, truncationIndicator: "***", expected: "foo"},
		{s: "foofoo", maxLength: 5, truncationIndicator: "***", expected: "fo***"},
		{s: "mail@email.com", maxLength: 10, truncationIndicator: "*****", expected: "mail@*****"},
		{s: "happy", maxLength: 4, truncationIndicator: "***", expected: "h***"},
		{s: "happy", maxLength: 4, truncationIndicator: "*****", expected: "", shouldPanic: true},
	}
	assertor := assert.New(t)
	var truncated string
	for r := range tt {
		if tt[r].shouldPanic {
			f := func() { TruncateSuffix(tt[r].s, tt[r].maxLength, tt[r].truncationIndicator) }
			assertor.AssertThatFunction(f).Panics()
		} else {
			truncated = TruncateSuffix(tt[r].s, tt[r].maxLength, tt[r].truncationIndicator)
			assertor.AssertThatString(truncated).IsEqual(tt[r].expected)
		}
	}
}

func TestTruncatePrefix(t *testing.T) {
	tt := []struct{
		s string
		maxLength int
		truncationIndicator string
		expected string
		shouldPanic bool
	}{
		{s: "foo", maxLength: 3, truncationIndicator: "***", expected: "foo"},
		{s: "foofoo", maxLength: 5, truncationIndicator: "***", expected: "***oo"},
		{s: "mail@email.com", maxLength: 10, truncationIndicator: "****", expected: "****il.com"},
		{s: "happy", maxLength: 4, truncationIndicator: "***", expected: "***y"},
		{s: "happy", maxLength: 4, truncationIndicator: "*****", expected: "", shouldPanic: true},
	}
	assertor := assert.New(t)
	var truncated string
	for r := range tt {
		if tt[r].shouldPanic {
			f := func() { TruncatePrefix(tt[r].s, tt[r].maxLength, tt[r].truncationIndicator) }
			assertor.AssertThatFunction(f).Panics()
		} else {
			truncated = TruncatePrefix(tt[r].s, tt[r].maxLength, tt[r].truncationIndicator)
			assertor.AssertThatString(truncated).IsEqual(tt[r].expected)
		}
	}
}