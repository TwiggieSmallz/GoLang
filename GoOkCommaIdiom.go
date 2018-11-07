/*Ok Comma Idiom - Maps Continued | Go-Lang*/
/*Go has a property called the "OK Comma Idiom" which tests for the
presence of a key/value pair in a given map.*/

package main
  import "fmt"
    func main () {
      digits := map[string]int {        //digits is a var of type map with string as keys and int as values.
        "zero" : 0,                      //entries to the left of the colon (:) are the keys in the form of strings.
        "one" : 1,                       //entries to the right of the colon (:) are the values in the form of ints.
        "two" : 2,                      //each entry to be followed by a comma, even last entry.
        "three" : 3,
        "four" : 4,
                                }
            if digits["four"] != 0 {    //change str value inbetween "" to change output.
              fmt.Printf("digit value is %v\n", digits["four"])
            } else {
                fmt.Println("digit not in map\n")
            }

  key := "three"

var digit int
var ok bool

        if digit, ok = digits[key]; ok {    //note syntax used here.
          fmt.Printf("%v is %v\n", key, digit)
        } else {
          fmt.Printf("%v not in map.\n", key)
        }
          if _, ok = digits["nine"]; ok {
            fmt.Printf("9 is in the map!\n")
          } else {
            fmt.Printf("9 is not in the map.\n")
          }
        //delete(digits, "zero")      //delete k/v from map; uncomment and go build to test.
          fmt.Printf("%v\n", digits) //prints the entire map
    }
