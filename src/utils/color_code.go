package utils

const (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func TextGreen(text string) string {
	return Green + text + Reset
}
