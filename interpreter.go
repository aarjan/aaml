package main

type Interpreter struct {
	parser Parser
	data   map[string][]string
}

// NewInterpreter ...
func NewInterpreter(p Parser) *Interpreter {
	return &Interpreter{p, make(map[string][]string)}
}

func (i *Interpreter) interpret() {
	list := i.parser.statementList()
	i.evaluate(list)
}

func (i *Interpreter) evaluate(node interface{}) []string {
	switch node.(type) {
	case *StatementListNode:
		for _, n := range node.(*StatementListNode).nodes {
			if n != nil {
				i.evaluate(n)
			}
		}
	case *AssignmentNode:
		ast := node.(*AssignmentNode)
		i.data[ast.left.token.value] = i.evaluate(ast.right)
	case ExprNode:
		ast := node.(ExprNode)
		if len(ast.nodes) == 1 {
			return i.evaluate(ast.nodes[0])
		}
		list := []string{}
		for _, n := range ast.nodes {
			list = append(list, i.evaluate(n)...)
		}
		return list
	case StringNode:
		ast := node.(StringNode)
		if ast.token == nil {
			return []string{"<nil>"}
		}
		return []string{ast.token.value}
	}
	return nil
}
