package parser

import (
	"github.com/NuruProgramming/KaNuru/ast"
	"github.com/NuruProgramming/KaNuru/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}
}
