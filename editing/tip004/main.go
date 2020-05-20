package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Code completion works in Go Templates as well

	// E.g. Change the message to add {{.Message}} and {{.Count}} instead of |

	//language=GoTemplate
	message := `{{- /*gotype: github.com/dlsniper/tipsandtricks/editing/tip004.Output*/ -}}
各位 | 晚上好!
We have | Gophers here today!
`
	printMessage(message, "Gopher")
}

type Output struct {
	Message string
	Count   int
}

func printMessage(msg, name string) {
	t := template.Must(template.New("output").Parse(msg))

	o := Output{Message: name}

	err := t.Execute(os.Stdout, o)
	if err != nil {
		log.Fatalln(err)
	}
}
