package main
import "fmt"

// TODO: implement the function computeGrade
func computeGrade(gotPoints int, maxPoints int) float64 {
	grade := float64(gotPoints) / float64(maxPoints) * 5 + 1
	return grade
	
}

func main() {
	// TODO: call the function computeGrade
		fmt.Println("Deine Note ist", computeGrade(17, 28)); // 4.125
}
