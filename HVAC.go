package HVAC

import (
	"math"
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
	anyNum interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
	}
)

const (
	EnthalpyOfVaporization    = 2300             // kJ/kg
	NormalPressure            = 101325           // Pa
	AirMolarMass              = 28.98            // g/mol
	CelsiusToKelvinDifference = 273.15           // K or ºC
	MolarGasConstant          = 8.31446261815324 // J/K mol
	WaterMolarMass            = 18.01528         // mol
	DryAirHeatCapacity        = 1.007            // kJ/kg K
	WaterSteamHeatCapacity    = 2.0784           // kJ/kg K
	SecondsInHour             = 3600             // s
)

// First is kg/h, second l/h
func CoolingCondensationFlowrate[T anyFloat](InletHumidity T, OutgoingTemperature T, VolumetricFlowrate T) (T, T) {
	water_mass_flowrate := (InletHumidity - .95*VaporDensity(OutgoingTemperature)) * VolumetricFlowrate
	if water_mass_flowrate < 0 {
		water_mass_flowrate = 0
	}
	water_volumetric_flowrate := 1000 * water_mass_flowrate / (WaterDensity(OutgoingTemperature))
	return water_mass_flowrate, water_volumetric_flowrate
}

// Returns the power of evaporation/condensation process
func SteamToWaterPhaseTransitionPower[T anyFloat](MassFlowRate /* kg/h */ T) (TransitionPower T) {
	TransitionPower = MassFlowRate * EnthalpyOfVaporization / SecondsInHour
	return
}

// Returns the density (kg/m3) of the air at the selected temperature and normal pressure
func AirDensity[T anyFloat](Temperature T) (Density T) {
	return (NormalPressure * AirMolarMass) / ((Temperature + CelsiusToKelvinDifference) * MolarGasConstant * 1000)
}

// Returns the density of saturated steam at the selected temperature and normal pressure
func VaporDensity[T anyFloat](Temperature /* ºC */ T) (Density T) /* kg/m3 */ {
	return (VaporPressure(Temperature) * WaterMolarMass) / ((Temperature + CelsiusToKelvinDifference) * MolarGasConstant)
}

// Returns the pressure of saturated steam at the selected temperature
func VaporPressure[T anyFloat](Temperature /* ºC */ T) (Pressure T /* kPa */) {
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

// Returns the power required to heat/cool selected flowrate (Q) from the initial temperature (InletTemperature) to the target temperature (OutgoingTemperature)
func AirHeatPower[T anyFloat](InletTemperature /* ºC */ T, Humidity /* g/kg */ T, OutgoingTemperature /* ºC */ T, VolumetricFlowrate /* m3/h */ T) T /* kW */ {
	return T(math.Abs(float64(OutgoingTemperature-InletTemperature))) * VolumetricFlowrate * (AirDensity(InletTemperature) * (DryAirHeatCapacity + Humidity*WaterSteamHeatCapacity)) / SecondsInHour
}

// Returns the temperature difference of the air before and after heating/cooling
// HeatingPower - kW
// VolumetricFlowrate - m3/h
// InletTemperature - ºC
// Humidity - g/kg
func AirHeatOutgoingTemperature[T anyFloat](HeatingPower /* kW */ T, VolumetricFlowrate /* m3/h */ T, InletTemperature /* ºC */ T, Humidity /* g/kg */ T) (OutgoingTemperature /* ºC */ T) {
	return HeatingPower/(VolumetricFlowrate*(AirDensity(InletTemperature)*(DryAirHeatCapacity+Humidity*WaterSteamHeatCapacity))/SecondsInHour) + InletTemperature
}

// Returns the power that the water flow (Q) gives off when cooled from InletTemperature to OutgoingTemperature
// Q - water volumetric flowrate
// InletTemperature - inlet water temperature
// OutgoingTemperature - outgoing water temperature
func WaterHeatPower(InletTemperature /* ºC */ float64, OutgoingTemperature /* ºC */ float64, VolumetricFlowrate /* dm3/h */ float64) float64 /* kW */ {
	return float64(math.Abs(float64(OutgoingTemperature-InletTemperature))) * VolumetricFlowrate * WaterDensity(InletTemperature) * WaterHeatCapacity(InletTemperature) / (SecondsInHour * 1000)
}

// Returns the temperature difference of the in- and outgoing water flow after it has heated/cooled the air
// P - air heating power
// Q - water volumetric flowrate
// InletTemperature - inlet water temperature
func WaterHeatTemperature(HeatingPower /* kW */ float64, VolumetricFlowrate /* dm3/h */ float64, InletTemperature /* ºC */ float64) float64 /* ºC */ {
	return HeatingPower / (VolumetricFlowrate * WaterDensity(InletTemperature) * WaterHeatCapacity(InletTemperature) / (SecondsInHour * 1000))
}

// Returns the volumetric flowrate required to produce selected power
// P - required power
// InletTemperature - inlet water temperature
// OutgoingTemperature - outgoing water temperature
func WaterHeatVolumetricFlowRate(HeatingPower /* kW */ float64, InletTemperature /* ºC */ float64, OutgoingTemperature /* ºC */ float64) float64 /* dm3/h */ {
	return HeatingPower / (float64(math.Abs(float64(OutgoingTemperature-InletTemperature))) * WaterDensity(InletTemperature) * WaterHeatCapacity(InletTemperature) / (SecondsInHour * 1000))
}

func WaterDensity[T anyFloat](t T) T {
	x := float64(t)
	return T(1000 * (0.999922 + 0.0000475377*x - 7.34753*math.Pow10(-6)*math.Pow(x, 2) + 3.92894*math.Pow10(-8)*math.Pow(x, 3) - 1.2144*math.Pow10(-10)*math.Pow(x, 4)))
}

func WaterHeatCapacity[T anyFloat](t T) T {
	x := float64(t)
	y := 4.21701568841111 - 0.007849993116176771*x + 0.0010067730629546086*math.Pow(x, 2) - 0.00006957457567073951*math.Pow(x, 3) + 2.605607116793847*math.Pow10(-6)*math.Pow(x, 4) - 5.609144300760924*math.Pow10(-8)*math.Pow(x, 5) + 7.141795677520566*math.Pow10(-10)*math.Pow(x, 6) - 5.286838922592877*math.Pow10(-12)*math.Pow(x, 7) + 2.0881591874928128*math.Pow10(-14)*math.Pow(x, 8) - 3.366302230773775*math.Pow10(-17)*math.Pow(x, 9)
	return T(y)
}
