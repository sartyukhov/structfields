package linters

import (
	"golang.org/x/tools/go/analysis"

	"github.com/sartyukhov/structfields/internal"
)

var StructFieldsAnalyzer = &analysis.Analyzer{
	Name: "structfields",
	Doc:  "Checks that all fields in a struct are initialized",
	Run:  internal.Run,
}
