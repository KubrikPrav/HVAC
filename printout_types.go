package HVAC

import (
	"math"
)

type (
    UnitPrintout struct {
        IsHeatedWaterPreHeater    bool
        IsElectricHeaterPreHeater bool
        IsHeatedWater             bool
        IsElectricHeater          bool
        IsChilledWater            bool
        IsDirectExpansion         bool
        IsMediaHumidifier         bool
        IsSteamHumidifier         bool
        IsSoundModerator          bool
        IsSupplyFilter            bool
        IsExhaustFilter           bool
        IsThermalWheel            bool
        IsPlateHeatExchanger      bool
        IsSupplyBlower            bool
        IsExhaustBlower           bool
        Name                      string
        Plot                      string
        Drawing                   string
        Price                     float64
        Task                      TaskPrint2
        Result                    ResultPrint2
        HeatRecovery              HeatRecoveryPrint
        PreHeater                 HeaterPrint
        Heater                    HeaterPrint
        Cooler                    HeaterPrint
        Humidifier                HeaterPrint
        SoundModerator            SoundModeratorArray
        SupplyFilter              []FilterPrint
        SupplyBlower              BlowerResp
        ExhaustFilter             []FilterPrint
        ExhaustBlower             BlowerResp
        Extra                     Extra
        TotalNoise                struct {
            Inside       Noise
            Outside      Noise
            Body         Noise
            InsideTotal  float64
            OutsideTotal float64
            BodyTotal    float64
        }
        Description []struct {
            Name  string
            Value string
        }
    }

    TaskPrint2 struct{
        Summer TaskPrint
        Winter TaskPrint
        TSet bool
        HSet bool
    }
    TaskPrint struct{
        Outdoor PrintAir
        Indoor PrintAir
        SupplyTarget PrintAir
        SupplyFlowrate uint64
        ExhaustFlowrate uint64
        SupplyPressure uint64
        ExhaustPressure uint64
    }

    ResultPrint2 struct{
        Summer ResultPrint
        Winter ResultPrint
    }
    ResultPrint struct{
        Outdoor PrintAir
        Indoor PrintAir
        Supply PrintAir
        Exhaust PrintAir
        Flowrate uint64
        ExhaustFlowrate uint64
        SupplyPressure uint64
        ExhaustPressure uint64
    }
	FilterPrint struct {
		Class      string
		SizeAndQty []struct {
			Size string
			Qty  uint64
		}
		Summer struct {
			Flowrate     float64
			PressureDrop float64
		}
		Winter struct {
			Flowrate     float64
			PressureDrop float64
		}
	}
	HeatRecoveryPrint struct {
		LongName  string
		ShortName string
		Used      bool
		Summer    struct {
			Used                  bool
			Supply                PrintAir
			Exhaust               PrintAir
			Inside                PrintAir
			Outside               PrintAir
			SupplyPressureDrop    float64
			ExhaustPressureDrop   float64
			TemperatureEfficiency float64
			HumidityEfficiency    float64
		}
		Winter struct {
			Used                  bool
			Supply                PrintAir
			Exhaust               PrintAir
			Inside                PrintAir
			Outside               PrintAir
			SupplyPressureDrop    float64
			ExhaustPressureDrop   float64
			TemperatureEfficiency float64
			HumidityEfficiency    float64
		}
	}
	HeaterPrint struct {
		LongName  string
		ShortName string
		Used      bool
		Summer    struct {
			Used                    bool
			Inlet                   PrintAir
			Outgoing                PrintAir
			PressureDrop            float64
			VolumetricFlowrate      float64
			Capacity                float64 // max available power
			Power                   float64
			WaterVolumetricFlowrate float64
			WaterPressureDrop       float64
		}
		Winter struct {
			Used                    bool
			Inlet                   PrintAir
			Outgoing                PrintAir
			PressureDrop            float64
			VolumetricFlowrate      float64
			Capacity                float64 // max available power
			Power                   float64
			WaterVolumetricFlowrate float64
			WaterPressureDrop       float64
		}
	}
)

// Converters

func (s *PrintAir) Air() (out Air) {
	out.Temperature = s.Temperature
	out.SetHumidityRatio(s.HumidityRatio)
	return
}

func (s *Air) PrintAir(digits int) (out PrintAir) {
	return PrintAir{math.Round(s.Temperature*math.Pow10(digits)) / math.Pow10(digits), math.Round(s.HumidityRatio()*math.Pow10(digits)) / math.Pow10(digits)}
}

func round(num float64, digits int) float64 {
	return math.Round(num*math.Pow10(digits)) / math.Pow10(digits)
}

