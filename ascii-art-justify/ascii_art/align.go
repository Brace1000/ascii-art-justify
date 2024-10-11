package ascii_art

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	dim := strings.Fields(string(out))

	w, _ := strconv.Atoi(dim[1])

	return w
}

func Align(str, align string, lenColor int) string {
	if len(str) >= getTerminalSize() {
		return str
	}

	padding := getTerminalSize() - (len(str) - lenColor)

	if align == "center" {
		padding = (padding / 2)
	}

	if align != "justify" {
		return strings.Repeat(" ", padding) + str
	}

	arrStr := strings.Split(str, "j")
	if len(arrStr) < 2 {
		return str
	}

	spaces := len(arrStr) - 1
	padding /= spaces
	spaceWidth := strings.Repeat(" ", padding)
	addSpace := padding % spaces
	justStr := ""

	for i, wordArt := range arrStr {
		justStr += wordArt
		if i < len(arrStr)-1 {
			justStr += spaceWidth
		}
		if addSpace > 0 {
			justStr += " "
			addSpace --
		}
	}

	return justStr
}
