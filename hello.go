/* In Go, any source code that has a function is allocated to a container which is referred to as a package.

Each package within the source code is defined with a name, like variable assignment.

To access the different packages functionalities, we have to import them from the libraries.

Once imported, we can dictate functionality parameters to suite our purposes.

Example:*/

package main       /*This is the entry point to our program and contains all variables, functions, etc. from a global perspective. 'Main' is also the name of this created package and its contents.*/
  import "fmt"       /*This imports the fmt (pronounced "fumpt") library which implements formatted I/O functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.*/
    func main (){      /*The "main" function within the program; 'calls-up' to "package main" like calling a global var down into your function.*/
      fmt.Printf("Hello, world!\n")     /*Go's print to screen command; the '\n' dictates that a new-line is to follow after command execution.*/
    }                                     /*The end of the package.*/
