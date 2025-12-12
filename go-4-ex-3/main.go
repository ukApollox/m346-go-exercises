package main

import "fmt"
import "math"

func computeQuadraticFormula(a float64, b float64, c float64) []float64 {
    D := b*b - 4*a*c

    if D < 0 {
        return []float64{} 
    } else if D == 0 {
        x := -b / (2 * a)
        return []float64{x}
    } else {
        sqrtD := math.Sqrt(D)
        x1 := (-b + sqrtD) / (2 * a)
        x2 := (-b - sqrtD) / (2 * a)
        return []float64{x1, x2}
    }
}

func main() {
    solutions := computeQuadraticFormula(1, -3, 2)
    if len(solutions) == 0 {
        fmt.Println("Die quadratische Gleichung hat keine reellen Lösungen.")
    } else {
        fmt.Println("Die Lösungen der quadratischen Gleichung sind:", solutions)
    }

    // Testfälle
    fmt.Println(computeQuadraticFormula(1, 0, -4))  //-2 und 2
    fmt.Println(computeQuadraticFormula(1, 2, 1))   // -1
    fmt.Println(computeQuadraticFormula(1, 0, 1))   // keine Lösung
}