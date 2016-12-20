package main

import (
	"fmt"
	"flag"
	"github.com/enodev0/integrate/integral"
)

func main() {
	n := flag.Int("n", 10, "No. of iteration")
	ll := flag.Float64("ll", 0.0, "Lower limit")
	ul := flag.Float64("ul", 0.0, "Upper limit")
	exp := flag.String("fn", "", "Function to integrate")

	flag.Parse()
	// a := integral.NewIntegrator(20, -1, 1, "(1-(x**3))**(1/2)")
	a := integral.NewIntegrator(*n, *ll, *ul, *exp)
	b := a.Run()
	fmt.Println(b)
}
