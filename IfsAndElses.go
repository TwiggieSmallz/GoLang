/*Decision Making: If, Else if, and Else*/
package main
  import "fmt"
    func main () {
      a,b := 2,1
        if a == b {
          fmt.Printf("%v and %v are equal\n", a, b)
        } else if a > b { fmt.Printf("%v is greater than %v\n", a, b) //when using else statements, else's must be on the same line as the closing curly-brace from the if/else before it.
        } else {        //again, notice the else sandwhiched between a closing brace on the left and an opening brace on the right | } else {
          fmt.Printf("%v less than %v", a, b)
                }
    }
/*The lexar for Go is adding ;'s by inference and so if it sees a closing brace without an else on the same line, it assumes };
which 'blocks' the parameters of the statement from being seen by the else and so it looks empty and gives a syntax error.*/
