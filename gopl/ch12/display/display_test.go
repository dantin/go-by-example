package display

import (
	"io"
	"os"
	"testing"

	"github.com/dantin/go-by-example/gopl/ch7/eval"
)

func TestExampleExpr(t *testing.T) {
	e, _ := eval.Parse("sqrt(A / pi)")
	Display("e", e)
}

func TestSlice(t *testing.T) {
	Display("slice", []*int{new(int), nil})
}

func TestNilInterface(t *testing.T) {
	var w io.Writer
	Display("w", w)
}

func TestPtrToInterface(t *testing.T) {
	var w io.Writer
	Display("&w", &w)
}

func TestStruct(t *testing.T) {
	Display("x", struct{ x interface{} }{3})
}

func TestInterface(t *testing.T) {
	var i interface{} = 3
	Display("i", i)
}

func TestPtrToInterface2(t *testing.T) {
	var i interface{} = 3
	Display("&i", &i)
}

func TestArray(t *testing.T) {
	Display("x", [1]interface{}{3})
}

func TestMovie(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Color           bool
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}

	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mardrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	Display("strangelove", strangelove)
}

// This test ensures that the program terminates without crashing.
func TestDemo(t *testing.T) {

	Display("os.Stderr", os.Stderr)
}
