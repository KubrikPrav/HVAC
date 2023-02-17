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
	RequestType1 struct {
		Types map[string]HeaterTask2
		Extra Extra
	}
	ResponseType1 map[string]RespT1struct

	RequestType2  map[string]Flowrate2
	ResponseType2 map[string][]FilterDescription2

	RequestType3  map[string]BlowerReq
	ResponseType3 map[string]BlowerResp

	RequestType4  map[string]SoundModeratorTask2
	ResponseType4 map[string]SoundModeratorArray

	HeaterTask2 struct {
		Summer HeaterTask
		Winter HeaterTask
	}
	RespT1struct struct {
		Summer HeaterResult
		Winter HeaterResult
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
			InletNoise     Noise
			OutgoingNoise  Noise
			Efficiency     float64
			ConsumingPower float64
			Speed          float64
		}
		Winter struct {
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
		Outgoing                Air
		PressureDrop            float64
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
