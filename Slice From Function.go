/*Returning a slice from a Function*/
package main
  import "fmt"
    func main () {
      testString := "This is a test."       //a string is an array (or a slice) of bytes:
        fmt.Printf("The 1st word of the string is \"%s\" \n", firstWord(testString))       //use forward slashes to allow for additional quotes within already-quoted parameters.
    }
    func firstWord(str string) []byte {
      word := []byte{}
        for ch := range str {
          if str[ch] == ' '{        //space betwixt the '' is how to place the "space character" into Go.
          break
          } else {
            word = append(word, str[ch])
          }
        }
              return word
    }
