package myMath

import (
	"errors"
	"math"
	"reflect"
)

type (
	anyFloat interface {
		float32 | float64
	}
	anyInt interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
	}
	anyUInt interface {
		uint | uint8 | uint16 | uint32 | uint64
	}
	anySInt interface {
		int | int8 | int16 | int32 | int64
	}
)

const (
	EnthalpyOfVaporization    = 2300             // kJ/kg
	NormalPressure            = 101325           // Pa
	AirMolarMass              = 28.98            // g/mol
	CelsiusToKelvinDifference = 273.15           // K or ºC
	MolarGasConstant          = 8.31446261815324 // J/K mol
	WaterDensity              = 997              // kg/m3
	WaterMolarMass            = 18.01528         // mol
	DryAirHeatCapacity        = 1.007            // kJ/kg K
	WaterSteamHeatCapacity    = 2.0784           // kJ/kg K
	WaterHeatCapacity         = 4.1806           // kJ/kg K
	SecondsInHour             = 3600             // s
)

type (
	noise        [8]float64
	airCondition struct {
		Temperature   float64 // in ºC
		HumidityRatio float64 // in %
	}
)

func (s airCondition) vapourPressure() float64 {
	return VaporPressure(s.Temperature)
}
func (s airCondition) vapourDensity() float64 {
	return VaporDensity(s.Temperature)
}
func (s airCondition) partialPressure() float64 {
	return PartialPressure(s.Temperature, s.HumidityRatio)
}
func (s airCondition) absHumidity() float64 {
	return VaporDensity(s.Temperature) * s.HumidityRatio * 0.01
}
func (s noise) TotalNoise() float64 {
	return LogNoiseCounter(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7])
}

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

// First is kg/h, second l/h
func CoolingCondensationFlowrate(InputAirCondition airCondition, OutgoingTemperature float64, VolumetricFlowrate float64) (float64, float64) {
	water_mass_flowrate := (InputAirCondition.absHumidity() - .95*VaporDensity(OutgoingTemperature)) * VolumetricFlowrate
	if water_mass_flowrate < 0 {
		water_mass_flowrate = 0
	}
	water_volumetric_flowrate := 1000 * water_mass_flowrate / (WaterDensity)
	return water_mass_flowrate, water_volumetric_flowrate
}

// Returns the power of evaporation/condensation process
func SteamToWaterPhaseTransitionPower[T anyFloat](MassFlowRate /* kg/h */ T) T {
	transition_energy_per_hour := MassFlowRate * EnthalpyOfVaporization
	transition_power := transition_energy_per_hour / (60 * 60)
	return transition_power
}

