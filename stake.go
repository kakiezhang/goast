package main

import (
	"go/ast"
	"go/token"
)

func stake(fd *ast.FuncDecl) {
	if len(fd.Body.List) > 0 {
		list := make([]ast.Stmt, 0, len(fd.Body.List)+1)
		list = append(list, &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("fmt"),
					Sel: ast.NewIdent("Println"),
				},
				Args: []ast.Expr{&ast.BasicLit{
					Kind:  token.STRING,
					Value: `"i am the first line"`,
				}},
			},
		})
		list = append(list, fd.Body.List...)
		fd.Body.List = list
	}
}
