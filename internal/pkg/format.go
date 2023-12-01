package pkg

import (
	"bim/internal/models"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// FormatDate returns a formatted date
func FormatDate(date string) string {
	date = strings.TrimPrefix(date, "*")
	format, err := time.Parse("02-01-2006", date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		// os.Exit(0)
	}
	return format.Format("Jan 2, 2006")
}

// ProcessDates returns a formatted date
func ProcessDates(model models.Date) models.Date {
	for i, date := range model.Dates {
		model.Dates[i] = strings.TrimPrefix(date, "*")
		model.Dates[i] = FormatDate(model.Dates[i])
	}
	return model
}

func FormatString(str string) string {
	reg := regexp.MustCompile(`[_]`)
	str = reg.ReplaceAllString(str, " ")
	reg1 := regexp.MustCompile(`[-]`)
	str = reg1.ReplaceAllString(str, " - ")
	str = strings.Title(str)
	return str
}

func ProcessString(model models.Location) models.Location {
	for i, location := range model.Locations {
		model.Locations[i] = FormatString(location)
	}
	return model
}

// ProcessDatesLocations processes the dates and locations in the model
func ProcessDatesLocations(model models.Relation) models.Relation {
	newMap := make(map[string][]string)
	for old, dates := range model.DatesLocations {
		newDates := make([]string, len(dates))
		key := FormatString(old)
		for i, date := range dates {
			/*model.DatesLocations[old][i] = strings.TrimPrefix(date, "*")
			model.DatesLocations[old][i] = FormatDate(model.DatesLocations[old][i])*/
			newDates[i] = FormatDate(date)
		}
		newMap[key] = newDates
	}
	model.DatesLocations = newMap
	return model
}

// ProcessDatesLocations processes the dates and locations in the model
func ProcessDatesLocation(model models.Relation) models.Relation {
	// Use a map instead of a struct to make the code more efficient
	newDatesLocations := make(map[string][]string)

	for old, dates := range model.DatesLocations {
		newDates := make([]string, len(dates))
		key := FormatString(old)

		for i, date := range dates {
			newDates[i] = FormatDate(strings.TrimPrefix(date, "*"))
		}

		newDatesLocations[key] = newDates
	}

	model.DatesLocations = newDatesLocations
	return model
}
