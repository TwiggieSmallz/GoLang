/*Scan & ScanLn in Go*/
//to use scan, pointers must be used.

package main
  import "fmt"
    func main () {
      word1, word2 := "", ""
        fmt.Printf("Enter two words please: ")
          count, err := fmt.Scan(&word1, &word2)    //remember & is a pointer operator.
            if err != nil {
              panic(err)
            }
            fmt.Printf("You put %v. There are %v words.\n", word1 + " " + word2, count)
    }
/*Scan doesn't mind new lines and treats them as a space char. ScanLn (Scan Line)
DOES care and will error if a new line happens before it finishes it arguements.*/

/*In the event of a 'troll/malicious' user, be extremely careful using scan & scanln - if a user
intentionally overflows desired # of things to be entered at the prompt, it'll imediatelly grab
the first char that's over the limit, "eat it", and then DROP AND RUN whatever immediatelly followed 1st char!
For example: if when asked to enter 2 words the user puts "word#1, free dpwd" the package will take word#1,
then will take "free" then grab 'd' from the 3rd entry and basically discard it, and it will drop the rest
"pwd" onto a command line and run it which then prints the working directory immediately after the package
finishes its arguements; now imagine a different command than PWD being forceably ran immediately after program...*/
