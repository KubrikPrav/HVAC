package HVAC

type (
	Extra struct {
		HeatedWaterInletTemperature     float64
		HeatedWaterOutgoingTemperature  float64
		ChilledWaterInletTemperature    float64
		ChilledWaterOutgoingTemperature float64
		HeaterAntifreezeId              uint64
		HeaterAntifreezeQty             uint64
		CoolerAntifreezeId              uint64
		CoolerAntifreezeQty             uint64
	}
	SeasonInitData struct {
		Outdoor                   Air
		Indoor                    Air
		SupplyTarget              Air
		SupplyVolumetricFlowrate  uint64
		SupplyPressure            uint64
		ExhaustVolumetricFlowrate uint64
		ExhaustPressure           uint64
	}
	RequiredComponents struct {
		SupplyLine          bool
		ExhaustLine         bool
		ElectricHeater      bool
		HeatedWater         bool
		DirectExpansion     bool
		ChilledWater        bool
		SteamHumidifier     bool
		MediaHumidifier     bool
		Dryer               bool
		NoHeater            bool
		SupplyFilterClasses []uint8
		ExhaustFilterClass  []uint8
	}
	UnitTask struct {
		Summer                      SeasonInitData
		Winter                      SeasonInitData
		SupplyBlowerEfficiencyClass string
		MaxInsideNoise              uint64
		MaxOutsideNoise             uint64
		SectionList                 RequiredComponents
		Extra                       struct {
			HeatedWaterInletTemperature     float64
			HeatedWaterOutgoingTemperature  float64
			ChilledWaterInletTemperature    float64
			ChilledWaterOutgoingTemperature float64
			HeaterAntifreezeId              uint64
			HeaterAntifreezeQty             uint64
			CoolerAntifreezeId              uint64
			CoolerAntifreezeQty             uint64
		}
	}
	UnitDescription struct {
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
		HeatRecovery              HeatRecoveryResult2
		PreHeater                 HeaterResult2
		Heater                    HeaterResult2
		Cooler                    HeaterResult2
		Humidifier                HeaterResult2
		SoundModerator            SoundModeratorArray
		SupplyFilter              []FilterDescription2
		SupplyBlower              BlowerResp
		ExhaustFilter             []FilterDescription2
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
	}
	RequestType1 struct {
		Types map[string]HeaterTask2
		Extra Extra
	}
	ResponseType1 map[string]HeaterResult2

	RequestType2  map[string]Flowrate2
	ResponseType2 map[string][]FilterDescription2

	RequestType3  map[string]BlowerReq
	ResponseType3 map[string]BlowerResp

	RequestType4  map[string]SoundModeratorTask2
	ResponseType4 map[string]SoundModeratorArray

	RequestType5 struct {
		Task  HeatRecoveryTask2
		Names []string
	}
	ResponseType5 map[string]HeatRecoveryResult2

	HeatRecoveryTask struct {
		Inside                    Air
		Outside                   Air
		SupplyVolumetricFlowrate  uint64
		ExhaustVolumetricFlowrate uint64
	}

	HeatRecoveryResult struct {
		Supply                Air
		Exhaust               Air
		Inside                Air
		Outside               Air
		PreHeaterPwr          float64
		SupplyPressureDrop    float64
		ExhaustPressureDrop   float64
		TemperatureEfficiency float64
		HumidityEfficiency    float64
		Freeze                bool
	}
	HeatRecoveryTask2 struct {
		Summer HeatRecoveryTask
		Winter HeatRecoveryTask
	}
	HeatRecoveryResult2 struct {
		LongName  string
		ShortName string
		Summer    HeatRecoveryResult
		Winter    HeatRecoveryResult
	}

	HeaterTask2 struct {
		Name   string
		Summer HeaterTask
		Winter HeaterTask
	}
	HeaterResult2 struct {
		LongName  string
		ShortName string
		Summer    HeaterResult
		Winter    HeaterResult
	}
	Flowrate2 struct {
		Classes                  []uint8
		SummerVolumetricFlowrate float64
		WinterVolumetricFlowrate float64
	}
	FilterDescription2 struct {
		Summer FilterDescription
		Winter FilterDescription
	}
	FilterDescription struct {
		Flowrate     float64
		PressureDrop float64
		Class        string
		SizeAndQty   []struct {
			Size string
			Qty  uint64
		}
	}
	BlowerReq struct {
		Summer           BlowerOperatingPoint
		Winter           BlowerOperatingPoint
		EfficiencyClass  string
		MaxInletNoise    uint64
		MaxOutgoingNoise uint64
	}
	BlowerOperatingPoint struct {
		Flowrate uint64
		Pressure uint64
	}
	BlowerResp struct {
		Summer struct {
			OperatingPoint BlowerOperatingPoint
			InletNoise     Noise
			OutgoingNoise  Noise
			Efficiency     float64
			ConsumingPower float64
			Speed          float64
		}
		Winter struct {
			OperatingPoint BlowerOperatingPoint
			InletNoise     Noise
			OutgoingNoise  Noise
			Efficiency     float64
			ConsumingPower float64
			Speed          float64
		}
		Plot struct {
			Flowrate []float64
			Pressure []float64
		}
		LongName        string
		ShortName       string
		Voltage         string
		EfficiencyClass string
		MaxCurrent      float64
		WheelSize       uint64
		MotorPower      float64
		Length          int64
		TooLoud         bool
		Twin            bool
	}
	HeaterTask struct {
		Inlet              Air
		Target             float64
		VolumetricFlowrate uint64
	}
	HeaterResult struct {
		Inlet                   Air
		Outgoing                Air
		PressureDrop            float64
		VolumetricFlowrate      float64
		Capacity                float64 // max available power
		Power                   float64
		WaterVolumetricFlowrate float64
		WaterPressureDrop       float64
		LowCapacity             bool
	}

	SoundModeratorTask2 struct {
		InsideUpper  SoundModeratorTask
		InsideLower  SoundModeratorTask
		OutsideUpper SoundModeratorTask
		OutsideLower SoundModeratorTask
	}
	SoundModeratorArray struct {
		InsideUpper  SoundModeratorDescription
		InsideLower  SoundModeratorDescription
		OutsideUpper SoundModeratorDescription
		OutsideLower SoundModeratorDescription
	}
	SoundDutyPoint struct {
		InletNoise    Noise
		OutgoingNoise Noise
		Flowrate      float64
		PressureDrop  float64
	}
	SoundModeratorDescription struct {
		NotRequired bool
		LongName    string
		ShortName   string
		Summer      SoundDutyPoint
		Winter      SoundDutyPoint
	}
	SoundModeratorTask struct {
		SummerNoise    Noise
		WinterNoise    Noise
		SummerFlowrate float64
		WinterFlowrate float64
		Target         float64
	}
)
