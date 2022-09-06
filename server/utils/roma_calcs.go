package utils

import (
	"fmt"
	"strings"
)

func GetRomanNumber(x int) string {
	roman_number_bases_dec := [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	roman_letter_bases := [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	i := len(roman_number_bases_dec) - 1
	remainingNbr := x
	ret := ""

	//handle special case 0
	if x == 0 {
		return "nihil" //there is no predominant historical 'standard' as noted at https://en.wikipedia.org/wiki/Roman_numerals#Zero
	}

	for remainingNbr > 0 {
		div := remainingNbr / roman_number_bases_dec[i]
		remainingNbr = remainingNbr % roman_number_bases_dec[i]

		for div > 0 {
			ret += roman_letter_bases[i]
			div--
		}
		i--
	}

	return ret
}

// GetRomanMonth get month (1-based)
func GetRomanMonth(x int) string {
	months := [12]string{"Ianuarius", "Februarius", "Martius", "Aprilis", "Maius", "Iunius", "Iulius", "Augustus", "September", "October", "November", "December"}

	return months[x-1]
}

// GetRomanDay get day of week (1-based)
func GetRomanDay(x int) string {
	days := [7]string{"Solis", "Lunae", "Martis", "Mercuris", "Jovis", "Veneris", "Saturni"}

	return days[x-1]
}

// GetDaysInMonth get number of days in specified (1-based) month and year
func GetDaysInMonth(month int, year int) int {
	daysInMonth := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	const febLeapYearDays = 29
	const febMonthNumber = 2

	// check leap-year special case
	if month == febMonthNumber && year%4 == 0 {
		return febLeapYearDays
	}

	return daysInMonth[month-1]
}

// GetNextMonthNumber get next (1-based) month number, rolling over to next year if nec.
func GetNextMonthNumber(month int) int {
	if month == 12 {
		return 1
	} else {
		return month + 1
	}
}

// GetMonthDayExpression as described at https://www.wikihow.com/Write-in-Latin
func GetMonthDayExpression(month int, day int, year int) string {

	// variations by month for when kalendes, ides, nones fall (1-based day return)
	kalendesDays := [12]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	nonesDays := [12]int{5, 5, 7, 5, 7, 5, 7, 5, 5, 7, 5, 5}
	idesDays := [12]int{13, 13, 15, 13, 15, 13, 15, 13, 13, 15, 13, 13}

	const anteDiem = "ante diem"

	ret := ""
	yearToUse := year

	// month-date expression
	if day == kalendesDays[month-1] {
		ret = fmt.Sprintf("Kalendis %s", GetRomanMonth(month))
	} else if day == nonesDays[month-1] {
		ret = fmt.Sprintf("Nones %s", GetRomanMonth(month))
	} else if day == idesDays[month-1] {
		ret = fmt.Sprintf("Ides %s", GetRomanMonth(month))
	} else if day > idesDays[month-1] {
		ret = fmt.Sprintf("%s %s Kalendis %s", anteDiem, strings.ToLower(GetRomanNumber(
			GetDaysInMonth(month, year)-day+2)), GetRomanMonth(GetNextMonthNumber(month)))
		if GetNextMonthNumber(month) == 1 {
			yearToUse++
		}
	} else if day > nonesDays[month-1] {
		ret = fmt.Sprintf("%s %s Ides %s", anteDiem, strings.ToLower(GetRomanNumber(
			idesDays[month-1]-day+2)), GetRomanMonth(month))
	} else if day > kalendesDays[month-1] {
		ret = fmt.Sprintf("%s %s Nones %s", anteDiem, strings.ToLower(GetRomanNumber(
			nonesDays[month-1]-day+2)), GetRomanMonth(month))
	} else {
		panic(fmt.Sprintf("unsupported month %d day %d combination", month, day))
	}

	//year expression
	ret = fmt.Sprintf("%s %s A.D.", ret, GetRomanNumber(yearToUse)) // assume AD

	return ret
}
