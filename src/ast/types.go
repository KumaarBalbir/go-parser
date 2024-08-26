package ast

/*
Here is a succinct explanation of the `SymbolType` class definition:

* The `SymbolType` struct represents a symbol type in the abstract syntax tree (AST).
* It has one field:
  - `Name`: stores the name of the symbol as a string.

Note that this struct does not have any methods, it only has a single field to store the name of the symbol.
*/
type SymbolType struct {
	Name string // T name of the symbol
}

/*
This is a method declaration in Go. The method `_type` is defined on the `SymbolType` struct. The method is empty, meaning it does not perform any actions or return any values. It is likely a placeholder for future implementation or a way to satisfy an interface.
*/
func (t SymbolType) _type() {

}

/*
This class definition defines a struct called `ArrayType` in Go, which represents an array type in the abstract syntax tree (AST). Here's a succinct explanation of what each field does:

* `Underlying Type`: stores the underlying type of the array, which can be any type that implements the `Type` interface.
*/
type ArrayType struct {
	Underlying Type // []T
}

func (t ArrayType) _type() {}
