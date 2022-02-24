package square

import (
	"math"
	"testing"
)

func TestCalcSquare(t *testing.T) {
	const expectedCircle = 50.26548245743669
	if res := CalcSquare(4, SidesCircle); res != expectedCircle {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expectedCircle, res)
	}

	const expectedTriangle = 27.02432272509341
	if res := CalcSquare(7.9, SidesTriangle); res != expectedTriangle {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expectedTriangle, res)
	}

	const expectedSquare = 0.25
	if res := CalcSquare(0.5, SidesSquare); res != expectedSquare {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expectedSquare, res)
	}

	const expectedUnknownShape = 0
	if res := CalcSquare(0.5, 2); res != expectedUnknownShape {
		t.Errorf("Unexpected result:\n\tExpected: %v\n\tGot: %v", expectedUnknownShape, res)
	}
}

func almostEqual(x float64, y float64) bool {
	return math.Abs(x - y) <= math.Pow(10, -6)
}
