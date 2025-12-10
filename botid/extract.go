package botid

import (
	"strings"

	"github.com/t14raptor/go-fast/ast"
	"github.com/t14raptor/go-fast/generator"
	"github.com/t14raptor/go-fast/parser"
	"github.com/t14raptor/go-fast/transform/simplifier"
	deobf "github.com/xkiian/obfio-deobfuscator"
)

func ExtractFromScript(script *string) (*ScriptCtx, error) {
	parsed, err := parser.ParseFile(*script)
	if err != nil {
		return nil, err
	}

	deobf.Deobfuscate(parsed) // magic!!
	simplifier.Simplify(parsed, false)
	ReplaceFromCharCode(parsed)
	simplifier.Simplify(parsed, false) //its just for string addition but i cba

	//os.WriteFile("out.js", []byte(generator.Generate(parsed)), 0644)

	return runExtractionVisitor(parsed), nil
}

type ScriptCtx struct {
	ast.NoopVisitor

	assignments map[ast.Id]string

	key  string
	seed float64

	arg1      float64
	arg2      float64
	rand      float64
	signature string
	version   string
}

func (v *ScriptCtx) VisitAssignExpression(n *ast.AssignExpression) {
	n.VisitChildrenWith(v)
	right, ok := n.Right.Expr.(*ast.StringLiteral)
	if !ok {
		return
	}

	left, ok := n.Left.Expr.(*ast.Identifier)
	if !ok {
		return
	}

	v.assignments[left.ToId()] = right.Value
}

func (v *ScriptCtx) VisitCallExpression(n *ast.CallExpression) {
	n.VisitChildrenWith(v)

	args := n.ArgumentList

	switch len(args) {
	case 5:
		arg1, ok := args[0].Expr.(*ast.NumberLiteral)
		if !ok {
			return
		}
		v.arg1 = arg1.Value

		arg2, ok := args[1].Expr.(*ast.NumberLiteral)
		if !ok {
			return
		}
		v.arg2 = arg2.Value

		arg3, ok := args[2].Expr.(*ast.NumberLiteral)
		if !ok {
			return
		}
		v.rand = arg3.Value

		arg4, ok := args[3].Expr.(*ast.StringLiteral)
		if !ok {
			return
		}
		v.signature = arg4.Value

		arg5, ok := args[4].Expr.(*ast.StringLiteral)
		if !ok {
			return
		}
		v.version = arg5.Value
	case 2:

		callExpr, ok := args[0].Expr.(*ast.CallExpression)
		if !ok {
			return
		}

		callee, ok := callExpr.Callee.Expr.(*ast.MemberExpression)
		if !ok {
			return
		}

		prop, ok := callee.Property.Prop.(*ast.Identifier)
		if !ok {
			return
		}
		if prop.Name != "join" {
			return
		}

		elements, ok := callee.Object.Expr.(*ast.ArrayLiteral)
		if !ok {
			return
		}

		first, ok := elements.Value[0].Expr.(*ast.Identifier)
		if !ok {
			return
		}
		found := v.assignments[first.ToId()]
		if found == "" {
			return
		}
		v.key = found

		second, ok := elements.Value[1].Expr.(*ast.Identifier)
		if !ok {
			return
		}
		found = v.assignments[second.ToId()]
		if found == "" {
			return
		}
		v.key += found

		/*if callExpr.Operator.String() != "+" {
			return
		}

		left, ok := callExpr.Left.Expr.(*ast.Identifier)
		if !ok {
			return
		}
		found := v.assignments[left.ToId()]
		fmt.Print(found)
		if found == "" {
			return
		}
		v.key = found

		right, ok := callExpr.Right.Expr.(*ast.Identifier)
		if !ok {
			return
		}
		found = v.assignments[right.ToId()]
		fmt.Print(found)
		if found == "" {
			return
		}
		v.key += found*/
	}
}

func (v *ScriptCtx) VisitObjectLiteral(n *ast.ObjectLiteral) {
	n.VisitChildrenWith(v)
	if !strings.Contains(generator.Generate(n), "(window)") {
		return
	}

	for _, prop := range n.Value {
		propKeyed, ok := prop.Prop.(*ast.PropertyKeyed)
		if !ok {
			continue
		}

		strLit, ok := propKeyed.Key.Expr.(*ast.StringLiteral)
		if !ok {
			continue
		}

		if strLit.Value != "S" { //maybe needs better filtering in the future
			continue
		}
		num, ok := propKeyed.Value.Expr.(*ast.NumberLiteral)
		if !ok {
			continue
		}

		v.seed = num.Value
		return
	}

}

func runExtractionVisitor(program *ast.Program) *ScriptCtx {
	f := &ScriptCtx{
		assignments: make(map[ast.Id]string),
	}
	f.V = f
	program.VisitWith(f)
	return f
}
