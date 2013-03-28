package main

import (
	"go/parser"
	"go/token"
	"go/ast"
	"go/printer"
	"os"
	"strings"
	"path"
	"fmt"
)

type Visitor struct {
}

func (x *Visitor) Visit(node ast.Node) ast.Visitor {
	if ident, ok := node.(*ast.Ident); ok {
		pprefix := "optimization_messages_"

		if strings.HasPrefix(ident.Name, pprefix) {
			ident.Name = ident.Name[len(pprefix):]
		}
	}

	return x
}

func main() {
	for _, file := range os.Args[1:] {
		if len(file) == 0 {
			continue
		}

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, file, nil, 0)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse `%s': %s\n", file, err)
			os.Exit(1)
		}

		// Change the package name
		pprefix := "optimization_messages_"

		if strings.HasPrefix(f.Name.Name, pprefix) {
			f.Name.Name = f.Name.Name[len(pprefix):]
		}

		// Change imports
		for _, imp := range f.Imports {
			if imp.Name != nil && strings.HasPrefix(imp.Name.Name, pprefix) {
				// Change the name
				imp.Name = nil

				// Change the path
				imp.Path.Value = `"ponyo.epfl.ch/go/get/optimization-go/optimization/messages/` + path.Base(imp.Path.Value)
			}
		}

		ast.Walk(&Visitor{}, f)

		ff, err := os.Create(file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create file `%s': %s\n", file, err)
			os.Exit(1)
		}

		printer.Fprint(ff, fset, f)
		ff.Close()
	}
}
