package HVAC

import "errors"

type airCondition struct {
	Temperature float64 // in ºC
	Humidity    float64 // in g/kg
}

func (s airCondition) VapourPressure() float64 {
	return VaporPressure(s.Temperature)
}
func (s airCondition) VapourDensity() float64 {
	return VaporDensity(s.Temperature)
}
func (s airCondition) AirDensity() float64 {
	return AirDensity(s.Temperature)
}
func (s *airCondition) SetHumidityRatio(HumidityRatio float64) error {
	if HumidityRatio < 0 {
		return errors.New("invalid Humidity Ratio")
	}
	s.Humidity = HumidityRatio * s.VapourDensity() / s.AirDensity()
	return nil
}
func (s airCondition) HumidityRatio() float64 {
	return s.Humidity * s.AirDensity() / s.VapourDensity()
}

func (s airCondition) PartialPressure() float64 {
	return PartialPressure(s.Temperature, s.HumidityRatio())
}