func (s *HeaterResult2) Print(digits int) HeaterPrint {
	return HeaterPrint{
		LongName:  s.LongName,
		ShortName: s.ShortName,
		Summer: struct {
			Used                    bool
			Inlet                   PrintAir
			Outgoing                PrintAir
			PressureDrop            float64
			VolumetricFlowrate      float64
			Capacity                float64
			Power                   float64
			WaterVolumetricFlowrate float64
			WaterPressureDrop       float64
		}{
			Used:                    s.Summer.Inlet != s.Summer.Outgoing,
			Inlet:                   s.Summer.Inlet.PrintAir(digits),
			Outgoing:                s.Summer.Outgoing.PrintAir(digits),
			PressureDrop:            round(s.Summer.PressureDrop, digits),
			VolumetricFlowrate:      round(s.Summer.VolumetricFlowrate, digits),
			Capacity:                round(s.Summer.Capacity, digits),
			Power:                   round(s.Summer.Power, digits),
			WaterVolumetricFlowrate: round(s.Summer.WaterVolumetricFlowrate, digits),
			WaterPressureDrop:       round(s.Summer.WaterPressureDrop, digits),
		},
		Winter: struct {
			Used                    bool
			Inlet                   PrintAir
			Outgoing                PrintAir
			PressureDrop            float64
			VolumetricFlowrate      float64
			Capacity                float64
			Power                   float64
			WaterVolumetricFlowrate float64
			WaterPressureDrop       float64
		}{
			Used:                    s.Winter.Inlet != s.Winter.Outgoing,
			Inlet:                   s.Winter.Inlet.PrintAir(digits),
			Outgoing:                s.Winter.Outgoing.PrintAir(digits),
			PressureDrop:            round(s.Winter.PressureDrop, digits),
			VolumetricFlowrate:      round(s.Winter.VolumetricFlowrate, digits),
			Capacity:                round(s.Winter.Capacity, digits),
			Power:                   round(s.Winter.Power, digits),
			WaterVolumetricFlowrate: round(s.Winter.WaterVolumetricFlowrate, digits),
			WaterPressureDrop:       round(s.Winter.WaterPressureDrop, digits),
		},
	}
}

func (s HeatRecoveryResult2) Print(digits int) HeatRecoveryPrint {
	return HeatRecoveryPrint{
		ShortName: s.ShortName,
		LongName:  s.LongName,
		Summer: struct {
			Used                  bool
			Supply                PrintAir
			Exhaust               PrintAir
			Inside                PrintAir
			Outside               PrintAir
			SupplyPressureDrop    float64
			ExhaustPressureDrop   float64
			TemperatureEfficiency float64
			HumidityEfficiency    float64
		}{
			Used:                  s.Summer.Supply != s.Summer.Outside,
			Supply:                s.Summer.Supply.PrintAir(digits),
			Exhaust:               s.Summer.Exhaust.PrintAir(digits),
			Inside:                s.Summer.Inside.PrintAir(digits),
			Outside:               s.Summer.Outside.PrintAir(digits),
			SupplyPressureDrop:    round(s.Summer.SupplyPressureDrop, digits),
			ExhaustPressureDrop:   round(s.Summer.ExhaustPressureDrop, digits),
			TemperatureEfficiency: round(s.Summer.TemperatureEfficiency, digits),
			HumidityEfficiency:    round(s.Summer.HumidityEfficiency, digits),
		},
		Winter: struct {
			Used                  bool
			Supply                PrintAir
			Exhaust               PrintAir
			Inside                PrintAir
			Outside               PrintAir
			SupplyPressureDrop    float64
			ExhaustPressureDrop   float64
			TemperatureEfficiency float64
			HumidityEfficiency    float64
		}{
			Used:                  s.Winter.Supply != s.Winter.Outside,
			Supply:                s.Winter.Supply.PrintAir(digits),
			Exhaust:               s.Winter.Exhaust.PrintAir(digits),
			Inside:                s.Winter.Inside.PrintAir(digits),
			Outside:               s.Winter.Outside.PrintAir(digits),
			SupplyPressureDrop:    round(s.Winter.SupplyPressureDrop, digits),
			ExhaustPressureDrop:   round(s.Winter.ExhaustPressureDrop, digits),
			TemperatureEfficiency: round(s.Winter.TemperatureEfficiency, digits),
			HumidityEfficiency:    round(s.Winter.HumidityEfficiency, digits),
		},
	}
}

