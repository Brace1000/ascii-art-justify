# ASCII Art Justfy

## Overview
ASCII Art justfy is a  program written in Go that converts input strings into their graphical representations using ASCII characters with the template  name and aligns it either to the left, right,centre or justify by itself.
The program  is also designed to handle colors , output and fs as a commandline argument .


## Features for fs 
```
* Convert any input string using ASCII characters into ASCII art.
* Add the name of the template of your choice as your second argument.
```

## Features for colors
```
* Convert any input string with ASCII characters into ASCII art.
* Customize the color of individual letters or sets of letters using the --color flag.
* Supports  only ansi color system e.g "black":  "\u001b[30m"],
```

## Installation
To use ASCII Art Justify,or ASCII Art fs and/or ASCII Art Color with fs , you'll need Go installed on your system. You can install it via the official Go website.

Clone the repository:
```
git clone https://learn.zone01kisumu.ke/git/joseowino/ascii-art-justify
```

## Usage for Justify
Navigate to the project directory and run the program with the desired input string and template name:
```
Usage: go run .[OPTION] [STRING] [BANNER]
```

### Example:
```
 Example: go run . --align=right something standand
 ```
 The progam will output the  Ascii representation of the input string  using the specified bannerfile. 

## Usage for color with fs
Navigate to the project directory and run the program with the desired  color options, input string and template name:
```
go run . [OPTIONS] "Your input string" [BANNERFILE NAME]
```
### Options
```
* --color=<color> <letters to be colored>: Specifies the color and letters to be colored in the output.
* If the letters are not specified, the whole string will be colored.
```
### Example I:
```
go run . --color=red "Hello, World!" | cat -E
```
The program will output the ASCII representation of the input string with the specified letters colored in red.

### Example II:
```
go run . --color=red "kit"  "king kit kitten kit"  shadow | cat -E
```
The program will output the ASCII representation of the input string with substring kit colored red.

## Project Structure
* main.go: Contains the main program logic.
* banners/: Directory containing ASCII banner templates.
* ascii_art/: Directory containing unit tests for the program and functions

## Testing
Unit tests are included to ensure the correctness of the program, including color functionality. Run tests using:
```
go test ./...
```
## Contributing
Contributions are welcome! If you have suggestions for improvements, please open an issue or create a pull request on the GitHub repository.

# License
This project is licensed under the MIT License -> [https://mit-license.org/]

## Contributers

* joseOwino
https://learn.zone01kisumu.ke/git/joseowino
* hilaromondi
https://learn.zone01kisumu.ke/git/hilaromondi
* bobaigwa
https://learn.zone01kisumu.ke/git/bobaigwa

## Note
Only limited number of bannerFiles have been used in this project i.e
* standard.txt
* thinkertoy.txt
* shadow.txt
* rings.txt
