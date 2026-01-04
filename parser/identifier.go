package parser

import (
	"github.com/NuruProgramming/KaNuru/ast"
)

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}
