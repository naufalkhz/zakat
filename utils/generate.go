package utils

import (
	"fmt"
	"strings"
	"time"
)

// 230731/PHS/6randomNumber
func GenerateCode(code string) string {
	if code == "" {
		code = "DFL"
	}
	t := time.Now()

	currentDate := fmt.Sprintf("%d%02d%02d", t.Year()%100, t.Month(), t.Day())
	// Generate 8 Number
	randomNumber := fmt.Sprint(t.Nanosecond())[:8]

	return fmt.Sprintf("%s/%s/%s", currentDate, strings.ToUpper(code[0:3]), randomNumber)
}
