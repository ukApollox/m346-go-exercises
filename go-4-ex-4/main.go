package main
import "fmt"

// TODO: implement the function convertCelsiusToFahrenheit
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
// TODO: implement the function convertFahrenheitToCelsius

func main() {
	// TODO: call the function convertCelsiusToFahrenheit
	fmt.Println("32 Grad Celsius sind", convertCelsiusToFahrenheit(32), "Grad Fahrenheit"); // 89.6
	// TODO: call the function convertFahrenheitToCelsius
	fmt.Println("89.6 Grad Fahrenheit sind", convertFahrenheitToCelsius(89.6), "Grad Celsius"); // 32
}
