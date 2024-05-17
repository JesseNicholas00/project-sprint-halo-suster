package nip

import (
	"time"

	"github.com/JesseNicholas00/HaloSuster/utils/helper"
)

const (
	NipLengthMin = 13
	NipLengthMax = 15

	SuffLengthMin = 3
	SuffLengthMax = 5
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

func newWithSuffLen(
	role NipRole,
	gender NipGender,
	year int,
	month int,
	suffix int,
	suffLen int,
) int64 {
	rolePart := int64(role) % helper.Pow10[3] * helper.Pow10[suffLen+7]
	genderPart := int64(gender) % helper.Pow10[1] * helper.Pow10[suffLen+6]
	yearPart := int64(year) % helper.Pow10[4] * helper.Pow10[suffLen+2]
	monthPart := int64(month) % helper.Pow10[2] * helper.Pow10[suffLen]
	suffixPart := int64(suffix) % helper.Pow10[suffLen]

	return rolePart + genderPart + yearPart + monthPart + suffixPart
}

func New(
	role NipRole,
	gender NipGender,
	year int,
	month int,
	suffix int,
) int64 {
	suffix64 := int64(suffix)

	for suffLen := SuffLengthMin; suffLen < SuffLengthMax; suffLen++ {
		if helper.HasLen(suffix64, suffLen) {
			return newWithSuffLen(
				role,
				gender,
				year,
				month,
				suffix,
				suffLen,
			)
		}
	}

	return newWithSuffLen(
		role,
		gender,
		year,
		month,
		suffix,
		SuffLengthMax,
	)
}

func getLen(nip int64) int {
	for len := NipLengthMin; len <= NipLengthMax; len++ {
		if helper.HasLen(nip, len) {
			return len
		}
	}
	return -1
}

func IsValid(nip int64) bool {
	len := getLen(nip)
	if len == -1 {
		return false
	}

	role := GetRole(nip, len)
	if !(role == RoleIt || role == RoleNurse) {
		return false
	}

	gender := GetGender(nip, len)
	if !(gender == GenderMale || gender == GenderFemale) {
		return false
	}

	year := GetYear(nip, len)
	if !helper.IsBetween(year, 2000, curYear) {
		return false
	}

	month := GetMonth(nip, len)
	return helper.IsBetween(month, 1, 12)
}

func getLenFromParams(nip int64, length []int) int {
	if len(length) == 0 {
		return getLen(nip)
	}
	return length[0]
}

func GetRole(nip int64, length ...int) NipRole {
	return NipRole(helper.GetSubDigit(nip, getLenFromParams(nip, length), 1, 3))
}

func GetGender(nip int64, length ...int) NipGender {
	return NipGender(
		helper.GetSubDigit(nip, getLenFromParams(nip, length), 4, 4),
	)
}

func GetYear(nip int64, length ...int) int {
	return int(helper.GetSubDigit(nip, getLenFromParams(nip, length), 5, 8))
}

func GetMonth(nip int64, length ...int) int {
	return int(helper.GetSubDigit(nip, getLenFromParams(nip, length), 9, 10))
}

func GetSuffix(nip int64, length ...int) int {
	curLen := getLenFromParams(nip, length)
	return int(helper.GetSubDigit(nip, curLen, 11, curLen))
}
