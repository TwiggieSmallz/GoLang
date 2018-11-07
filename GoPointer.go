package main
  import "fmt"
    func main () {
      a := new(int)                                     //a is a pointer to an int
      *a = 9                                           //assign 9 as "contents of" a
        fmt.Printf("a is %v, *a is %v\n", a, *a)      //print address (in memory) and the contents of a
      b := a                                          //b points to the same location of a
        fmt.Printf("%v = b, %v =*b\n", b, *b)         //print address (in memory) and the contents of b
      *a = *a + 1                                     //increment a's contents
        fmt.Printf("%v is b, %v is *b\n", b, *b)      //print adress and new incremented contents of ab
    }
/*Because b references the same memory as a, changing a's
contents also changes b's contents.*/
/*Assigning a reference-type is prohibited in Go*/
