package ascii_art

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func colorMap(text, subString string) map[int]string {
	mapText := map[int]string{}
	colorStart := strings.Index(text, subString)
	colorReset := len(text) - 1
	if colorStart != -1 {
		colorReset = colorStart + len(subString) - 1
	}

	for i := 0; i < len(text); i++ {
		if colorStart == -1 {
			mapText[i] = "NA"
			continue
		}

		if text[i] == '\n' && i > colorStart && i < colorReset {
			mapText[i] = "CS"
			continue
		}

		if i < colorStart || (i > colorStart && i < colorReset) {
			mapText[i] = "NA"
			continue
		}

		if i == colorStart {
			mapText[i] = "CS"
			continue
		}

		if i == colorReset && i < len(text)-1 {
			mapText[i] = "CR"
			colorStart = strings.Index(text[i+1:], subString)
			colorReset = len(text) - 1
			if colorStart > -1 {
				colorStart += i + 1
				colorReset = colorStart + len(subString) - 1
			}
			continue
		}

		mapText[i] = "NA"

	}
	return mapText
}


// ColorPicker maps a color name or code to its corresponding ANSI color code
func ColorPicker(color string) (colorCode string) {
	if color == "" {
		return ""
	}

	
	colorChat := map[string]string{
		"reset":          "\u001b[39m",
		"black":          "\u001b[30m",
		"red":            "\u001b[31m",
		"green":          "\u001b[32m",
		"yellow":         "\u001b[33m",
		"blue":           "\u001b[34m",
		"magenta":        "\u001b[35m",
		"cyan":           "\u001b[36m",
		"white":          "\u001b[37m",
		"bright_black":   "\u001b[90m",
		"bright_red":     "\u001b[91m",
		"bright_green":   "\u001b[92m",
		"bright_yellow":  "\u001b[93m",
		"bright_blue":    "\u001b[94m",
		"bright_magenta": "\u001b[95m",
		"bright_cyan":    "\u001b[96m",
		"bright_white":   "\u001b[97m",
		"orange":         "\033[38;5;208m",
		"violet":         "\033[38;5;129m",
		"indigo":         "\033[38;5;54m",
		"maroon":         "\033[38;5;52m",
		"purple":         "\033[38;5;165m",
		"pink":           "\033[38;5;206m",
		"brown":          "\033[38;5;130m",
		"wheat":          "\033[38;5;229m",
		"tomato":         "\033[38;5;209m",
		"smoke":          "\033[38;5;245m",
		"gray":           "\033[38;5;240m",
		"gold":           "\033[38;5;178m",
		"avocado":        "\033[38;5;58m",
		"oceanblue":      "\033[38;5;27m",
		"navyblue":       "\033[38;5;18m",
		"amber":          "\033[38;5;214m",
	}

	// Retrieve the color code from the map
	colorCode, ok := colorChat[color]

	if ok {
		return colorCode
	}

	// If the color is not found, check if it is an ANSI code
	ansi, err := strconv.Atoi(color)
	if err != nil || ansi < 0 || ansi > 255 {
		fmt.Println("Error: color Not Found")
		os.Exit(1)
	}

	// Return the ANSI color code
	return fmt.Sprintf("\033[38;5;%dm", ansi)
}
