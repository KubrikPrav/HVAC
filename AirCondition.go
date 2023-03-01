package HVAC

import (
	"errors"
)

type Air struct {
	Temperature float64 // in ºC
	Humidity    float64 // in g/kg
}
type PrintAir struct {
	Temperature   float64
	HumidityRatio float64
}

func (s *Air) VapourPressure() float64 {
	return VaporPressure(s.Temperature)
}
func (s *Air) VapourDensity() float64 {
	return VaporDensity(s.Temperature)
}
func (s *Air) AirDensity() float64 {
	return AirDensity(s.Temperature)
}
func (s *Air) SetHumidityRatio(HumidityRatio float64) error {
	if HumidityRatio < 0 {
		return errors.New("invalid Humidity Ratio")
	}
	s.Humidity = (HumidityRatio / 100) * s.VapourDensity() / s.AirDensity()
	return nil
}
func (s *Air) HumidityRatio() float64 {
	return s.Humidity * 100 * s.AirDensity() / s.VapourDensity()
}

func (s *Air) PartialPressure() float64 {
	return PartialPressure(s.Temperature, s.HumidityRatio())
}
