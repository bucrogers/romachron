package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"server/models"
	"server/utils"
)

// get time in roman
func GetRomaDateTime(c *gin.Context) {
	tm := time.Now().Local() // Local() required in container context

	var romaTime models.RomaTime
	romaTime.Hr = utils.GetRomanNumber(tm.Hour())
	romaTime.Min = utils.GetRomanNumber(tm.Minute())
	romaTime.Sec = utils.GetRomanNumber(tm.Second())

	if tm.Hour() > 12 {
		romaTime.Hr12hFormat = utils.GetRomanNumber(tm.Hour() - 12)
		romaTime.AmPm = "p.m."
	} else {
		romaTime.Hr12hFormat = utils.GetRomanNumber(tm.Hour())
		romaTime.AmPm = "a.m."
	}

	romaTime.DateExpression = utils.GetMonthDayExpression(int(tm.Month()), tm.Day(), tm.Year())
	romaTime.DayExpression = utils.GetDayOfWeekExpression(int(tm.Month()), tm.Day(), tm.Year())

	c.JSON(http.StatusOK, romaTime)
}
