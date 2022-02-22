package govaluate

import (
	"github.com/Knetic/govaluate"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestGovaluate(t *testing.T) {
	expr, _ := govaluate.NewEvaluableExpression("1+1")
	result, err := expr.Evaluate(nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestGovaluateParams(t *testing.T) {
	expr, _ := govaluate.NewEvaluableExpression("a + b")
	parameters := make(map[string]interface{})
	parameters["a"] = 1
	parameters["b"] = 2
	result, _ := expr.Evaluate(parameters)
	assert.Equal(t, float64(3), result)

	parameters["a"] = 10
	parameters["b"] = 20
	result, _ = expr.Evaluate(parameters)
	assert.Equal(t, float64(30), result)
}

func TestGovaluateDivZero(t *testing.T) {
	expr, _ := govaluate.NewEvaluableExpression("1/a")
	result, err := expr.Evaluate(map[string]interface{}{"a": 0})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, true, math.IsInf(result.(float64), 0))
}

func TestGovaluateFunction(t *testing.T) {
	functions := map[string]govaluate.ExpressionFunction{
		"sum": func(args ...interface{}) (interface{}, error) {
			var r float64
			for _, arg := range args {
				r += arg.(float64)
			}
			return r, nil
		},
	}

	exprString := "sum(1,2)"
	expr, _ := govaluate.NewEvaluableExpressionWithFunctions(exprString, functions)
	result, _ := expr.Evaluate(nil)
	assert.Equal(t, float64(3), result)
}