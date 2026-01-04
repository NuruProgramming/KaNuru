package parser

import (
	"github.com/NuruProgramming/KaNuru/ast"
	"github.com/NuruProgramming/KaNuru/token"
)

func (p *Parser) parseBreak() *ast.Break {
	stmt := &ast.Break{Token: p.curToken}
	for p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}
