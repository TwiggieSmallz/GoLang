package main
  import "fmt"
    func main () {
      a,b := 3,5
        fmt.Printf("Before swap: a=%v, b=%v\n", a, b)
//swap a and b:
      a,b = b,a
      b,a = swap(a,b)
        fmt.Printf("After swap: a=%v, b=%v\n", a, b)
    }
    func swap(a,b int) (int, int) {       /*the 2 enclosed ints at the end of the function declaration are the desired return types; if multi values, need to define type and set parameters.*/
    return b, a
    }
