package main

import (
	"golang.org/x/tools/go/analysis"

	linters "github.com/sartyukhov/structfields"
)

func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{linters.StructFieldsAnalyzer}, nil
}
