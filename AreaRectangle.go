/*Structures in Go | Rectangle*/
package main
  import "fmt"           //Structure names should always be capitalized; anything with a capital letter in a package will be exported to work in other packages.
    type Rect struct{   //This structure is of type: rectangle
      w,h float64       //with width and height values that are floats.
    }                   /*A pointer is an address in memory that can be called to; the indirection operator (*) is used to establish the route.*/
      func main () {
    var r1 Rect         //r1 is of type Rect
      r1.w = 5.3
      r1.h = 2.7
        fmt.Printf("Area of r1 = %.2f\n", area(r1))
      r2 := new(Rect)     //"new" is a built-in operator for new pointer production. Allocates memory to store values and sets those vals to 0 unless user defines.
      r2.w = 3.5
      r2.h = 7.1
        fmt.Printf("Area of r2 = %.2f\n", area(*r2))      //the asterisk in the area(*r2) is known as the indirection operator in Go.
    }
      func area (r Rect) float64 {
        return r.w * r.h
      }

/*What is a pointer [in Go]?
A pointer is a variable that holds the address, in memory, of a value.
Each byte in memory is assigned an address.
A pointer holds the address information for a memory location.
Two operators are used:
  * - The indirection or "content of" operator.
  & - The "address of" operator.*/

/*We can make a pointer to a variable key by taking its address using a & operator:
  a := 5 //a is an int
  b := &a //b is a pointer to an int (a)
  a == *b (both are 5)
So why do pointers matter?
  When we use 'new' to allocate space for a variable, we get a reference (pointer) to that variable
  We can pass a pointer as a parameter to a function in order to change that parameter's value outside of the function.
All languages deal with/have/use pointers; some languages hide that fact very well!*/
