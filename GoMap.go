/*Maps in Go*/
/*A map is a table of key/value pairs. Maps are also known as hash tables,
dictionaries, or associative arrays in other programming languages.*/

/*The keys of a map MUST be "hashable", meaning that a unique value
can be assigned to each key and that equality can be evaluated for every key*/

/*Ints, floats, strings, etc. are all hashable types but arrays, slices, etc. are NOT.
The values of a map need not be hashable.*/

package main
  import "fmt"
    func main () {
      digits := map[string]int {        //digits is a var of type map with string as keys and int as values.
        "zero" : 0,                      //entries to the left of the colon (:) are the keys in the form of strings.
        "one" : 1,                       //entries to the right of the colon (:) are the values in the form of ints.
        "two" : 2,                      //each entry to be followed by a comma, even last entry.
        "three" : 3,
        "four" : 4,
        "five" : 5,
                                }
          digits["five"] = 5
            if digits["zero"] != 0 {    //change str value inbetween "" to change output.
              fmt.Printf("digit value is %v\n", digits["four"])
            } else {
                fmt.Printf("digit not in map")
            }
    }
