package ast 


// { statement 1; statement 2; }
type BlockStmt struct {
	Body []Stmt
}

func (n BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
} 