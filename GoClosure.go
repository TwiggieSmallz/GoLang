/*Closure Functions | Go*/

package main
  import "fmt"
    func main () {
      var a, b int
        add := func() int {
          return a + b
        }
        a, b = 3, 9
          fmt.Printf("%v + %v = %v\n", a, b, add())
        a, b = 4, -7
          fmt.Printf("%v + %v = %v\n", a, b, add())
    }
/*Functions defined by their types are called "closures" as they can access vars not defined in their scope.
Closures 'close-over' variables in the scope of which they're defined.*/