func (s SoundModeratorDescription) Round(digits int) SoundModeratorDescription {
	return SoundModeratorDescription{
		NotRequired: s.NotRequired,
		LongName:    s.LongName,
		ShortName:   s.ShortName,
		Summer: SoundDutyPoint{
			InletNoise:    s.Summer.InletNoise.Round(digits),
			OutgoingNoise: s.Summer.OutgoingNoise.Round(digits),
			Flowrate:      round(s.Summer.Flowrate, digits),
			PressureDrop:  round(s.Summer.PressureDrop, digits),
		},
		Winter: SoundDutyPoint{
			InletNoise:    s.Winter.InletNoise.Round(digits),
			OutgoingNoise: s.Winter.OutgoingNoise.Round(digits),
			Flowrate:      round(s.Winter.Flowrate, digits),
			PressureDrop:  round(s.Winter.PressureDrop, digits),
		},
	}
}
func (s SoundModeratorArray) Round(digits int) SoundModeratorArray {
	return SoundModeratorArray{
		InsideLower:  s.InsideLower.Round(digits),
		InsideUpper:  s.InsideUpper.Round(digits),
		OutsideLower: s.OutsideLower.Round(digits),
		OutsideUpper: s.OutsideUpper.Round(digits),
	}
}

func (s FilterDescription2) Print(digits int) FilterPrint {
	return FilterPrint{
		Class:      s.Summer.Class,
		SizeAndQty: s.Summer.SizeAndQty,
		Summer: struct {
			Flowrate     float64
			PressureDrop float64
		}{
			Flowrate:     round(s.Summer.Flowrate, digits),
			PressureDrop: round(s.Summer.PressureDrop, digits),
		},
		Winter: struct {
			Flowrate     float64
			PressureDrop float64
		}{
			Flowrate:     round(s.Winter.Flowrate, digits),
			PressureDrop: round(s.Winter.PressureDrop, digits),
		},
	}
}
func printFilters(in []FilterDescription2, digits int) (out []FilterPrint) {
	for i := 0; i < len(in); i++ {
		out = append(out, in[i].Print(digits))
	}
	return
}

func (s UnitDescription) Print(digits int) UnitPrintout {
	return UnitPrintout{
		IsHeatedWaterPreHeater:    s.IsHeatedWaterPreHeater,
		IsElectricHeaterPreHeater: s.IsElectricHeaterPreHeater,
		IsHeatedWater:             s.IsHeatedWater,
		IsElectricHeater:          s.IsElectricHeater,
		IsChilledWater:            s.IsChilledWater,
		IsDirectExpansion:         s.IsDirectExpansion,
		IsMediaHumidifier:         s.IsMediaHumidifier,
		IsSteamHumidifier:         s.IsSteamHumidifier,
		IsSoundModerator:          s.IsSoundModerator,
		IsSupplyFilter:            s.IsSupplyFilter,
		IsExhaustFilter:           s.IsExhaustFilter,
		IsThermalWheel:            s.IsThermalWheel,
		IsPlateHeatExchanger:      s.IsPlateHeatExchanger,
		IsSupplyBlower:            s.IsSupplyBlower,
		IsExhaustBlower:           s.IsExhaustBlower,
		Name:                      s.Name,
		Plot:                      s.Plot,
		Drawing:                   s.Drawing,
		Price:                     s.Price,
		HeatRecovery:              s.HeatRecovery.Print(digits),
		PreHeater:                 s.PreHeater.Print(digits),
		Heater:                    s.Heater.Print(digits),
		Cooler:                    s.Cooler.Print(digits),
		Humidifier:                s.Humidifier.Print(digits),
		SoundModerator:            s.SoundModerator.Round(digits),
		SupplyFilter:              printFilters(s.SupplyFilter, digits),
		SupplyBlower:              s.SupplyBlower,
		ExhaustFilter:             printFilters(s.ExhaustFilter, digits),
		ExhaustBlower:             s.ExhaustBlower,
		Extra:                     s.Extra,
		TotalNoise: struct {
			Inside       Noise
			Outside      Noise
			Body         Noise
			InsideTotal  float64
			OutsideTotal float64
			BodyTotal    float64
		}{
			Inside:       s.TotalNoise.Inside.Round(digits),
			Outside:      s.TotalNoise.Outside.Round(digits),
			Body:         s.TotalNoise.Body.Round(digits),
			InsideTotal:  round(s.TotalNoise.InsideTotal, digits),
			OutsideTotal: round(s.TotalNoise.OutsideTotal, digits),
			BodyTotal:    round(s.TotalNoise.BodyTotal, digits),
		},
	}
}
