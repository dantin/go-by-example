// version 1.0 2018-07-24

package main

import "fmt"

// Flags represents network device status.
type Flags uint

const (
	// FlagUp is up
	FlagUp Flags = 1 << iota
	// FlagBroadcast supports broadcast access capability
	FlagBroadcast
	// FlagLoopback is a loopback interface
	FlagLoopback
	// FlagPointToPoint belongs to a point-to-point link
	FlagPointToPoint
	// FlagMulticast supports multicast access capability
	FlagMulticast
)

// IsUp
func IsUp(v Flags) bool {
	return v&FlagUp == FlagUp
}

func TurnDown(v *Flags) {
	*v &^= FlagUp
}

func SetBroadcast(v *Flags) {
	*v |= FlagBroadcast
}

func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}
