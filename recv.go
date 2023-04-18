package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

func recvChecker(node *ast.File, fset *token.FileSet) {
	printNode(node, fset, 0)

	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		if callExpr, ok := n.(*ast.CallExpr); ok {
			fmt.Printf("callExpr: %+v\n", callExpr.Fun)

			if ident, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				fmt.Printf("ident: %+v\n", ident.X)

				if fmt.Sprint(ident.X) == "mc" && ident.Sel.Name == "Get" {
					fmt.Printf("mc.Get() called in function %s\n", node.Name.Name)
					fmt.Printf("ident.Sel.Obj: %+v\n", ident.Sel.Obj)

					if funType, ok := ident.Sel.Obj.Decl.(*ast.FuncDecl); ok {
						if retType, ok := funType.Type.Results.List[0].Type.(*ast.Ident); ok {
							fmt.Printf("retType: %+v\n", retType)

							// if retType.Name == "Item" {
							// 	fmt.Printf("ch.Get() called in function %s\n", node.Name.Name)
							// }
						}
					}

				}
			}
		}

		return true
	})
}

func printNode(node ast.Node, fset *token.FileSet, indent int) {
	if node == nil {
		return
	}

	// 获取节点类型的名称
	nodeType := fmt.Sprintf("%T", node)
	nodeType = strings.TrimPrefix(nodeType, "*ast.")
	nodeType = strings.TrimSuffix(nodeType, "Stmt")
	nodeType = strings.TrimSuffix(nodeType, "Expr")

	// 获取节点的源代码位置信息
	position := fset.Position(node.Pos())
	fileName := position.Filename
	line := position.Line
	column := position.Column

	// 在终端中输出节点的信息
	fmt.Printf("%s%s (line: %d, col: %d, file: %s)\n", strings.Repeat("  ", indent), nodeType, line, column, fileName)

	// 递归输出子节点的信息
	for _, child := range getChilds(node) {
		printNode(child, fset, indent+4)
	}
}

// 获取一个节点的所有子节点
func getChilds(node ast.Node) []ast.Node {
	var childs []ast.Node

	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil || n == node {
			return true
		}

		if _, ok := n.(ast.Expr); ok {
			childs = append(childs, n)
			return false
		}

		childs = append(childs, n)
		return true
	})

	return childs
}
