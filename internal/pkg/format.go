package pkg

import (
	"fmt"
	"os"
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
		os.Exit(0)
	}
	return format.Format("Jan 2, 2006")
}

func FormatString(str string) string {
	reg := regexp.MustCompile(`[-_]`)
	str = reg.ReplaceAllString(str, " ")
	return str
}

func Capitalize(str string) string {
	return strings.Title(str)
}
