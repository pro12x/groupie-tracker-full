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
	return str
}

func Capitalize(str string) string {
	return strings.Title(str)
}

func ProcessString(model models.Location) models.Location {
	for i, location := range model.Locations {
		model.Locations[i] = FormatString(location)
		model.Locations[i] = Capitalize(model.Locations[i])
	}
	return model
}

// ProcessDatesLocations processes the dates and locations in the model
func ProcessDatesLocations(model models.Relation) models.Relation {
	for old, dates := range model.DatesLocations {
		key := FormatString(old)
		key = Capitalize(key)
		model.DatesLocations[key] = dates
		delete(model.DatesLocations, old)
		/*for i, date := range dates {
			model.DatesLocations[key][i] = strings.TrimPrefix(date, "*")
			model.DatesLocations[key][i] = FormatDate(model.DatesLocations[key][j])
		}*/
	}
	return model
}
