package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	linters "github.com/sartyukhov/structfields"
)

func main() {
	singlechecker.Main(linters.StructFieldsAnalyzer)
}
