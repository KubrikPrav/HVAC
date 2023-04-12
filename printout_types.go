package HVAC

import (
	"math"
)

type (
	UnitPrintout struct {
		NOk                       bool
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
		IsAutomatics              bool
		IsOutside                 bool
		ServiceSide               string
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
		TotalNoise                NoiseResponse1
		Description               []struct {
			Name  string
			Value string
		}
	}

	TaskPrint2 struct {
		Summer TaskPrint
		Winter TaskPrint
		TSet   bool
		HSet   bool
	}
	TaskPrint struct {
		Outdoor         PrintAir
		Indoor          PrintAir
		SupplyTarget    PrintAir
		SupplyFlowrate  uint64
		ExhaustFlowrate uint64
		SupplyPressure  uint64
		ExhaustPressure uint64
	}

	ResultPrint2 struct {
		Summer ResultPrint
		Winter ResultPrint
	}
	ResultPrint struct {
		Outdoor              PrintAir
		Indoor               PrintAir
		Supply               PrintAir
		Exhaust              PrintAir
		SupplyFlowrate       uint64
		ExhaustFlowrate      uint64
		SupplyPressure       uint64
		ExhaustPressure      uint64
		SupplyTotalPressure  uint64
		ExhaustTotalPressure uint64
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
			HeatRecovery          float64
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
			HeatRecovery          float64
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
			HeatRecovery          float64
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
			HeatRecovery:          round(s.Summer.HeatRecovery, digits),
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
			HeatRecovery          float64
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
			HeatRecovery:          round(s.Summer.HeatRecovery, digits),
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
		Class:      s.Class,
		SizeAndQty: s.SizeAndQty,
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
func (s UnitResult) Print(d int) ResultPrint {
	return ResultPrint{
		Outdoor:              s.Outdoor.PrintAir(d),
		Indoor:               s.Indoor.PrintAir(d),
		Supply:               s.Supply.PrintAir(d),
		Exhaust:              s.Exhaust.PrintAir(d),
		SupplyFlowrate:       s.SupplyFlowrate,
		ExhaustFlowrate:      s.ExhaustFlowrate,
		SupplyPressure:       s.SupplyPressure,
		ExhaustPressure:      s.ExhaustPressure,
		SupplyTotalPressure:  s.SupplyTotalPressure,
		ExhaustTotalPressure: s.ExhaustTotalPressure,
	}
}
func (s UnitResult2) Print(d int) ResultPrint2 {
	return ResultPrint2{
		Summer: s.Summer.Print(d),
		Winter: s.Winter.Print(d),
	}
}
func (s SeasonInitData) Print(d int) TaskPrint {
	return TaskPrint{
		Outdoor:         s.Outdoor.PrintAir(d),
		Indoor:          s.Indoor.PrintAir(d),
		SupplyTarget:    s.SupplyTarget.PrintAir(d),
		SupplyFlowrate:  s.SupplyVolumetricFlowrate,
		ExhaustFlowrate: s.ExhaustVolumetricFlowrate,
		SupplyPressure:  s.SupplyPressure,
		ExhaustPressure: s.ExhaustPressure,
	}
}
func (s UnitTask) Print(d int) TaskPrint2 {
	return TaskPrint2{
		Summer: s.Summer.Print(d),
		Winter: s.Winter.Print(d),
		TSet:   s.Summer.SupplyTarget.Temperature != 0 || s.Winter.SupplyTarget.Temperature != 0,
		HSet:   s.Summer.SupplyTarget.Humidity != 0 || s.Winter.SupplyTarget.Humidity != 0,
	}
}
func serviceSidePrint(b bool) string {
	if b {
		return "Правая"
	} else {
		return "Левая"
	}
}
func (s UnitDescription) Print(digits int) UnitPrintout {
	return UnitPrintout{
		NOk:                       s.NOk,
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
		IsAutomatics:              s.IsAutomatics,
		IsOutside:                 s.IsOutside,
		ServiceSide:               serviceSidePrint(s.RightServiceSide),
		Name:                      s.Name,
		Plot:                      s.Plot,
		Drawing:                   s.Drawing,
		Price:                     s.Price,
		Task:                      s.Task.Print(digits),
		Result:                    s.Result.Print(digits),
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
		TotalNoise: NoiseResponse1{
			ODA:  s.TotalNoise.ODA.Round(0),
			SUP:  s.TotalNoise.SUP.Round(0),
			ETA:  s.TotalNoise.ETA.Round(0),
			EHA:  s.TotalNoise.EHA.Round(0),
			Body: s.TotalNoise.Body.Round(0),
			Room: s.TotalNoise.Room.Round(0),
		},
	}
}
