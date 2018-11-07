/*Buffered String Read | Go-Lang*/
/*Getting input from stdin (Standard Input)*/

package main
  import (
    "bufio"                                         //means buffer input/output
    "fmt"
    "os"
  )
      func main () {                                //Remember caps are needed to export functionality for use in other packages.
        reader := bufio.NewReader(os.Stdin)         //file-reader functionality
          fmt.Printf("Enter a string: ")
            str, err := reader.ReadString('\n')
              if err != nil {
                panic("Error: new line not found.")
              }
                fmt.Println(str)                    //Println = Print on new line.
      }
/*Recap - these other pkgs/libs (bufio/os) bring different functions into the program.
Bufio (in this case) brings a reader utility so the program can "take-in" information.
OS brings utilities as well to modify system properties/files. In this case: enter some
string, program reads supplied info, decides who/where to pass to, and the new lib functions
perform desired operations.*/
