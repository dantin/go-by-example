// version 1.0 2018-07-16

// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

// Celsius declaration defines a new named type that has the same underlying type as a float64.
type Celsius float64

// Fahrenheit declaration defines a new named type that has the same underlying type as a float64.
type Fahrenheit float64

const (
	// AbsoluteZeroC is the absolute zero temperature in celsius.
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing point temperature in celsius.
	FreezingC Celsius = 0
	// BoilingC is the boiling point temperature in celsius.
	BoilingC Celsius = 100
)

// CToF convert Celsius to Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC convert Fahrenheit to Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
