package ascii_art

import (
	"fmt"
	"os"
	"path"
	"strings"
)

func Output(str, outputFile string) {
	if path.Ext(outputFile) != ".txt" {
		fmt.Println("Error: The output file must be a text file")
		os.Exit(0)
	}

	if strings.Contains(outputFile, "/") {
		fmt.Println("Error: incorrect output file format \n The file name should not contain / character")
		os.Exit(0)
	}

	err := os.WriteFile(outputFile, []byte(str), 0o644)
	if err != nil {
		fmt.Println("Erorr writing file", err)
		os.Exit(1)
	}
	fmt.Println(outputFile, "written successfully")
}
