package main

import (
	"image/color"
	"testing"
)

func benchmarkMadelbrot(b *testing.B, f func(complex128) color.Color) {
	for i := 0; i < b.N; i++ {
		f(complex(float64(i), float64(i)))
	}
}

func BenchmarkMadelbrotComplex128(b *testing.B) {
	benchmarkMadelbrot(b, mandelbrot)
}

func BenchmarkMadelbrotComplex64(b *testing.B) {
	benchmarkMadelbrot(b, mandelbrot64)
}

func BenchmarkMadelbrotBigFloat(b *testing.B) {
	benchmarkMadelbrot(b, mandelbrotBigFloat)
}

func BenchmarkMadelbrotRat(b *testing.B) {
	benchmarkMadelbrot(b, mandelbrotRat)
}