func HalfLengthValueSearcher[T anyFloat](Function func(T) T, Xmin T, Xmax T, TargetY T, Accuracy T) (T, error) {

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

// Returns the density (kg/m3) of the air at the selected temperature and normal pressure
func AirDensity[T anyFloat](Temperature T) T {
	return (NormalPressure * AirMolarMass) / ((Temperature + CelsiusToKelvinDifference) * MolarGasConstant * 1000)
}

// Returns the density of saturated steam at the selected temperature and normal pressure
func VaporDensity[T anyFloat](Temperature /* ºC */ T) T /* kg/m3 */ {
	return (VaporPressure(Temperature) * WaterMolarMass) / ((Temperature + CelsiusToKelvinDifference) * MolarGasConstant)
}

// Returns the pressure of saturated steam at the selected temperature
func VaporPressure[T anyFloat](Temperature /* ºC */ T) T /* kPa */ {
	var p float64
	t := float64(Temperature)
	if t > 0 {
		p = .61121 * math.Exp((18.678-t/234.5)*(t/(257.14+t)))
	} else {
		p = .61115 * math.Exp((23.036-t/333.7)*(t/(279.82+t)))
	}
	return T(p)
}

// Returns the partial pressure of the water vapour at the selected temperature and humidity ratio
func PartialPressure[T anyFloat](Temperature /* ºC */ T, HumidityRatio /* % */ T) T /* kPa */ {
	return HumidityRatio * .01 * VaporPressure(Temperature)
}

func Round[T anyFloat](DecimalPlaces int, x ...*T) {
	k := T(math.Pow10(DecimalPlaces))
	for i := 0; i < len(x); i++ {
		*x[i] = T(math.Round(float64(*x[i]*k))) / k
	}
}

// Returns the power required to heat/cool selected flowrate (Q) from the initial temperature (T_in) to the target temperature (T_out)
// T_in - inlet air temperature
// H_in - Humidity ratio
// T_out - outgoing air temperature
// Q - Air volumetric flowrate
func AirHeatPower[T anyFloat](InputAirCondition airCondition, T_out /* ºC */ float64, Q /* m3/h */ float64) float64 /* kW */ {
	return float64(math.Abs(float64(T_out-InputAirCondition.Temperature))) * Q * (AirDensity(InputAirCondition.Temperature)*DryAirHeatCapacity + InputAirCondition.HumidityRatio*0.01*VaporDensity(InputAirCondition.Temperature)*WaterSteamHeatCapacity) / SecondsInHour
}

// Returns the temperature difference of the air before and after heating/cooling
// Q - air volumetric flowrate
// P - heating power
// T_in - inlet air temperature
// H_in - Humidity ratio
func AirHeatTemperature(P /* kW */ float64, Q /* m3/h */ float64, InputAirCondition airCondition) float64 /* ºC */ {
	return P / (Q * (AirDensity(InputAirCondition.Temperature)*DryAirHeatCapacity + InputAirCondition.HumidityRatio*0.01*VaporDensity(InputAirCondition.Temperature)*WaterSteamHeatCapacity) / SecondsInHour)
}

// Returns the power that the water flow (Q) gives off when cooled from T_in to T_out
// Q - water volumetric flowrate
// T_in - inlet water temperature
// T_out - outgoing water temperature
func WaterHeatPower(T_in /* ºC */ float64, T_out /* ºC */ float64, Q /* dm3/h */ float64) float64 /* kW */ {
	return float64(math.Abs(float64(T_out-T_in))) * Q * WaterDensity * WaterHeatCapacity / (SecondsInHour * 1000)
}

// Returns the temperature difference of the in- and outgoing water flow after it has heated/cooled the air
// P - air heating power
// Q - water volumetric flowrate
// T_in - inlet water temperature
func WaterHeatTemperature(P /* kW */ float64, Q /* dm3/h */ float64, T_in /* ºC */ float64) float64 /* ºC */ {
	return P / (Q * WaterDensity * WaterHeatCapacity / (SecondsInHour * 1000))
}

// Returns the volumetric flowrate required to produce selected power
// P - required power
// T_in - inlet water temperature
// T_out - outgoing water temperature
func WaterHeatVolumetricFlowRate(P /* kW */ float64, T_in /* ºC */ float64, T_out /* ºC */ float64) float64 /* dm3/h */ {
	return P / (float64(math.Abs(float64(T_out-T_in))) * WaterDensity * WaterHeatCapacity / (SecondsInHour * 1000))
}

// Returns the air humidity ratio after heating
// H_in - initial air humidity
// T_in - initial air temperature
// T_out - temperature of air after heating
func AirHeatingHumidityRatio(H_in /* % */ float64, T_in /* ºC */ float64, T_out /* ºC */ float64) float64 /* % */ {
	return H_in * VaporPressure(T_in) / VaporPressure(T_out)
}

// Logarithmic addition of noise
func LogNoiseCounter(value /* dB(A) */ ...float64) float64 /* dB(A) */ {
	var temp float64
	for i := 0; i < len(value); i++ {
		temp += math.Pow(10, 0.1*float64(value[i]))
	}
	return float64(10 * math.Log10(temp))
}

// Returns result of linear interpolation
// f( x ) = a * x + b, f( x1 ) = val1, f( x2 ) = val2, f( target_x ) = result
func LinearInterpolation(target_x float64, x1 float64, x2 float64, val1 float64, val2 float64) float64 {
	if x1 == x2 {
		return val1
	} else {
		return -((x2*val1 - x1*val2) / (x1 - x2)) - ((-val1+val2)*target_x)/(x1-x2)
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

// This function search for the nearest higher and lower values in low to high sorted array & return it's indexes
func searchNearestId(val float64, arr *[]float64) (int, int, error) {
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
func ZeroParabolicApproximator(x1 float64, x2 float64, val1 float64, val2 float64, target_x float64) (float64, error) {
	if x1 == 0 || x2 == 0 || x1 == x2 {
		return 0, errors.New("divide dy zero")
	}

	a := -((-x2*val1 + x1*val2) / (x1 * (x1 - x2) * x2))
	b := -((x2*x2*val1 - x1*x1*val2) / (x1 * (x1 - x2) * x2))
	return a*target_x*target_x + b*target_x, nil
}
