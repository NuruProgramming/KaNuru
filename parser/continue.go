package parser

import (
	"github.com/NuruProgramming/KaNuru/ast"
	"github.com/NuruProgramming/KaNuru/token"
)

func (p *Parser) parseContinue() *ast.Continue {
	stmt := &ast.Continue{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
