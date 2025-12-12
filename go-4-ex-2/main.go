package main
import "fmt"
import "math"

// TODO: implement the function computeHypotenuse using math.Pow and math.Sqrt
func computeHypotenuse(a float64, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2));
}
func main() {
	// TODO: call the function computeHypotenuse
	fmt.Println("Die Hypotenuse ist", computeHypotenuse(3, 4)); // 5
}
