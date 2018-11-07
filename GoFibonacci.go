/*Fibonacci Closure Demo | Go*/
//The Fibonacci sequence is where current number is the sum of the 2 preceeding numbers.

package main
  import "fmt"
    func main () {
      nextFib := fib()
        fmt.Printf("1 1 ")
          for i := 0; i < 10; i++ {
            fmt.Printf("%v\n", nextFib())
          }
    }
    func fib() func() uint {
    var prevf uint = 1
    var f uint = 1
      return func() uint {
        f, prevf = f + prevf, f
          return f
      }
    }
