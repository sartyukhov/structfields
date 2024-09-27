package internal

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const triggerComment = "allrequired"

func Run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if compLit, ok := n.(*ast.CompositeLit); ok {
				if typ, ok := pass.TypesInfo.TypeOf(compLit).Underlying().(*types.Struct); ok {
					if hasTriggerComment(file, compLit) {
						requiredFields := map[string]bool{}

						for i := 0; i < typ.NumFields(); i++ {
							requiredFields[typ.Field(i).Name()] = false
						}

						for _, elt := range compLit.Elts {
							if kv, ok := elt.(*ast.KeyValueExpr); ok {
								if ident, ok := kv.Key.(*ast.Ident); ok {
									if _, exists := requiredFields[ident.Name]; exists {
										requiredFields[ident.Name] = true
									}
								}
							}
						}

						for fieldName, initialized := range requiredFields {
							if !initialized {
								pass.Reportf(compLit.Pos(), "field %s is required but not initialized", fieldName)
							}
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}

func hasTriggerComment(file *ast.File, compLit *ast.CompositeLit) bool {
	for _, commentGroup := range file.Comments {
		for _, comment := range commentGroup.List {
			if compLit.Pos() < comment.Pos() && compLit.End() > comment.End() {
				if strings.Contains(comment.Text, triggerComment) {
					return true
				}
			}
		}
	}
	return false
}
