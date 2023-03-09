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
		SupplyFilterClasses []string
		ExhaustFilterClass  []string
	}
	UnitTask struct {
		Summer                      SeasonInitData
		Winter                      SeasonInitData
		SupplyBlowerEfficiencyClass string
		MaxInsideNoise              uint64
		MaxOutsideNoise             uint64
		SectionList                 RequiredComponents
		ServiceSide                 string
		IsAutomatics                bool
		IsOutside                   bool
		HousingType                 struct {
			Vertical   bool
			Flat       bool
			Horizontal bool
		}
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
		RightServiceSide          bool
		IsAutomatics              bool
		IsOutside                 bool
		Name                      string
		Plot                      string
		Drawing                   string
		Price                     float64
		Task                      UnitTask
		Result                    UnitResult2
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
		UnitSpec                  UnitSpec
		TotalNoise                struct {
			Inside       Noise
			Outside      Noise
			Body         Noise
			InsideTotal  float64
			OutsideTotal float64
			BodyTotal    float64
		}
		Dimensions struct {
			Height uint64
			Width  uint64
			Length uint64
		}
	}
	UnitSpec struct {
		Internals PartList
		Housing   PartList
	}
	PartList []struct {
		LongName  string
		ShortName string
		Qty       uint64
	}
	UnitResult2 struct {
		Summer UnitResult
		Winter UnitResult
	}
	UnitResult struct {
		Outdoor              Air
		Indoor               Air
		Supply               Air
		Exhaust              Air
		SupplyFlowrate       uint64
		ExhaustFlowrate      uint64
		SupplyPressure       uint64
		ExhaustPressure      uint64
		SupplyTotalPressure  uint64
		ExhaustTotalPressure uint64
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
	RequestType5  struct {
		Task  HeatRecoveryTask2
		Names []string
	}
	ResponseType5    map[string]HeatRecoveryResult2
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
		Length    uint64
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
		Length    uint64
		Summer    HeaterResult
		Winter    HeaterResult
	}
	Flowrate2 struct {
		Classes                  []string
		SummerVolumetricFlowrate float64
		WinterVolumetricFlowrate float64
	}
	FilterDescription2 struct {
		Length     uint64
		Class      string
		SizeAndQty []struct {
			Size string
			Qty  uint64
		}
		Summer FilterDescription
		Winter FilterDescription
	}
	FilterDescription struct {
		Flowrate     float64
		PressureDrop float64
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
		Length          uint64
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
		Length      uint64
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
	DrawingTask struct {
		RtoL  bool
		High  []string // List of high sections. 0 is left, len-1 is right
		Upper []string // List of upper sections. 0 is left, len-1 is right
		Lower []string // List of lower sections. 0 is left, len-1 is right
	}
	DrawingResult struct {
		Picture     string // Here should be a picture in svg format
		TotalLength string // <Sum of all High sections length> + Max(<Sum of all Lower sections length>, <Sum of all Upper sections length>)
		TotalHeight string
		TotalWidth  string
	}
	DrawingRequest  map[string]DrawingTask
	DrawingResponse map[string]DrawingResult
)

func (s *PartList) Add(LongName string, ShortName string, qty uint64) {
	for i := 0; i < len(*s); i++ {
		if (*s)[i].LongName == LongName {
			val := (*s)[i]
			val.Qty += qty
			(*s)[i] = val
			return
		}
	}
	val := struct {
		LongName  string
		ShortName string
		Qty       uint64
	}{
		LongName:  LongName,
		ShortName: ShortName,
		Qty:       qty,
	}
	*s = append(*s, val)
}

func (s UnitSpec) Names(LongNames []string) {
	LongNames = make([]string, len(s.Internals)+len(s.Housing))
	for i := 0; i < len(s.Internals)+len(s.Housing); i++ {
		if i < len(s.Housing) {
			LongNames[i] = s.Housing[i].LongName
		} else {
			LongNames[i] = s.Internals[i-len(s.Housing)].LongName
		}
	}
}
