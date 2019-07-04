// version 1.0 2018-07-16

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import "fmt"

// Celsius declaration defines a new named type that has the same underlying type as a float64.
type Celsius float64

// Fahrenheit declaration defines a new named type that has the same underlying type as a float64.
type Fahrenheit float64

// Kelvin declaration defines a new named type that has the same underlying type as a float64.
type Kelvin float64

const (
	// AbsoluteZeroC is the absolute zero temperature in celsius.
	AbsoluteZeroC Celsius = -273.15
	// FreezingC is the freezing point temperature in celsius.
	FreezingC Celsius = 0
	// BoilingC is the boiling point temperature in celsius.
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
