package parser

import (
	"github.com/NuruProgramming/KaNuru/ast"
)

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{Token: p.curToken}
}
