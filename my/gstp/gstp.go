package gstp

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
)

const specTab = "\t"
const specEnter = "\n"

func GenProto(r io.Reader, w io.Writer, exp regexp.Regexp) error {
	src, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var content string
	fset := token.NewFileSet()
	astf, err := parser.ParseFile(fset, "", string(src), parser.ParseComments)
	if err != nil {
		return err
	}

	cmap := ast.NewCommentMap(fset, astf, astf.Comments)

	var declCmt string
	for _, decl := range astf.Decls {
		d, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		declCmt = fmt.Sprintf("%s", genComment(cmap[d], specEnter))
		for _, spec := range d.Specs {
			switch spec.(type) {
			case *ast.TypeSpec:
				tpy := spec.(*ast.TypeSpec)
				name := tpy.Name.Obj.Name
				if !exp.MatchString(name) {
					continue
				}

				st, ok := tpy.Type.(*ast.StructType)
				if !ok {
					continue
				}

				content += declCmt
				content += specEnter
				content += genMsg(cmap, st, name)
			}
		}
		content += specEnter
	}

	_, err = fmt.Fprint(w, content)
	return err
}

func genMsg(cmap ast.CommentMap, st *ast.StructType, name string) string {
	var msg = fmt.Sprintf("message %s {\n", name)

	for i, field := range st.Fields.List {
		msg += fmt.Sprintf("%s\n", genComment(cmap[field], specTab))
		// 生成field
		msg += fmt.Sprintf("\t%s %s = %d;\n", genFiledTyp(field.Type), snakeName(field.Names[0].Name), i+1)
	}
	msg += fmt.Sprintf("}")

	return msg
}

func genComment(cg []*ast.CommentGroup, spec string) (comment string) {
	last := len(cg) - 1
	if last <= -1 {
		return ""
	}

	for _, c := range cg[last].List {
		comment += spec + c.Text
	}
	//comment += spec

	return comment
}

func genFiledTyp(expr ast.Expr) (name string) {
	switch expr.(type) {
	case *ast.Ident:
		name = getIdentName(expr.(*ast.Ident))

	case *ast.SelectorExpr:
		name = getSelectorExprName(expr.(*ast.SelectorExpr))

	case *ast.StarExpr:
		name = genFiledTyp(expr.(*ast.StarExpr).X)

	case *ast.MapType:
		typ := expr.(*ast.MapType)
		name = fmt.Sprintf("map<%s,%s>", genFiledTyp(typ.Key), genFiledTyp(typ.Value))

	case *ast.ArrayType:
		// TODO[Dokiy] 2022/8/2: deal [][]string type
		name = "repeated" + " " + genFiledTyp(expr.(*ast.ArrayType).Elt)

	}

	return name
}

func getSelectorExprName(expr *ast.SelectorExpr) (name string) {
	name = expr.Sel.Name
	if expr.Sel.Name == "Time" && genFiledTyp(expr.X) == "time" {
		name = "google.protobuf.Timestamp"
	}
	return
}

func getIdentName(ident *ast.Ident) (name string) {
	name = ident.Name
	if ident.Name == "int" {
		name = "int32"
	}
	return
}

func snakeName(name string) string {
	l := len(name)
	if l <= 0 {
		return ""
	}

	s := make([]byte, 0, l*2)
	s = append(s, name[0])
	for i := 1; i < l; i++ {
		p, c := name[i-1], name[i]
		if p != '_' && c >= 'A' && c <= 'Z' {
			s = append(s, '_')
		}
		s = append(s, c)
	}

	return strings.ToLower(string(s))
}
