package main

import "go/ast"

func findFunction(name string, f *ast.File) *ast.FuncDecl {
	var result *ast.FuncDecl
	ast.Inspect(f, func(n ast.Node) bool {
		if fd, ok := n.(*ast.FuncDecl); ok {
			if fd.Name.Name == name {
				result = fd
				return false
			}
		}
		return true
	})
	return result
}
