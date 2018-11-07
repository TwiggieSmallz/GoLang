package main
  import "fmt"
    func main () {
      num := 5
      negate(&num)     //negate is a built-in operator in Go which makes an int value negative.
        fmt.Printf("num is now %v\n", num)
      }
      func negate (a *int) {
        *a = *a * -1
      }
