package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	a "ascii/ascii_art"
)

const (
	colorErr = `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`

	justifyErr = `Usage: go run . [OPTION] [STRING] [BANNER]

Example: go run . --align=right something standard`

	genErr = `Usage: go run . [OPTIONs] [STRING]

EX: go run . "something"`

	fsErr = `Usage: go run . [STRING] [BANNER]

EX: go run . something standard`

	outputErr = `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`

	flagErr = `Flag list:
--color
--align
--output`
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(genErr)
		os.Exit(1)
	}

	// Default arguments
	input := ""
	bannerFile := "standard"
	subString := input

	flgColor := flag.String("color", "", "Color")
	flgAlign := flag.String("align", "", "Align")
	flgOutput := flag.String("output", "", "Output")

	flag.Parse()
	flagPattern := regexp.MustCompile(`^--\w+=.+$`)

	flagArgs := os.Args[1 : flag.NFlag()+1]
	for _, arg := range flagArgs {
		if !flagPattern.MatchString(arg) {
			if strings.Contains(arg, "color") {
				fmt.Println(colorErr)
			} else if strings.Contains(arg, "align") {
				fmt.Println(justifyErr)
			} else if strings.Contains(arg, "output") {
				fmt.Println(outputErr)
			} else {
				fmt.Println(flagErr)
			}
			os.Exit(0)
		}
	}

	color := a.ColorPicker(*flgColor)
	align := *flgAlign
	outputFile := *flgOutput

	args := flag.Args() // Non flags arguments
	nArgs := len(args)  // Count of non flag arguments

	if color == "" {
		switch nArgs {
		case 1:
			input = args[0]
			subString = input
		case 2:
			input = args[0]
			bannerFile = args[1]
			subString = input
		default:
			fmt.Println(fsErr)
			os.Exit(0)
		}
	} else {
		switch nArgs {
		case 1:
			input = args[0]
			subString = input
		case 2:
			input = args[0]
			subString = args[0]
			bannerFile = args[1]
		case 3:
			subString = args[0]
			input = args[1]
			bannerFile = args[2]
		default:
			fmt.Println(genErr)
			os.Exit(0)
		}
	}

	if strings.Contains(bannerFile, ".") {
		fmt.Println(fsErr)
		os.Exit(0)
	}
	bannerFile = "./.banners/" + bannerFile + ".txt"
	contents := a.GetFile(bannerFile)

	str := a.ProcessInput(contents, input, color, subString, align)

	if outputFile == "" {
		if align == "" {
			fmt.Print(str)
		} else {
			if align != "left" && align != "center" && align != "justify" && align != "right" {
				fmt.Println("Error: Align option are left, center, right or justify")
				os.Exit(0)
			}
			print(str)
		}
	} else {
		a.Output(str, outputFile)
	}
}
