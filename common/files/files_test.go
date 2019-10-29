package files

import (
	"github.com/pierdipi/go/common/assert"
	"testing"
)

func TestSameFiles(t *testing.T) {
	f1 := "testdata/f1.txt"
	f2 := "testdata/f1_clone.txt"
	f3 := "testdata/f2.txt"
	assertor := assert.New(t)

	isEqual, err := SameFiles(f1, f2)
	assertor.AssertThatError(err).IsNil()
	assertor.AssertThatBool(isEqual).IsEqual(true)

	isEqual, err = SameFiles(f1, f2, f3)
	assertor.AssertThatError(err).IsNil()
	assertor.AssertThatBool(isEqual).IsEqual(false)

	isEqual, err = SameFiles(f1, "_not_exists.txt")
	assertor.AssertThatError(err).IsNotNil()
	assertor.AssertThatBool(isEqual).IsEqual(false)

	isEqual, err = SameFiles("_not_exists.txt", f1)
	assertor.AssertThatError(err).IsNotNil()
	assertor.AssertThatBool(isEqual).IsEqual(false)

	isEqual, err = SameFiles("_not_exists.txt")
	assertor.AssertThatError(err).IsNil()
	assertor.AssertThatBool(isEqual).IsEqual(true)
}
