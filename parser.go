/*
Grammer:
	statementList 		: statement
							| statementList
	statement	  		: assignmentStatement
							| empty
	assignmentStatement : factor ARROW expr
	expr				: factor (COMMA factor)*
	factor 				: STRING
*/
package main

type StringNode struct {
	token *Token
}

type ExprNode struct {
	nodes []StringNode
}

type AssignmentNode struct {
	left  StringNode
	right ExprNode
}

type StatementListNode struct {
	nodes []*AssignmentNode
}

type Parser struct {
	lexer    Lexer
	curToken Token
}

func (p *Parser) eat(tt string) {
	if p.curToken._type == tt {
		p.curToken = p.lexer.getNextToken()
	}
}

func (p *Parser) getFactor() StringNode {
	tt := p.curToken
	if tt._type == STRING {
		return StringNode{&tt}
	}
	return StringNode{nil}
}

func (p *Parser) expr() ExprNode {
	var exprNode ExprNode

	exprNode.nodes = []StringNode{p.getFactor()}
	for p.curToken._type == COMMA {
		p.eat(COMMA)
		exprNode.nodes = append(exprNode.nodes, p.getFactor())
	}
	return exprNode
}

func (p *Parser) assignmentStatement() *AssignmentNode {
	variable := p.getFactor()
	p.eat(ARROW)
	exprNode := p.expr()
	return &AssignmentNode{variable, exprNode}
}

func (p *Parser) statement() *AssignmentNode {
	if p.curToken._type == STRING {
		return p.assignmentStatement()
	}
	return nil
}

func (p *Parser) statementList() *StatementListNode {
	statementList := StatementListNode{}
	statementList.nodes = []*AssignmentNode{p.statement()}
	for p.curToken._type == NEWLINE {
		p.eat(NEWLINE)
		statementList.nodes = append(statementList.nodes, p.statement())
	}
	return &statementList
}
