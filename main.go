package main

import (
	"fmt"
	"flag"
	"github.com/enodev0/tint/integral"
)

func main() {
	n := flag.Int("n", 10, "No. of iteration")
	ll := flag.Float64("ll", 0.0, "Lower limit")
	ul := flag.Float64("ul", 0.0, "Upper limit")
	exp := flag.String("fn", "", "Function to integrate")

	flag.Parse()

	fmt.Printf("\nIntegrating f(x) = %s from %0.2f to %0.2f in %d steps\n\n",
		*exp, *ll, *ul, *n)
	a := integral.NewIntegrator(*n, *ll, *ul, *exp)
	b := a.Run()
	fmt.Printf("Result: %0.3f unit sq\n\n", b)
}
