package HVAC

import (
	"math"
	"reflect"
)

func AssignByName(input any, output any) {
	input_value := reflect.Indirect(reflect.ValueOf(input))
	output_value := reflect.Indirect(reflect.ValueOf(output))
	len_i := input_value.NumField()
	len_o := output_value.NumField()
	for i := 0; i < len_i; i++ {
		for o := 0; o < len_o; o++ {
			if input_value.Type().Field(i).Name == output_value.Type().Field(o).Name && input_value.Type().Field(i).Type == output_value.Type().Field(o).Type {
				if output_value.Field(o).CanSet() {
					output_value.Field(o).Set(input_value.Field(i))
				}
				break
			}
		}
	}
}

// standard logistic sigmoid function for array of args
func LogisticSigmoid(x ...float64) []float64 {
	result := make([]float64, len(x))
	for index, val := range x {
		result[index] = float64(1 / (1 + math.Exp(-float64(val))))
	}
	return result
}
