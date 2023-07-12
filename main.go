package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
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

	tokenRecvCheck(f)

	// // Step 2: Find the function with the specified name
	// fd := findFunction(fname, f)
	// if fd == nil {
	// 	log.Fatal("Function not found!")
	// }

	// Step 3: Add the new annotation to the function

	// stake(fd)

	// Step 4: Write the modified AST back to file
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, f); err != nil {
		log.Fatal(err)
	}

	// if err := ioutil.WriteFile(fpath, buf.Bytes(), 0644); err != nil {
	// 	log.Fatal(err)
	// }
}

func tokenRecvCheck(f *ast.File) {
	r := "mc"
	m := "Get"

	ok := true
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.AssignStmt:
			if x.Tok == token.DEFINE || x.Tok == token.ASSIGN {
				i := 0

				if callExpr, ok := x.Rhs[i].(*ast.CallExpr); ok {
					// fmt.Printf("callExpr: %+v\n", callExpr.Fun)

					if ident, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
						// fmt.Printf("ident.X: %+v\n", ident.X)

						if ident.Sel.Name == m {
							if fmt.Sprint(ident.X) == r {
								fmt.Printf("dsn sep called: %s.%s()\n", r, m)
								ok = false
								// return false
							}
						}
					}
				}
			}
		}

		fmt.Println("hhhhhh")
		return ok
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
