package main
  import "fmt"
    func main () {
      a := []int{2, 4, 6, 8}        //short-hand for making a slice with make([], x, y)
      b := []int{1, 3, 5, 7}
        a = append(a, 10, 12, 14)   //appending multi-values to a slice.
          fmt.Printf("a is now %v\n", a)
        a = append(a,b...)          //appending a slice to another slice. Since a and b's values are changing, the ellipses (...) is needed to denote a variadic parameter being used.
          fmt.Printf("a is now %v\n", a)
        a = append(a[:2], b[:2]...) //appending a slice of a slice to another slice of a slice (slice-ception); include all values up too but not including index value for 2.
          fmt.Printf("a is now %v\n", a)
    }
