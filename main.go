package main

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/sartyukhov/structfieldschecker/internal"
)

func main() {
	var Analyzer = &analysis.Analyzer{
		Name: "structfieldschecker",
		Doc:  "Checks that all fields in a struct are initialized",
		Run:  internal.Run,
	}

	singlechecker.Main(Analyzer)
}
