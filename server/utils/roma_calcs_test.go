package utils_test

import (
	"fmt"
	"server/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanNumber(t *testing.T) {
	assert.Equal(t, "nihil", utils.GetRomanNumber(0))

	assert.Equal(t, "I", utils.GetRomanNumber(1))
	assert.Equal(t, "VIII", utils.GetRomanNumber(8))
	assert.Equal(t, "X", utils.GetRomanNumber(10))
	assert.Equal(t, "C", utils.GetRomanNumber(100))
	assert.Equal(t, "M", utils.GetRomanNumber(1000))
}

func TestRomanMonth(t *testing.T) {
	assert.Equal(t, "Ianuarius", utils.GetRomanMonth(1))
	assert.Equal(t, "December", utils.GetRomanMonth(12))
}

func TestRomanDay(t *testing.T) {
	assert.Equal(t, "Solis", utils.GetRomanDay(1))
	assert.Equal(t, "Saturni", utils.GetRomanDay(7))
}

func TestDaysInMonth(t *testing.T) {
	assert.Equal(t, 28, utils.GetDaysInMonth(2, 1999))
	assert.Equal(t, 29, utils.GetDaysInMonth(2, 2000))
}

func TestGetMonthDayExpression(t *testing.T) {
	//approach here is to verify by inspection, and that there is no panic that occurs

	// leap-year
	fmt.Println(2, 28, 2000, utils.GetMonthDayExpression(2, 28, 2000))
	fmt.Println(2, 29, 2000, utils.GetMonthDayExpression(2, 29, 2000))

	// non-leap-year
	for m := 1; m <= 12; m++ {
		for d := 1; d <= utils.GetDaysInMonth(m, 1999); d++ {
			fmt.Println(m, d, 1999, utils.GetMonthDayExpression(m, d, 1999))
		}
	}

}
