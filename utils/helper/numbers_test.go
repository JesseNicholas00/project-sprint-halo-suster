package helper_test

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/utils/helper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHasLen(t *testing.T) {
	Convey("When given a number of the correct length", t, func() {
		number := int64(123456789)
		Convey("Should return true", func() {
			So(helper.HasLen(number, 9), ShouldBeTrue)
		})
	})
	Convey("When given a number of an incorrect length", t, func() {
		number := int64(123456789)
		Convey("Should return false", func() {
			So(helper.HasLen(number, 8), ShouldBeFalse)
			So(helper.HasLen(number, 10), ShouldBeFalse)
		})
	})
}

func TestGetSubDigit(t *testing.T) {
	Convey("When getting subdigits of a number", t, func() {
		len := 9
		number := int64(123456789)
		Convey("Should return the expected values", func() {
			So(helper.GetSubDigit(number, len, 1, 3), ShouldEqual, 123)
			So(helper.GetSubDigit(number, len, 4, 8), ShouldEqual, 45678)
			So(helper.GetSubDigit(number, len, 9, 9), ShouldEqual, 9)
		})
	})
}
