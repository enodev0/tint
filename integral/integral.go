package integral

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"math"
	"strings"
	"time"
)

const (
	DEG = (180.0 / math.Pi) //rad -> deg
	RAD = (math.Pi / 180.0) //deg -> rad
	/*RED    = "\033[31m"
	GREEN  = "\033[32m"
	YELLOW = "\033[33m"
	RESET  = "\033[m"*/
)

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

func (t TrapezoidIntegrator) stringHasFunctions() bool {
	if strings.Contains(t.expression, "log") ||
		strings.Contains(t.expression, "ln") ||
		strings.Contains(t.expression, "sin") ||
		strings.Contains(t.expression, "cos") ||
		strings.Contains(t.expression, "tan") ||
		strings.Contains(t.expression, "asin") ||
		strings.Contains(t.expression, "acos") ||
		strings.Contains(t.expression, "atan") ||
		strings.Contains(t.expression, "sinh") ||
		strings.Contains(t.expression, "cosh") ||
		strings.Contains(t.expression, "tanh") ||
		strings.Contains(t.expression, "deg") ||
		strings.Contains(t.expression, "sqrt") ||
		strings.Contains(t.expression, "rad") ||
		strings.Contains(t.expression, "exp") {
		return true
	} else {
		return false
	}
}

// Tracktime courtesy stathat
func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}

func (t TrapezoidIntegrator) Run() float64 { // return an error?
	defer timeTrack(time.Now())
	delta := (t.upperLimit - t.lowerLimit) / (float64(t.steps))
	x := make([]float64, t.steps+1)
	f_x := make([]float64, t.steps+1)
	for i := 0; i <= t.steps; i++ {
		//println(i, x[i])
		if i > 0 {
			x[i] = x[i-1] + delta
		} else {
			x[i] = t.lowerLimit
		}
	}

	if t.stringHasFunctions() {
	
		functions := map[string]govaluate.ExpressionFunction{
			"log": func(arg float64) float64 {
				return math.Log10(arg)
			},
			"ln": func(arg float64) float64 {
				return math.Log(arg)
			},
			"sin": func(arg float64) float64 {
				return math.Sin(arg)
			},
			"cos": func(arg float64) float64 {
				return math.Cos(arg)
			},
			"tan": func(arg float64) float64 {
				return math.Tan(arg)
			},
			"asin": func(arg float64) float64 {
				return math.Asin(arg)
			},
			"acos": func(arg float64) float64 {
				return math.Acos(arg)
			},
			"atan": func(arg float64) float64 {
				return math.Atan(arg)
			},
			"sinh": func(arg float64) float64 {
				return math.Sinh(arg)
			},
			"cosh": func(arg float64) float64 {
				return math.Cosh(arg)
			},
			"tanh": func(arg float64) float64 {
				return math.Tanh(arg)
			},
			"sqrt": func(arg float64) float64 {
				return math.Sqrt(arg)
			},
			"deg": func(arg float64) float64 {
				return arg * DEG
			},
			"rad": func(arg float64) float64 {
				return arg * RAD
			},
			"exp": func(arg float64) float64 {
				return math.Exp(arg)
			},
		}

		fn, _ := govaluate.NewEvaluableExpressionWithFunctions(t.expression, functions)
		parameter := make(map[string]interface{}, 1)
		for j := 0; j <= t.steps; j++ {
			parameter["x"] = x[j]
			temp, _ := exp.Evaluate(parameter)
			f_x[j] = temp.(float64)
		}
	} else {
		exp, _ := govaluate.NewEvaluableExpression(t.expression)
		// allocating only one, since we would be
		// re-assigning to the same parameter
		parameter := make(map[string]interface{}, 1) // govaluate mandate
		// var err error
		for j := 0; j <= t.steps; j++ {
			//FIXME: The variable name must be the same as
			// that in the expression.
			// Here, I'm choosing "x" since it is the most common.
			// Thus, integral of 1-x**1/2 will work,
			//but 1-y**1/2 won't work,
			// since the expectant variable name is hardcoded in.
			// Refer to the README
			// of the govaluate package for more information.
			parameter["x"] = x[j] // here we go.
			temp, err := exp.Evaluate(parameter)
			// the return type of exp.Evaluate() is interface{},
			// not float64.
			if err != nil { //FIXME: Need type assertions. Whats that?
				panic("Panic: FATAL: Expression evaluation error")
			}
			// What's the concrete type implementing the interface? Below.
			f_x[j] = temp.(float64) // type assert interface{} to float
		}
	}
	// calculate the sum
	sum := 0.0
	for index, elem := range f_x {
		if index == 0 || index == t.steps {
			sum = sum + elem
		} else {
			sum = sum + (2 * elem)
		}
	}
	return (delta / 2) * sum // the integral
}
