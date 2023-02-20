package HVAC

import "math"

type (
	Noise [8]float32
)

func (s Noise) TotalNoise() float32 {
	var isZero bool
	for i := 0; i < 8; i++ {
		isZero = isZero || s[i] <= .01
	}
	if isZero {
		return 0
	} else {
		return logNoiseCounter(s[0], s[1], s[2], s[3], s[4], s[5], s[6], s[7])
	}

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
