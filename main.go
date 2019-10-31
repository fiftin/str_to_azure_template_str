package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type State int
const (
	Text State = 0
	Dollar State = 1
	Statement State = 2
)

func main() {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	var reader *bufio.Reader

	if len(os.Args[1:]) > 0 {
		f, err := os.Open(os.Args[1])
		if err != nil {
			fmt.Println("First argument must be filename")
			return
		}
		reader = bufio.NewReader(f)
	} else {
		if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
			fmt.Println("The command is intended to work with pipes.")
			fmt.Println("Usage: str_to_azure_template_str < str")
			return
		}
		reader = bufio.NewReader(os.Stdin)
	}

	output := []rune("concat('")
	state := Text
	for {
		input, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		switch input {
		case '$':
			switch state {
			case Text:
				state = Dollar
			case Dollar:
				output = append(output, input)
				state = Text
			case Statement:
				output = append(output, input)
			}
		case '\\':
			switch state {
			case Dollar:
				output = append(output, '\\', '\\')
				state = Text
			default:
				output = append(output, '\\', '\\')
			}
		case '"':
			switch state {
			case Dollar:
				output = append(output, '\\', '"')
				state = Text
			default:
				output = append(output, '\\', '"')
			}
		case '{':
			switch state {
			case Dollar:
				output = append(output, '\'', ',')
				state = Statement
			default:
				output = append(output, input)
			}
		case '}':
			switch state {
			case Text:
				output = append(output, input)
			case Dollar:
				output = append(output, input)
				state = Text
			case Statement:
				output = append(output, ',', '\'')
				state = Text
			}
		case '\n':
			switch state {
			case Text:
				output = append(output, '\\', 'n')
			case Dollar:
				output = append(output, '\\', '"')
				state = Text
			case Statement:
				panic("Invalid new line")
			}
		case '\r':
		default:
			switch state {
			case Dollar:
				output = append(output, input)
				state = Text
			default:
				output = append(output, input)
			}
		}
	}
	output = append(output, '\'', ')')
	fmt.Print(string(output))
}