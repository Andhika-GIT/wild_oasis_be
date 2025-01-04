package date

import (
	"fmt"
	"time"
)

func ConvertIsoToDate(isoDate string) string {

	// convert IsoDate to time.time
	parsedTime, err := time.Parse(time.RFC3339, isoDate)

	if err != nil {
		fmt.Println("error when parsing from iso to time : ", err)
		return ""
	}

	// convert time.Time back to string, with format "Mon Jan 02 2006 15:04:05"
	formattedDate := parsedTime.Format("Mon Jan 02 2006 15:04:05")

	return formattedDate

}

func EachDayOfInterval(startDate, endDate time.Time) []string {
	var result []string
	var formattedDate string

	for days := startDate; !days.After(endDate); days = days.Add(24 * time.Hour) {
		// convert time.Time back to string, with format "Mon Jan 02 2006 15:04:05"
		formattedDate = days.Format("Mon Jan 02 2006 15:04:05")
		result = append(result, formattedDate)
	}

	return result
}
