package main 
import (
	"fmt"
	"os"
	"github.com/go-parser/src/lexer"
)

func main (){
	bytes, _ := os.ReadFile("./examples/00.lang");
	source := string(bytes);

	fmt.Printf("Code: %s\n", source);

	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens{
		token.Debug()
	}

}