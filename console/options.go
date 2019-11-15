package console

import (
	"strings"

	"github.com/urfave/cli"
)

// ConsoleOption console arguments representations
type ConsoleOption map[string]string

// ConvertArgsToMap convert console arguments to map
func ConvertArgsToMap(args cli.Args) ConsoleOption {

	options := make(ConsoleOption)

	for i := 0; i < args.Len(); i++ {
		parsedOption := strings.Split(args.Get(i), "=")

		if len(parsedOption) == 1 {
			options[parsedOption[0]] = "true"
		} else if len(parsedOption) == 2 {
			options[parsedOption[0]] = parsedOption[1]
		}
	}

	return options
}
