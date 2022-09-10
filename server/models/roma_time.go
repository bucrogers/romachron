package models

type RomaTime struct {
	Hr          string `json:"hr"`
	Hr12hFormat string `json:"hr_12h_format""`
	AmPm        string `json:"am_pm"`
	Min         string `json:"min"`
	Sec         string `json:"sec"`

	DateExpression string `json:"date_expression""`
	DayExpression  string `json:"day_expression""`
}
