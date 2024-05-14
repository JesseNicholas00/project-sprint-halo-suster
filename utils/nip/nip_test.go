package nip_test

import (
	"testing"

	"github.com/JesseNicholas00/HaloSuster/utils/nip"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNip(t *testing.T) {
	Convey("When creating a NIP", t, func() {
		role := nip.RoleIt
		gender := nip.GenderMale
		year := 2002
		month := 12
		suffix := 420

		val := nip.New(role, gender, year, month, suffix)
		Convey("Should return a NIP with the correct values", func() {
			So(nip.IsValid(val), ShouldBeTrue)
			So(nip.GetRole(val), ShouldEqual, role)
			So(nip.GetGender(val), ShouldEqual, gender)
			So(nip.GetYear(val), ShouldEqual, year)
			So(nip.GetMonth(val), ShouldEqual, month)
			So(nip.GetSuffix(val), ShouldEqual, suffix)
		})
	})
}
