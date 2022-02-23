package square

import (
	"math"
)

type num byte;

const (
	SidesTriangle num = 3 
	SidesSquare num = 4 
	SidesCircle num = 0
)

func CalcSquare(sideLen float64, sidesNum num) float64 {
	if(sidesNum==3) {
		return (math.Sqrt(3)/4)*sideLen*sideLen
	}
	if(sidesNum==4) {
		return sideLen*sideLen
	}
	if(sidesNum==0) {
		return math.Pi*sideLen*sideLen
	}
	return 0
}