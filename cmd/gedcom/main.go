package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func fatalln(args ...interface{}) {
	log.Fatalln(append([]interface{}{"ERROR:"}, args...)...)
}

func check(err error) {
	if err != nil {
		fatalln(err)
	}
}

func usage() string {
	lines := []string{
		"Missing command, use one of:",
		fmt.Sprintf("\t%s diff      - Compare gedcom files", os.Args[0]),
		fmt.Sprintf("\t%s publish   - Publish as HTML", os.Args[0]),
		fmt.Sprintf("\t%s query     - Query with gedcomq", os.Args[0]),
		fmt.Sprintf("\t%s tune      - Used to calculate ideal weights and similarities", os.Args[0]),
	}

	return strings.Join(lines, "\n")
}

func main() {
	if len(os.Args) < 2 {
		fatalln(usage())
	}

	switch os.Args[1] {
	case "diff":
		runDiffCommand()

	case "publish":
		runPublishCommand()

	case "query":
		runQueryCommand()

	case "tune":
		runTuneCommand()

	default:
		fatalln("unknown command:", os.Args[1])
	}
}
