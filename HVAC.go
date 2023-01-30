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
	anuNum interface {
		int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
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

// First is kg/h, second l/h
func CoolingCondensationFlowrate[T anyFloat](InletHumidity T, OutgoingTemperature T, VolumetricFlowrate T) (T, T) {
	water_mass_flowrate := (InletHumidity - .95*VaporDensity(OutgoingTemperature)) * VolumetricFlowrate
	if water_mass_flowrate < 0 {
		water_mass_flowrate = 0
	}
	water_volumetric_flowrate := 1000 * water_mass_flowrate / (WaterDensity)
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
	return HeatingPower / (VolumetricFlowrate * (AirDensity(InletTemperature) * (DryAirHeatCapacity + Humidity*WaterSteamHeatCapacity)) / SecondsInHour)
}

// Returns the power that the water flow (Q) gives off when cooled from InletTemperature to OutgoingTemperature
// Q - water volumetric flowrate
// InletTemperature - inlet water temperature
// OutgoingTemperature - outgoing water temperature
func WaterHeatPower(InletTemperature /* ºC */ float64, OutgoingTemperature /* ºC */ float64, VolumetricFlowrate /* dm3/h */ float64) float64 /* kW */ {
	return float64(math.Abs(float64(OutgoingTemperature-InletTemperature))) * VolumetricFlowrate * WaterDensity * WaterHeatCapacity / (SecondsInHour * 1000)
}

// Returns the temperature difference of the in- and outgoing water flow after it has heated/cooled the air
// P - air heating power
// Q - water volumetric flowrate
// InletTemperature - inlet water temperature
func WaterHeatTemperature(HeatingPower /* kW */ float64, VolumetricFlowrate /* dm3/h */ float64, InletTemperature /* ºC */ float64) float64 /* ºC */ {
	return HeatingPower / (VolumetricFlowrate * WaterDensity * WaterHeatCapacity / (SecondsInHour * 1000))
}

// Returns the volumetric flowrate required to produce selected power
// P - required power
// InletTemperature - inlet water temperature
// OutgoingTemperature - outgoing water temperature
func WaterHeatVolumetricFlowRate(HeatingPower /* kW */ float64, InletTemperature /* ºC */ float64, OutgoingTemperature /* ºC */ float64) float64 /* dm3/h */ {
	return HeatingPower / (float64(math.Abs(float64(OutgoingTemperature-InletTemperature))) * WaterDensity * WaterHeatCapacity / (SecondsInHour * 1000))
}
