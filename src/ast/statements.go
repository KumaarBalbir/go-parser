package ast

// { statement 1; statement 2; }
/*


Here is a succinct explanation of the `BlockStmt` class definition:

* The `BlockStmt` struct represents a block of statements in the abstract syntax tree (AST).
* The `Body` field is a slice of `Stmt` objects, which contains the individual statements within the block.

Note that the `BlockStmt` struct does not have any methods, it only has a single field `Body` to store the statements.
*/
type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

/*
Here is a succinct explanation of the `ExpressionStmt` class definition:

* The `ExpressionStmt` struct represents an expression statement in the abstract syntax tree (AST).
* It has a single field `Expression` of type `Expr`, which stores the expression being evaluated.

The class method is:
* `stmt()`: This method is a placeholder for any implementation that wants to provide a way to execute or evaluate the expression statement.
*/
type ExpressionStmt struct {
	Expression Expr
}

// Implement the stmt method for ExpressionStmt
func (e ExpressionStmt) stmt() {}

/*
This class definition defines a `VarDeclStmt` struct in Go, which represents a variable declaration statement in an abstract syntax tree (AST). Here's a succinct explanation of what each field does:

* `VariableName`: stores the name of the variable being declared.
* `IsConstant`: indicates whether the variable is declared as a constant.
* `AssignedValue`: stores the value assigned to the variable, if any.
* `ExplicitType`: stores the explicit type of the variable, if any.

stmt() method:
This method takes a VarDeclStmt receiver (v) and returns no value (i.e., it's a void function).
The method is currently empty, so it doesn't perform any actions or return any values.
*/
type VarDeclStmt struct {
	VariableName  string
	IsConstant    bool
	AssignedValue Expr
	ExplicitType  Type
}

func (v VarDeclStmt) stmt() {}
