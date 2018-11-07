/*In Go, a 'switch' is a sort of replacement for if/else stacks*/
package main
  import "fmt"
    func main () {
      grades := [7]int{93, 72, 68, 54, 100, 43, 86}
      strScore := ""        //In Go, in order to have a var be exportable for use in other packages, the name must contain a capital character.
        for score := range grades {     //read as "'for' the score var within this range of grades...
          switch {                      //Could be an if statement here followed by else ifs for each score case range
            case grades [score] > 90:
              strScore = "A"
            case grades [score] > 80:
              strScore = "B"
            case grades [score] > 70:
              strScore = "C"
            case grades [score] > 60:
              strScore = "D"
            default:        //will catch any other values not defined; in this case, any number less than 60.
              strScore = "F"
                }
                  fmt.Printf("%v, letter grade %v\n", grades[score], strScore)
              }
    }
