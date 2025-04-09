package utils

import (
	"math"
	"math/cmplx"
)

func RotateDegrees(zNumber complex128, thetaDegrees float64) complex128 {
	var thetaRad = thetaDegrees * math.Pi / 180

	sinTheta := math.Sin(thetaRad)
	cosTheta := math.Cos(thetaRad)

	var rotation complex128 = cmplx.Rect(cosTheta, sinTheta)

	return zNumber * rotation
}
