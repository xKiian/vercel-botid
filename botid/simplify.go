package botid

import (
	"github.com/t14raptor/go-fast/ast"
)

type FromCharCodeReplacerVisitor struct {
	ast.NoopVisitor
}

func (v *FromCharCodeReplacerVisitor) VisitExpression(n *ast.Expression) {
	n.VisitChildrenWith(v)

	callExpr, ok := n.Expr.(*ast.CallExpression)
	if !ok {
		return
	}

	if len(callExpr.ArgumentList) != 1 {
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

	if prop.Name != "fromCharCode" {
		return
	}

	code := callExpr.ArgumentList[0]
	if numLit, ok := code.Expr.(*ast.NumberLiteral); ok {
		value := string(rune(int(numLit.Value)))
		*n = ast.Expression{
			Expr: &ast.StringLiteral{
				Value: value,
			},
		}
	}
}

func ReplaceFromCharCode(p *ast.Program) {

	r := &FromCharCodeReplacerVisitor{}
	r.V = r
	p.VisitWith(r)
}
