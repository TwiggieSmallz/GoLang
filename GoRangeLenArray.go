/*Range & Len of an Array*/     //len - length of an arrays index
/*In Go, an array of unknown length cannot be specified/declared. An arrays length is part of the declaration.
Arrays are used only for groups of data that have a known amount of initial elements; in all other cases a "slice" is used.*/
package main
  import "fmt"
    func main (){
    scores := [8]float64{87.3, 92.0, 100, 78.9, 84, 98, 72.6, 64} //The index operator and parameter ([] and 8) are essential to index declaration; Go does not allow for an index to be declared with empty parameters.
      fmt.Printf("Scores:%v\n", scores)
        avg := 0.0
          for i := range scores {               //range is a keyword in Go and is thought of as "for" each index in the 'scores' array.
          avg += scores[i]
              }                                 //len is a tool built into Go that tells the length of an arrays index.
        avg = avg / float64(len(scores))        //must specify float64 here else infers an int and operations of this nature must be performed on same var type.
          fmt.Printf("Average Score:%.2f\n", avg)
        }
//It is possible in Go to "cast" to different var types but this must always be an "up-cast" such that whatever we're casting too is bigger than what we cast from.
//Example: float64 > int so possible to cast-up 'int' type and transform it to a float.
