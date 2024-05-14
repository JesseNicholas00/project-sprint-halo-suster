package nip

import (
	"time"

	"github.com/JesseNicholas00/HaloSuster/utils/helper"
)

const (
	NipLength = 13
)

type NipGender int

const (
	GenderMale   = NipGender(1)
	GenderFemale = NipGender(2)
)

type NipRole int

const (
	RoleIt    = NipRole(615)
	RoleNurse = NipRole(303)
)

var curYear = time.Now().Year()

// format:
//  6  1  5 | 2 | 2 0 0 1 | 0 2 | 9 8 7
// 12 11 10   9   8 7 6 5   4 3   2 1 0

func New(
	role NipRole,
	gender NipGender,
	year int,
	month int,
	suffix int,
) int64 {
	rolePart := int64(role) % helper.Pow10[3] * helper.Pow10[10]
	genderPart := int64(gender) % helper.Pow10[1] * helper.Pow10[9]
	yearPart := int64(year) % helper.Pow10[4] * helper.Pow10[5]
	monthPart := int64(month) % helper.Pow10[2] * helper.Pow10[3]
	suffixPart := int64(suffix) % helper.Pow10[3]

	return rolePart + genderPart + yearPart + monthPart + suffixPart
}

func IsValid(nip int64) bool {
	if !helper.HasLen(nip, NipLength) {
		return false
	}

	role := GetRole(nip)
	if !(role == RoleIt || role == RoleNurse) {
		return false
	}

	gender := GetGender(nip)
	if !(gender == GenderMale || gender == GenderFemale) {
		return false
	}

	year := GetYear(nip)
	if !helper.IsBetween(year, 2000, curYear) {
		return false
	}

	month := GetMonth(nip)
	return helper.IsBetween(month, 1, 12)
}

func GetRole(nip int64) NipRole {
	return NipRole(helper.GetSubDigit(nip, NipLength, 1, 3))
}

func GetGender(nip int64) NipGender {
	return NipGender(helper.GetSubDigit(nip, NipLength, 4, 4))
}

func GetYear(nip int64) int {
	return int(helper.GetSubDigit(nip, NipLength, 5, 8))
}

func GetMonth(nip int64) int {
	return int(helper.GetSubDigit(nip, NipLength, 9, 10))
}

func GetSuffix(nip int64) int {
	return int(helper.GetSubDigit(nip, NipLength, 11, 13))
}
