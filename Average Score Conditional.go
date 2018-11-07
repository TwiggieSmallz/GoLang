package main
  import "fmt"
    func main () {
    //initialize an array of test scores
    scores := [8]float64 {87.3, 92.0, 100, 78.9, 84, 98, 72.6, 64} //the # in the [] indicates the number of indices in the array; remember first value (87.3) == [0] - all arrays in Go are zero based.
    //print the array:
      fmt.Printf("Scores:%v\n", scores)
    //calculate the average:
      avg := 0.0 //inferred to be float64
    //add the scores:
    for i := 0; i < 8; i++ {
      avg += scores [i]
      }
    //divide the sum by total number of scores listed (8)
      avg = avg / 8
    //print the average:
      fmt.Printf("Average score: %.2f\n", avg) //%v prints entire value, %.Xf prints indicated number of decimal places to follow the float (indicated by the 'f')
    }
//Average score should be 84.60, if you get another value, recheck code.
//Arrays indexes begin at 0, so 0 through 7 == 8
//As a general rule, answers should specify to one more decimal point than what's needed.
