package main

import (
	"github.com/joshuaferrara/go-satellite"

	"math"
)

const (
	EARTH_SEMI_MAJOR_AXIS float64 = 6378.137
	EARTH_SEMI_MINOR_AXIS float64 = 6356.7523142
)
/*

func eciToLLA(eciCoords satellite.Vector3, gmst float64) (altitude float64, latLon satellite.LatLong) {
	f := (EARTH_SEMI_MAJOR_AXIS - EARTH_SEMI_MINOR_AXIS) / EARTH_SEMI_MAJOR_AXIS
	e2 := (2*f - f*f)
	xyLen := math.Sqrt(eciCoords.X*eciCoords.X + eciCoords.Y*eciCoords.Y)

	latLon.Longitude = math.Atan2(eciCoords.Y, eciCoords.X) - gmst
	latitude := math.Atan2(eciCoords.Z, xyLen)
	c := 0.0
	for i := 0; i < 20; i++ {
		c = 1 / math.Sqrt(1-e2*(math.Sin(latitude)*math.Sin(latitude)))
		latitude = math.Atan2(eciCoords.Z+(EARTH_SEMI_MAJOR_AXIS*c*e2*math.Sin(latitude)), xyLen)
	}
	altitude = (xyLen / math.Cos(latitude))
	latLon.Latitude = latitude
	return
}

func llaToECI(latLonRad satellite.LatLong, altitude, gmst float64) (eciCoords satellite.Vector3) {
	theta := gmst + latLonRad.Longitude
	pa := (EARTH_SEMI_MINOR_AXIS + altitude) * math.Cos(latLonRad.Latitude)
	eciCoords.X = pa * math.Cos(theta)
	eciCoords.Y = pa * math.Sin(theta)
	eciCoords.Z = (EARTH_SEMI_MAJOR_AXIS + altitude) * math.Sin(latLonRad.Latitude)
	return
}
*/

func latLongRadians(latLon satellite.LatLong) (latLonRad satellite.LatLong) {
	latLonRad.Latitude = latLon.Latitude * degreesToRadiansMul
	latLonRad.Longitude = latLon.Longitude * degreesToRadiansMul
	return
}

func diffBetweenECIs(eciA, eciB satellite.Vector3) (eciAB satellite.Vector3) {
	eciAB.X = eciA.X - eciB.X
	eciAB.Y = eciA.Y - eciB.Y
	eciAB.Z = eciA.Z - eciB.Z
	return
}

func ecefToECI(ecef satellite.Vector3, gmst float64) (eci satellite.Vector3) {
  eci.X = (ecef.X * math.Cos(-gmst)) + (ecef.Y*math.Sin(-gmst))
  eci.Y = (ecef.X * -math.Sin(-gmst)) + (ecef.Y*math.Cos(-gmst))
  eci.Z = ecef.Z
  return
}
