package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	buildAst()
}

func buildAst() {
	// Step 1: Parse source code into AST
	fpath := "example.go"
	// fname := "foo"

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fpath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	tokenRecvCheck(f, "mc", "Get")

	// // Step 2: Find the function with the specified name
	// fd := findFunction(fname, f)
	// if fd == nil {
	// 	log.Fatal("Function not found!")
	// }

	// Step 3: Add the new annotation to the function

	// stake(fd)

	// // Step 4: Write the modified AST back to file
	// var buf bytes.Buffer
	// if err := printer.Fprint(&buf, fset, f); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := ioutil.WriteFile(fpath, buf.Bytes(), 0644); err != nil {
	// 	log.Fatal(err)
	// }
}

func tokenRecvCheck(f *ast.File, r, m string) {
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			if x.Tok == token.DEFINE || x.Tok == token.ASSIGN {
				i := 0

				// callExpr := x.Rhs[i].(*ast.CallExpr)
				// fnName := callExpr.Fun.(*ast.SelectorExpr).Sel.Name

				// fmt.Printf("x.Rhs: %+v\n", x.Rhs)

				if callExpr, ok := x.Rhs[i].(*ast.CallExpr); ok {
					// fmt.Printf("callExpr: %+v\n", callExpr.Fun)

					if ident, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
						// fmt.Printf("ident.X: %+v\n", ident.X)

						if ident.Sel.Name == m {
							if fmt.Sprint(ident.X) == r {
								fmt.Printf("dsn sep called: %s.%s()\n", r, m)
								// fmt.Printf("aaa x.Lhs: %+v\n", x.Lhs)
							} else {
								isSrcDotGet(ident.X, r, m)
							}
						}
					}
				}

				if len(x.Lhs) == 1 {
					r = x.Lhs[0].(*ast.Ident).Name
					// m = "Get"
				}
			}
		}
		return true
	})
}

func isSrcDotGet(n ast.Node, r, m string) bool {
	if callExpr, ok := n.(*ast.CallExpr); ok {
		// fmt.Printf("callExpr: %+v\n", callExpr.Fun)

		if ident, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
			// fmt.Printf("ident.X: %+v\n", ident.X)

			if fmt.Sprint(ident.X) == r && ident.Sel.Name == m {
				fmt.Printf("chain called: %s.%s()\n", r, m)
				return true
			}
		}
	}

	return false
}
