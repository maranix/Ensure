package ensure

import "fmt"

type color string

// Red
// return fmt.Sprintf("\033[31m%s\033[0m", "Failure")
const (
	green  color = "\033[32m"
	red          = "\033[31m"
	yellow       = "\033[33m"
	reset        = "\033[0m"
)

func colorizeString(c color, s string) string {
	return fmt.Sprintf("%s%s%s", c, s, reset)
}
