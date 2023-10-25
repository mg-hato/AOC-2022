package models

import (
	c "aoc/common"
	"fmt"
)

type SnafuNumber struct {
	snafu string
}

func SnafuNumberFromString(s string) SnafuNumber {

	return SnafuNumber{
		snafu: string(c.Reverse([]rune(s))),
	}
}

func SnafuNumberFromInt(i int64) SnafuNumber {
	if i <= 0 {
		return SnafuNumber{snafu: "0"}
	}
	digits := []int64{}
	x := i
	for x > 0 {
		digits = append(digits, x%5)
		x /= 5
	}

	snafu_digits := []rune{}
	carry := int64(0)
	for _, digit := range digits {
		digit += carry
		carry = digit / 5
		digit %= 5
		switch digit {
		case 0, 1, 2:
			snafu_digits = append(snafu_digits, '0'+rune(digit))
		case 3:
			snafu_digits = append(snafu_digits, '=')
			carry++
		case 4:
			snafu_digits = append(snafu_digits, '-')
			carry++
		default:
			fmt.Printf("warning: digit is %d", digit)
		}
	}
	switch carry {
	case 1, 2:
		snafu_digits = append(snafu_digits, '0'+rune(carry))
	case 0:
	default:
		fmt.Printf("warning: carry is %d", carry)
	}

	return SnafuNumber{
		snafu: string(snafu_digits),
	}
}

func (s SnafuNumber) String() string {
	return string(c.Reverse([]rune(s.snafu)))
}

func (s SnafuNumber) ToInt() int64 {

	number := int64(0)
	power := int64(1)
	snafu_digit_to_int := map[rune]int64{
		'0': int64(0),
		'1': int64(1),
		'2': int64(2),
		'-': int64(-1),
		'=': int64(-2),
	}

	for _, snafu_digit := range s.snafu {
		number += power * snafu_digit_to_int[snafu_digit]
		power *= 5
	}
	return number
}
