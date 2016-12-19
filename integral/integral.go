package integral

import "github.com/Knetic/govaluate"

type TrapezoidIntegrator struct {
	steps      int
	expression string
	lowerLimit float64
	upperLimit float64
}

func NewIntegrator(stp int, ll, ul float64, exp string) TrapezoidIntegrator {
	return TrapezoidIntegrator{
		steps:      stp,
		expression: exp,
		lowerLimit: ll,
		upperLimit: ul,
	}
}

func (t TrapezoidIntegrator) Run() float64 { // return an error?
	delta := (t.upperLimit - t.lowerLimit) / (float64(t.steps))
	x := make([]float64, t.steps)
	f_x := make([]float64, t.steps)
	for i := 0; i <= t.steps; i++ {
		if i > 0 {
			x[i] = x[i-1] + delta
		} else {
			x[i] = t.lowerLimit
		}
	}
	//if 
	exp, _ := govaluate.NewEvaluableExpression(t.expression) 
	//err != nil {
	//	panic("Panic: FATAL: Expression could not be generated")
	//}
	// allocating only one, since we would be re-assigning to the same parameter
	parameter := make(map[string]interface{}, 1) // govaluate mandate
	var err error
	for j := 0; j <= t.steps; j++ {
		//FIXME: The variable name must be the same as that in the expression.
		// Here, I'm choosing "x" since it is the most common.
		// Thus, integral of 1-x**1/2 will work, but 1-y**1/2 won't work,
		// since the expectant variable name is hardcoded in. Refer to the README
		// of the govaluate package for more information. Hence the cumbersome
		// re-assignments, rather than all at a time.
		parameter["x"] = x[j] // here we go.
		temp, err:= exp.Evaluate(parameter)
		// the return type of exp.Evaluate() is interface{}, not float64.
		if err != nil { //FIXME: Need type assertions. Whats that?
			panic("Panic: FATAL: Expression evaluation error")
		}
		f_x[j] = float64(temp)
	}
	// calculate the sum
	sum := 0.0
	for index, elem := range f_x {
		if index == 0 || index == t.steps { // see calculation on sheet
			sum = sum + elem
		} else {
			sum = sum + (2 * elem)
		}			
	}
	return (delta/2) * sum // the integral
}
