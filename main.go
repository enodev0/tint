package main

import "fmt"
import "github.com/enodev0/integrate/integral"

const DEFAULT_STEPS = 20

func main() {
	a := integral.NewIntegrator(DEFAULT_STEPS, -1, 1, "(1-(x**3))**(1/2)")
	b := a.Run()
	fmt.Println(b)
}
