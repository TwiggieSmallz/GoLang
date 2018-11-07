/*Defer | Go*/
/*Defer will delay the execution of something until after its container function is finished.*/

package main
  import "fmt"
    func main () {
      defer f1()      //function f1 to be deferred until function f2 completes.
      f2()
    }
    func f1(){
      fmt.Println("Hello from f1!\n")
    }
    func f2() {
      fmt.Println("Hello from f2!")
        for i := 1; i <= 3; i++ {
          fmt.Println(i)
    }
}
/*
Output of this package should be:
Hello from f2!
1
2
3
Hello from f1!
*/
