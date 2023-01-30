package HVAC

import (
	"errors"
	"math"
)

func HalfLengthValueSearcher[T anyNum](Function func(T) T, Xmin T, Xmax T, TargetY T, Accuracy T) (T, error) {

	x_lower := Xmin
	x_higher := Xmax
	y_lower := Function(x_lower)
	y_higher := Function(x_higher)
	if (TargetY > y_higher && TargetY > y_lower) || (TargetY < y_higher && TargetY < y_lower) {
		return 0, errors.New("out of range")
	}
	for math.Abs(float64(y_higher-y_lower)) > float64(Accuracy) {
		x_mid := (x_higher + x_lower) / 2
		y_mid := Function(x_mid)
		if y_mid > TargetY {
			x_higher = x_mid
			y_higher = y_mid
		} else {
			x_lower = x_mid
			y_lower = y_mid
		}
	}
	return (x_higher + x_lower) / 2, nil
}

func Round[T anyFloat](DecimalPlaces int, x ...*T) {
	k := T(math.Pow10(DecimalPlaces))
	for i := 0; i < len(x); i++ {
		*x[i] = T(math.Round(float64(*x[i]*k))) / k
	}
}

// Returns result of linear interpolation
// f( x ) = a * x + b, f( x1 ) = val1, f( x2 ) = val2, f( target_x ) = result
func LinearInterpolation[T anyNum](target_x T, x1 T, x2 T, val1 T, val2 T) T {
	if x1 == x2 {
		return val1
	} else {
		return -((x2*val1 - x1*val2) / (x1 - x2)) - ((-val1+val2)*target_x)/(x1-x2)
	}
}

func BiLinearInterpolation[T anyNum](targetX T, x1 T, x2 T, targetY T, y1 T, y2 T, valX1Y1 T, valX1Y2 T, valX2Y1 T, valX2Y2 T) T {
	if x1 == x2 {
		if y1 == y2 {
			return valX1Y1
		} else {
			return LinearInterpolation(targetY, y1, y2, valX1Y1, valX1Y2)
		}
	} else if y1 == y2 {
		return LinearInterpolation(targetX, x1, x2, valX1Y1, valX2Y1)
	} else {
		return LinearInterpolation(targetX, x1, x2, LinearInterpolation(targetY, y1, y2, valX1Y1, valX1Y2), LinearInterpolation(targetY, y1, y2, valX2Y1, valX2Y2))
	}
}

func FindBLI[T anyNum](targetX T, targetY T, sortedArray []T, ValueArray any, GetVAlue func(ValueArray any, X T, Y T) (Res T, err error)) (Result T, err error) {
	var (
		higherXid int
		lowerXid  int
		higherYid int
		lowerYid  int
	)
	higherXid, lowerXid, err = SearchNearestId(targetX, &sortedArray)
	if err != nil {
		return 0, err
	}
	higherYid, lowerYid, err = SearchNearestId(targetY, &sortedArray)
	if err != nil {
		return 0, err
	}
	higherX := sortedArray[higherXid]
	lowerX := sortedArray[lowerXid]
	higherY := sortedArray[higherYid]
	lowerY := sortedArray[lowerYid]
	ValXlYl, err := GetVAlue(ValueArray, lowerX, lowerY)
	ValXlYh, err := GetVAlue(ValueArray, lowerX, higherY)
	ValXhYl, err := GetVAlue(ValueArray, higherX, lowerY)
	ValXhYh, err := GetVAlue(ValueArray, higherX, higherY)

	Result = BiLinearInterpolation(targetX, lowerX, higherX, targetY, lowerY, higherY, ValXlYl, ValXlYh, ValXhYl, ValXhYh)
	return
}

// This function search for the nearest higher and lower values in low to high sorted array & return it's indexes
func SearchNearestId[T anyNum](val T, arr *[]T) (int, int, error) {
	var (
		lower_id  int
		higher_id int
		err       error
	)
	err = errors.New("out of range")
	for i := 0; i < len(*arr)-1; i++ {
		if (*arr)[i] < val && (*arr)[i+1] > val {
			lower_id = i
			higher_id = i + 1
			err = nil
			break
		} else if (*arr)[i] == val {
			lower_id = i
			higher_id = i
			err = nil
			break
		}
	}
	if (*arr)[len(*arr)-1] == val {
		lower_id = len(*arr) - 1
		higher_id = len(*arr) - 1
		err = nil
	}
	return lower_id, higher_id, err
}

// Approximation by a 2nd order polynomial passing through zero
func ZeroParabolicApproximator[T anyNum](x1 T, x2 T, val1 T, val2 T, target_x T) (T, error) {
	if x1 == 0 || x2 == 0 || x1 == x2 {
		return 0, errors.New("divide dy zero")
	}

	a := -((-x2*val1 + x1*val2) / (x1 * (x1 - x2) * x2))
	b := -((x2*x2*val1 - x1*x1*val2) / (x1 * (x1 - x2) * x2))
	return a*target_x*target_x + b*target_x, nil
}
