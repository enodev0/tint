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
		strings.Contains(t.expression, "sqrt") ||
		strings.Contains(t.expression, "exp") {
		return true
	} else {
		return false
	}
}

func timeTrack(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}

func (t TrapezoidIntegrator) Run() float64 {

	defer timeTrack(time.Now())
	delta := (t.upperLimit - t.lowerLimit) / (float64(t.steps))
	x   := make([]float64, t.steps+1)
	f_x := make([]float64, t.steps+1)


	for i := 0; i <= t.steps; i++ {
		if i > 0 {
			x[i] = x[i-1] + delta
		} else {
			x[i] = t.lowerLimit
		}
	}


	if t.stringHasFunctions() {
	//TODO: Re-implement the expression evaluator
		functions := map[string]govaluate.ExpressionFunction {
			"log": func(arg ...interface{}) (interface{}, error) {
				return math.Log10(arg[0].(float64)), nil
			},
			"ln": func(arg ...interface{}) (interface{}, error) {
				return math.Log(arg[0].(float64)), nil
			},
			"sin": func(arg ...interface{}) (interface{}, error) {
				return math.Sin(arg[0].(float64)*RAD), nil
			},
			"cos": func(arg ...interface{}) (interface{}, error) {
				return math.Cos(arg[0].(float64)*RAD), nil
			},
			"tan": func(arg ...interface{}) (interface{}, error) {
				return math.Tan(arg[0].(float64)*RAD), nil
			},
			"asin": func(arg ...interface{}) (interface{}, error) {
				return math.Asin(arg[0].(float64)*DEG), nil
			},
			"acos": func(arg ...interface{}) (interface{}, error) {
				return math.Acos(arg[0].(float64)*DEG), nil
			},
			"atan": func(arg ...interface{}) (interface{}, error) {
				return math.Atan(arg[0].(float64)*DEG), nil
			},
			"sinh": func(arg ...interface{}) (interface{}, error) {
				return math.Sinh(arg[0].(float64)), nil
			},
			"cosh": func(arg ...interface{}) (interface{}, error) {
				return math.Cosh(arg[0].(float64)), nil
			},
			"tanh": func(arg ...interface{}) (interface{}, error) {
				return math.Tanh(arg[0].(float64)), nil
			},
			"sqrt": func(arg ...interface{}) (interface{}, error) {
				return math.Sqrt(arg[0].(float64)), nil
			},
			"exp": func(arg ...interface{}) (interface{}, error) {
				return math.Exp(arg[0].(float64)), nil
			},
		}

		fn, _ := govaluate.NewEvaluableExpressionWithFunctions(t.expression, functions)
		parameter := make(map[string]interface{}, 1)
		for j := 0; j <= t.steps; j++ {
			parameter["x"] = x[j]
			temp, _ := fn.Evaluate(parameter)
			f_x[j] = temp.(float64)
		}
	} else {
		exp, _ := govaluate.NewEvaluableExpression(t.expression)
		parameter := make(map[string]interface{}, 1)
		for j := 0; j <= t.steps; j++ {

			parameter["x"] = x[j]

			temp, err := exp.Evaluate(parameter)
			if err != nil {
				panic("Panic: FATAL: Expression evaluation error")
			}

			f_x[j] = temp.(float64)
		}
	}

	sum := 0.0
	for index, elem := range f_x {
		if index == 0 || index == t.steps {
			sum = sum + elem
		} else {
			sum = sum + (2 * elem)
		}
	}
	return (delta / 2) * sum
}
