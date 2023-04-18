package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"strings"
)

func addAnnotation(fd *ast.FuncDecl, annotation string) {
	// Step 1: Parse the old annotation information
	fmt.Printf("fd: %+v\n", fd)
	var oldAnnotations []string
	if fd.Doc != nil {
		for _, c := range fd.Doc.List {
			fmt.Printf("c.Text: %s\n", c.Text)
			if strings.HasPrefix(c.Text, "//") {
				oldAnnotations = append(oldAnnotations, strings.TrimSpace(c.Text[2:]))
			}
		}
	}

	// Step 2: Merge the new and old annotation information
	var newAnnotations []string
	newAnnotations = append(newAnnotations, annotation)
	newAnnotations = append(newAnnotations, oldAnnotations...)

	fmt.Printf("newAnnotations: %+v\n", newAnnotations)

	// Step 3: Generate the new annotation code
	var buf bytes.Buffer
	for _, a := range newAnnotations {
		buf.WriteString("// " + a + "\n")
	}

	fmt.Printf("buf: %+v\n", buf.String())
	fmt.Printf("fd.Pos(): %d\n", fd.Pos())
	fd.Doc = &ast.CommentGroup{List: []*ast.Comment{
		{
			Text: buf.String(),
		},
	}}
}
