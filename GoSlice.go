/*Slices in Go*/
/*Slice - A slice or slices in Go are, in technical terms, a reference to an underlying array*/
package main
  import "fmt"
    func main () {
      a := make([]int, 10, 100)        //make is a function built-in to Go. Within the parameters (), we see: an empty index box [], as this is the array being examined. Followed by its type, int. And lastly, the slice parameters - the first being the initial size of the slice and the second being the capacity of the slice/size of the array.
        i := 0
          for i < 10 {
            a[i] = 2 * i
              i++
          }
            fmt.Printf("a is %v\n", a)
            fmt.Printf("Adding an element to a...\n")
              a = append(a, 20)       //append is a function built-in to Go. Adds the 2nd parameter amount to the var specified in parameter 1 and produces a "new" slice.
                fmt.Printf("a is now %v\n", a)
              b := a[1:5]             //a "slice" of slice-a ranging from index 1's value up to, but not including, index 5's value.
                fmt.Printf("b is %v\n", b)
    }
