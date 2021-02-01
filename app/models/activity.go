package models

import "gin-admin/pkg/utils"

type Activity struct {
	Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	StartDate   utils.Time `json:"startDate"`
	EndDate     utils.Time `json:"endDate"`
}
