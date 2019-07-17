// version 1.0 2018-07-16

// Package tempconv performs Celsius and Fahrenheit conversions.
package tempconv

import (
	"flag"
	"fmt"
)

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

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%gK", k)
}
