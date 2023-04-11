package HVAC

import "math"

type (
	Noise [8]float32
)

func (s Noise) Round(digits int) Noise {
	return Noise{
		float32(round(float64(s[0]), digits)),
		float32(round(float64(s[1]), digits)),
		float32(round(float64(s[2]), digits)),
		float32(round(float64(s[3]), digits)),
		float32(round(float64(s[4]), digits)),
		float32(round(float64(s[5]), digits)),
		float32(round(float64(s[6]), digits)),
		float32(round(float64(s[7]), digits)),
	}
}

func (s Noise) TotalNoise() float32 {
	var isZero bool = true
	for i := 0; i < 8; i++ {
		isZero = isZero && s[i] <= .01
	}
	if isZero {
		return 0
	} else {
		return float32(math.Round(float64(logNoiseCounter(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7]))))
	}
}

func (s Noise) TotalNoiseA() float32 {
	var isZero bool = true
	tmp := s.AScale()
	for i := 0; i < 8; i++ {
		isZero = isZero && tmp[i] <= .01
	}
	if isZero {
		return 0
	} else {
		return float32(math.Round(float64(logNoiseCounter(tmp[0], tmp[1], tmp[2], tmp[3], tmp[4], tmp[5], tmp[6], tmp[7]))))
	}
}

func (s Noise) AScale() Noise {
	aScale := [8]float32{-26.2, -16.1, -8.6, -3.2, 0, 1.2, 1, -1.1}
	var res [8]float32
	for i := uint(0); i < 8; i++ {
		res[i] = s[i] + aScale[i]
	}
	return res
}

func AddNoise(in ...Noise) (out Noise) {
	var temp [8][]float32
	for i := 0; i < 8; i++ {
		temp[i] = make([]float32, len(in))
	}
	for i := 0; i < len(in); i++ {
		for j := 0; j < 8; j++ {
			temp[j][i] = in[i][j]
		}
	}
	for i := 0; i < 8; i++ {
		out[i] = logNoiseCounter(temp[i]...)
	}
	return
}

// Logarithmic addition of noise
func logNoiseCounter[T anyFloat](value /* dB(A) */ ...T) T /* dB(A) */ {
	var temp float64
	for i := 0; i < len(value); i++ {
		temp += math.Pow(10, 0.1*float64(value[i]))
	}
	return T(10 * math.Log10(temp))
}
