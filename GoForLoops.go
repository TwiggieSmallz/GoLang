/*For Loops | Go-Lang*/
/*Go only uses "For" loops but they can be in different forms.
  One form of the for loop has 3 sections in its conditional expression:
    1. Setup a var to be incrimented
      for i := 0;
    2. The second conditional is for how to treat the loop; "While this is true..."
      for i := 0; i < 8;
    3. The third conditional incriments the var (at least in this example)
      for i := 0; i < 8; i++ */

/*The 2nd form of the for loop is similar to a "while" loop:*/
package main
  import "fmt"
    func main () {
    scores := [8]float64{87.3, 92.0, 100, 78.9, 84, 98, 72.6, 64}
      fmt.Printf("Scores: %v\n", scores)
      avg := 0.0
        i := 0
          for i < 8 {       //"While" i is less than 8 do...
            avg += scores[i]
              i++     //Could alternately be written as i+1 or i+=1 | Loop exit here: eventually "i" will be incremented to 9 making the condition 'false'.
          }
          avg = avg / 8
            fmt.Printf("Average score: %.2f\n", avg)
    }
