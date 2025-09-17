package utils

const (
	Red   = "\033[31m"
	Green = "\033[32m"
	Reset = "\033[0m"
)

func TextGreen(text string) string {
	return Green + text + Reset
}

func TextRed(text string) string {
	return Red + text + Reset
}
