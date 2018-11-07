package main
  import "fmt"
    func main() {                   /*note the empty parameters (); always required, even if empty*/
    a,b := 3,5
    fmt.Printf("%v + %v = %v\n", a, b, a + b)
    }

/* function name/ability = add, function parameters () = n1, n2 int, followed by 'int' which is the desired return TYPE. */
  func add(n1, n2 int) int {
  return n1 + n2
  }
