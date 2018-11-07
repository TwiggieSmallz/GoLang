/*Structure Initialization and Method Syntax | Go*/

package main
  import (
  "fmt"
  "math"
  )     //Go doesn't have constructors or initializers but we can return a pointer to a structure to mimic this.
    type point struct {
      x,y float64
    }
      func makePoint (a,b float64) *point {
        p := new(point)
          p.x, p.y = a, b
            return p
      }
      func(p point) distanceTo(a, b float64) float64 {
        dist := math.Sqrt(math.Pow(a - p.x, 2) + math.Pow(b - p.y, 2))
          return dist
      }
      func main () {
        x,y := 0.0, 0.0
          fmt.Printf("Enter point coordinates using x,y format please: ")
          fmt.Scanf("%v, %v", &x, &y)
          myPoint := &point{x, y}        //myPoint = point value; myPoint is a pointer to a point; myPoint = &point{x, y}
          fmt.Printf("Point = (%v, %v)\n", myPoint.x, myPoint.y)
          fmt.Printf("Distance of given point from origin (0, 0) is %.2v\n", myPoint.distanceTo(0, 0))
      }     //call a function that returns a pointer to a point.
