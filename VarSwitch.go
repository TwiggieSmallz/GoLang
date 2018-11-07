/*Switching on a Variable
One of the unique things about Go is that switching can be
performed "anywhere" - on strings, vars, ints, floats, etc. Other
languages like C or Jave, only allow switching on an integer or scalar number*/

package main
  import "fmt"
    func main () {
      numSize := ""                   //in-line declaration var set equal to empty string (to be populated later)
        for i := 1; i <= 10; i++ {    //loop from 1 to 10
          if i < 4 {                  //assigning numSize to i
            numSize = "Small"
          } else if i < 8 {
            numSize = "Medium"
          } else {
            numSize = "Large"
          }
          switch numSize {
            case "Small":
              fmt.Printf("%v is a small-sized number\n", i)
            case "Medium":
              fmt.Printf("%v is a medium-sized number\n", i)
            case "Large":
              fmt.Printf("%v is a large-sized number\n", i)
            }
          }
    }
