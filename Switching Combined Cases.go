/*Switching with combined cases in Go*/
package main
  import "fmt"
    func main () {
    for i := 1; i <= 10; i++ {
      switch i {
        case 1, 2, 3:
          fmt.Printf("%v is a small number\n", i)
        case 4, 5, 6, 7:
          fmt.Printf("%v is a medium number\n", i)
        case 8, 9, 10:
          fmt.Printf("%v is a large number\n", i)
              }
          }
    }
/*This package actually has the exact same outcome as the "Switching on a Variable"
package in the decision making=>switching subfolder but is cleaner, easier to read, and has better simplicity.*/
