//+build make

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	mainContents := `package main

func main()  {
	
}
`

	modFileContents := `module github.com/dlsniper/tipsandtricks/%s

go 1.11
`

	for i := 0; i < 42; i++ {
		currentIndex := fmt.Sprintf("tip%03d", i+1)
		dirName := "./" + currentIndex
		if err := os.Mkdir(dirName, os.ModeDir); err != nil {
			log.Fatalln(err)
		}

		mainFileName := fmt.Sprintf("%s/main.go", dirName)
		if err := ioutil.WriteFile(mainFileName, []byte(mainContents), os.ModePerm); err != nil {
			log.Fatalln(err)
		}

		modFileName := fmt.Sprintf("%s/go.mod", dirName)
		if err := ioutil.WriteFile(modFileName, []byte(fmt.Sprintf(modFileContents, currentIndex)), os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
}
