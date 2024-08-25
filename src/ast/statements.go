package ast 


// { statement 1; statement 2; }
type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
} 

// Implement the stmt method for ExpressionStmt
func (e ExpressionStmt) stmt() {}