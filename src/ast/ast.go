package ast 

type Stmt interface { // statement
	stmt ()
}

type Expr interface { // expression
	expr ()
}