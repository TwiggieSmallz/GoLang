/*Return multi-values from Functions */
package main
  import "fmt"
    func main () {
    a,b := 3,5
      fmt.Printf("Before swap: a=%v, b=%v\n", a, b)
//swap a and b
        tmp := a      /*<= temporary variable*/
        a=b
        b=tmp
      fmt.Printf("After swap a=%v, b=%v\n", a, b)
    }
