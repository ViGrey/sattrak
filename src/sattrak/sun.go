package main
/*

import (
	"github.com/joshuaferrara/go-satellite"

	"math"
	"time"
)

// (28.2) Astronomical Algorithms p 183
// Sun Mean Longitude
func getSunMeanLongitude(t float64) float64 {
	//a := 280.4606184 + t*36000.77005361
	//return a
	return (280.466567 + t*36000.76982779 + t*t*0.0003032028 +
		t*t*t/49.931 - t*t*t*t/1.53 - t*t*t*t*t/20)
}

// (47.3) Astronomical Algorithms p 338
// Sun Mean Anomaly
func getSunMeanAnomaly(t float64) float64 {
	return (357.5291092 + t*35999.0502909 - t*t*0.0001536 +
		t*t*t/24490000)
}

func getSunEccentricity(t float64) float64 {
	//(46.8093/3600)
	return 23.43929 - t*0.01300258333
}

func getSunLocation(t time.Time) (eciCoords satellite.Vector3) {
	year, month, day := t.UTC().Date()
	hour, minute, second := t.UTC().Clock()
	jday := satellite.JDay(year, int(month), day, hour, minute, second)
	jC := getJ2000Century(jday)

	meanLongitude := getSunMeanLongitude(jC)
	meanAnomaly := getSunMeanAnomaly(jC)
	meanAnomalyRad := degreesToRadians(meanAnomaly)

	eclipticLongitude := (meanLongitude + 1.914666471*(math.Sin(meanAnomalyRad)) +
		(0.918994643 * math.Sin(2*meanAnomalyRad)))
	eclipticLongitudeRad := degreesToRadians(eclipticLongitude)

	eccentricity := getSunEccentricity(jC)
	eccentricityRad := degreesToRadians(eccentricity)
	var auInKM float64 = 149597870.700
	eciCoords.X = math.Cos(eclipticLongitudeRad) * auInKM
	eciCoords.Y = math.Cos(eccentricityRad) * math.Sin(eclipticLongitudeRad) * auInKM
	eciCoords.Z = math.Sin(eccentricityRad) * math.Sin(eclipticLongitudeRad) * auInKM
	return
}
*/
