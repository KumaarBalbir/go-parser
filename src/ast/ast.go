package ast

/*
This interface defines a contract for any type that wants to represent a statement in the abstract syntax tree (AST).
The stmt() method is a placeholder for any implementation that wants to provide a way to execute or evaluate the statement.
*/
type Stmt interface { // statement
	stmt()
}

/*
It defines an interface called Expr that represents an expression in the abstract syntax tree (AST).
The interface has one method:
expr(): This method is a placeholder for any implementation that wants to provide a way to evaluate or execute the expression.
*/
type Expr interface { // expression
	expr()
}

/*
This class definition defines an interface called Type in Go. Here's a succinct explanation of what the class method does:

_type(): This method is a placeholder for any implementation that wants to provide a way to represent or handle a type in the abstract syntax tree (AST). The exact behavior is left to the implementing classes.
*/
type Type interface {
	_type()
}
