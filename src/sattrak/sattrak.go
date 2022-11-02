package main

import (
	"github.com/joshuaferrara/go-satellite"

	"fmt"

	"encoding/json"
  "unicode"
  "strings"
  "strconv"
	"math"
	"sync"
	"time"
)

const (
	STATUS_SAT_IN_VIEW            = 1 << 0
	STATUS_SAT_IN_SHADOW          = 1 << 7
	STATUS_HOME_LOCATION_IN_NIGHT = 1 << 1
	STATUS_HOME_LONGITUDE_EAST    = 1 << 2
)

var (
	locationTime time.Time

	chosenOMM       OMM
	chosenOMMOffset int
	ommNames        = []string{}
	omms            OMMs

	newOrbitRefreshRate      chan bool
	newOrbitRefreshRateMutex = sync.RWMutex{}
	degreesToRadiansMul      = math.Pi / 180
	twoPi                    = math.Pi * 2

	isUTC, is12Hr                        bool
	homeEnabled, sunEnabled, moonEnabled bool
	accurateMoonPhase                    bool
	isCurrentTime                        bool
	autoStartDevices                     bool

	satInputs         = []byte{}
	homeSunMoonInputs = []byte{}

	satSwapTicker *time.Ticker

	ommsMutex = sync.RWMutex{}
)

type WSPacket struct {
	PacketType string      `json:"packetType"`
	Packet     interface{} `json:"packet"`
}

func getSatAltitudeSlice(altitude float64) (altitudeBCDFinal []byte) {
	altitudeBCD := intToBCD(int64(math.Round(altitude * 100)))
	altitudeBCDLen := len(altitudeBCD)
	if altitudeBCDLen > 7 {
		exponent := intToBCD(altitudeBCDLen - 3)
		exponentLen := len(exponent)
		altitudeBCD = intToBCD(int64(math.Round(altitude /
			math.Pow(10, float64(altitudeBCDLen-8+exponentLen)))))
		altitudeBCDFinal = append(altitudeBCDFinal, altitudeBCD[0])
		altitudeBCDFinal = append(altitudeBCDFinal, 0xfe)
		altitudeBCDFinal = append(altitudeBCDFinal, altitudeBCD[1:6-exponentLen]...)
		altitudeBCDFinal = append(altitudeBCDFinal, 0x0e)
		altitudeBCDFinal = append(altitudeBCDFinal, exponent...)
	} else {
		altitudeBCD = padByteSlice(altitudeBCD, 7)
		altitudeBCDFinal = append(altitudeBCDFinal, altitudeBCD[:5]...)
		altitudeBCDFinal = append(altitudeBCDFinal, 0xfe)
		altitudeBCDFinal = append(altitudeBCDFinal, altitudeBCD[5:]...)
	}
	return
}

func setSatelliteInputs(t time.Time, timeIs12Hr bool, satLatLon satellite.LatLong, altitude float64, noradID int64) (controllerInputs []byte) {
	sec := padByteSlice(intToBCD(t.Second()), 2)
	min := padByteSlice(intToBCD(t.Minute()), 2)
	hour := t.Hour()
	var hr []byte
	if timeIs12Hr {
		hr = padByteSlice(intToBCD(hour%12), 2)
		if hr[0] == 0 && hr[1] == 0 {
			hr[0] = 1
			hr[1] = 2
		}
	} else {
		hr = padByteSlice(intToBCD(hour), 2)
	}
	day := padByteSlice(intToBCD(t.Day()), 2)
	yr := padByteSlice(intToBCD(t.Year()), 4)
	nID := padByteSlice(intToBCD(noradID), 9)
	altitudeSlice := getSatAltitudeSlice(altitude)
	satLat := padByteSlice(intToBCD(int(math.Abs(math.Round(satLatLon.Latitude*100)))), 4)
	satLon := padByteSlice(intToBCD(int(math.Abs(math.Round(satLatLon.Longitude*100)))), 5)
	var satStatus byte
	satStatus = setBitFlag(satLatLon.Latitude >= 0, satStatus, 7)
	satStatus = setBitFlag(satLatLon.Longitude >= 0, satStatus, 6)
	satStatus = setBitFlag(hour >= 12, satStatus, 2)
	satStatus = setBitFlag(is12Hr, satStatus, 1)
	_, offset := t.Zone()
	satStatus = setBitFlag(offset == 0, satStatus, 0)
	controllerInputs = []byte{0xf0}
	controllerInputs = append(controllerInputs, day...)
	controllerInputs = append(controllerInputs, uint8(t.Month()))
	controllerInputs = append(controllerInputs, yr...)
	controllerInputs = append(controllerInputs, hr...)
	controllerInputs = append(controllerInputs, min...)
	controllerInputs = append(controllerInputs, sec...)
	controllerInputs = append(controllerInputs, altitudeSlice...)
	controllerInputs = append(controllerInputs, nID...)
	controllerInputs = append(controllerInputs, satLat...)
	controllerInputs = append(controllerInputs, satLon...)
	controllerInputs = append(controllerInputs, satStatus)
	return
}

func homeSunMoonToInputs(homeLatLon, sunLatLon, moonLatLon satellite.LatLong, illumination float64, satInDaylight, satInView bool) (controllerInputs []byte) {
	homeLat := uint8(math.Abs(math.Round(homeLatLon.Latitude))+90) % 181
	homeLon := uint8(math.Abs(math.Round(homeLatLon.Longitude))) % 181

	sunLat := uint8(math.Round(sunLatLon.Latitude)+90) % 181
	sunLon := uint8(math.Abs(math.Round(sunLatLon.Longitude))) % 181

	moonLat := uint8(math.Round(moonLatLon.Latitude)+90) % 181
	moonLon := uint8(math.Abs(math.Round(moonLatLon.Longitude))) % 181
	if !accurateMoonPhase {
		illumination = 8
	}
	moonPhase := moonPhaseToSpriteOffset(illumination)
	var homeSunMoonStatus byte

	homeSunMoonStatus = setBitFlag(homeEnabled, homeSunMoonStatus, 0)
	homeSunMoonStatus = setBitFlag(sunEnabled, homeSunMoonStatus, 1)
	homeSunMoonStatus = setBitFlag(moonEnabled, homeSunMoonStatus, 2)

	homeSunMoonStatus = setBitFlag(homeLatLon.Longitude >= 0, homeSunMoonStatus, 3)
	homeSunMoonStatus = setBitFlag(sunLatLon.Longitude >= 0, homeSunMoonStatus, 4)
	homeSunMoonStatus = setBitFlag(moonLatLon.Longitude >= 0, homeSunMoonStatus, 5)

	homeSunMoonStatus = setBitFlag(satInDaylight, homeSunMoonStatus, 6)
	homeSunMoonStatus = setBitFlag(satInView, homeSunMoonStatus, 7)

	controllerInputs = []byte{0xf5}
	controllerInputs = append(controllerInputs, homeLat)
	controllerInputs = append(controllerInputs, homeLon)
	controllerInputs = append(controllerInputs, sunLat)
	controllerInputs = append(controllerInputs, sunLon)
	controllerInputs = append(controllerInputs, moonLat)
	controllerInputs = append(controllerInputs, moonLon)
	controllerInputs = append(controllerInputs, moonPhase)
	controllerInputs = append(controllerInputs, homeSunMoonStatus)
	return
}

func getLineChecksum(s string) (checksum int) {
  for _, letter := range s {
    if letter == '-' {
      checksum++
    } else if unicode.IsNumber(letter) {
      checksum += int(letter-'0')
    }
  }
  return
}

func handleNewSecondData(t time.Time, omm OMM) {
  date, _ := time.Parse("2006-01-02T15:04:05.999999", omm.EPOCH)
  f := float64(date.Hour()*3600e9+date.Minute()*60e9+date.Second()*1e9+date.Nanosecond())/float64(24*3600e9)
  line1 := "1 "
  line1 += int64ToFixedString(omm.NORAD_CAT_ID, 5)
  line1 += stringToFixedString(omm.CLASSIFICATION_TYPE, 1)
  line1 += " "
  s := stringToFixedString(strings.ReplaceAll(omm.OBJECT_ID, "-", ""), 10)
  line1 += s[2:]
  line1 += " "
  line1 += intToFixedString(date.Year(), 2)
  line1 += intToFixedString(date.YearDay(), 3)
  line1 += "." + float64ToFixedString(f, 8)
  line1 += " "
  if omm.MEAN_MOTION_DOT < 0 {
    line1 += "-"
  } else {
    line1 += " "
  }
  line1 += "." + float64ToFixedString(omm.MEAN_MOTION_DOT, 8)
  line1 += " "
  line1 += float64ToDecimalPointAssumed(omm.MEAN_MOTION_DDOT)
  line1 += " "
  line1 += float64ToDecimalPointAssumed(omm.BSTAR)
  line1 += " 0 "
  s = strconv.Itoa(int(omm.ELEMENT_SET_NO))
  for len(s) < 4 {
    s = " " + s
  }
  line1 += s
  line1 += intToFixedString(getLineChecksum(line1), 1)

  line2 := "2 "
  line2 += int64ToFixedString(omm.NORAD_CAT_ID, 5)
  line2 += " "
  i, d := math.Modf(omm.INCLINATION)
  if i < 10 {
    line2 += "  " + intToFixedString(int(i), 1)
  } else if i < 100 {
    line2 += " " + intToFixedString(int(i), 2)
  } else {
    line2 += intToFixedString(int(i), 3)
  }
  line2 += "." + float64ToFixedString(d, 4)
  line2 += " "
  i, d = math.Modf(omm.RA_OF_ASC_NODE)
  if i < 10 {
    line2 += "  " + intToFixedString(int(i), 1)
  } else if i < 100 {
    line2 += " " + intToFixedString(int(i), 2)
  } else {
    line2 += intToFixedString(int(i), 3)
  }
  line2 += "." + float64ToFixedString(d, 4)
  line2 += " "
  line2 += float64ToFixedString(omm.ECCENTRICITY, 7)
  line2 += " "
  i, d = math.Modf(omm.ARG_OF_PERICENTER)
  if i < 10 {
    line2 += "  " + intToFixedString(int(i), 1)
  } else if i < 100 {
    line2 += " " + intToFixedString(int(i), 2)
  } else {
    line2 += intToFixedString(int(i), 3)
  }
  line2 += "." + float64ToFixedString(d, 4)
  line2 += " "
  i, d = math.Modf(omm.MEAN_ANOMALY)
  if i < 10 {
    line2 += "  " + intToFixedString(int(i), 1)
  } else if i < 100 {
    line2 += " " + intToFixedString(int(i), 2)
  } else {
    line2 += intToFixedString(int(i), 3)
  }
  line2 += "." + float64ToFixedString(d, 4)
  line2 += " "
  i, d = math.Modf(omm.MEAN_MOTION)
  if i < 10 {
    line2 += " " + intToFixedString(int(i), 1)
  } else {
    line2 += intToFixedString(int(i), 2)
  }
  line2 += "." + float64ToFixedString(d, 8)
  s = strconv.Itoa(int(omm.REV_AT_EPOCH))
  for len(s) < 5 {
    s = " " + s
  }
  line2 += s
  line2 += intToFixedString(getLineChecksum(line2), 1)

	sat := satellite.TLEToSat(line1, line2, "wgs84")
	timeNow := t
	if isUTC {
		timeNow = timeNow.UTC()
	}
	yr, mon, day := timeNow.UTC().Date()
	hr, min, sec := timeNow.UTC().Clock()

	gmst := satellite.ThetaG_JD(satellite.JDay(yr, int(mon), day, hr, min, sec))
	satECI, _ := satellite.Propagate(sat, yr, int(mon), day, hr, min, sec)
	satAltitude, _, satLatLon := satellite.ECIToLLA(satECI, gmst)
	satLatLon = satellite.LatLongDeg(satLatLon)

	homeLatLon := satellite.LatLong{config.HomeLatitude, config.HomeLongitude}

	homeECI := satellite.LLAToECI(latLongRadians(homeLatLon), 0, gmst)
  sunECI, moonECI, illumination, phase := getSunMoonECIMoonIlluminationPhase(strconv.FormatInt(t.Unix(), 10), gmst)


	_, _, sunLatLon := satellite.ECIToLLA(sunECI, gmst)
	sunLatLon = satellite.LatLongDeg(sunLatLon)
	_, _, moonLatLon := satellite.ECIToLLA(moonECI, gmst)
	moonLatLon = satellite.LatLongDeg(moonLatLon)

	homeToSatLine := diffBetweenECIs(homeECI, satECI)
	satToSunLine := diffBetweenECIs(satECI, sunECI)

	distanceECIToHome := math.Sqrt(homeECI.X*homeECI.X + homeECI.Y*homeECI.Y +
		homeECI.Z*homeECI.Z)
	distanceECIToSat := math.Sqrt(satECI.X*satECI.X + satECI.Y*satECI.Y +
		satECI.Z*satECI.Z)
	distanceECIToSun := math.Sqrt(sunECI.X*sunECI.X + sunECI.Y*sunECI.Y +
		sunECI.Z*sunECI.Z)
	distanceHomeToSat := math.Sqrt(homeToSatLine.X*homeToSatLine.X +
		homeToSatLine.Y*homeToSatLine.Y + homeToSatLine.Z*homeToSatLine.Z)
	distanceSatToSun := math.Sqrt(satToSunLine.X*satToSunLine.X +
		satToSunLine.Y*satToSunLine.Y + satToSunLine.Z*satToSunLine.Z)

	angleECIHomeSat := radiansToDegrees(math.Acos(
		(distanceHomeToSat*distanceHomeToSat +
			distanceECIToHome*distanceECIToHome -
			distanceECIToSat*distanceECIToSat) /
			(2 * distanceHomeToSat * distanceECIToHome)))
	angleSatECISun := radiansToDegrees(math.Acos(
		(distanceECIToSat*distanceECIToSat +
			distanceECIToSun*distanceECIToSun -
			distanceSatToSun*distanceSatToSun) /
			(2 * distanceECIToSat * distanceECIToSun)))
	angleSatSunECI := radiansToDegrees(math.Acos(
		(distanceSatToSun*distanceSatToSun +
			distanceECIToSun*distanceECIToSun -
			distanceECIToSat*distanceECIToSat) /
			(2 * distanceSatToSun * distanceECIToSun)))

	angleEarthSunECI := radiansToDegrees(
		math.Atan(EARTH_SEMI_MAJOR_AXIS / distanceECIToSun))

	satInDaylight := angleSatECISun < 90 || (angleSatECISun >= 90 && angleSatSunECI > angleEarthSunECI)
	satInView := angleECIHomeSat >= 90

	utc := isUTC
  moonSprite := illumination / 2
  if phase < 0.5 {
    moonSprite = 1 - moonSprite
  }
	satInputs = setSatelliteInputs(timeNow, is12Hr, satLatLon, satAltitude,
		omm.NORAD_CAT_ID)
	homeSunMoonInputs = homeSunMoonToInputs(homeLatLon, sunLatLon, moonLatLon,
		moonSprite, satInDaylight, satInView)

	infoDataMutex.Lock()
	infoData.SatelliteName = omm.OBJECT_NAME
	infoData.NoradID = omm.NORAD_CAT_ID
	infoData.Time = timeNow
	infoData.Altitude = satAltitude
	infoData.SatelliteLatitude = satLatLon.Latitude
	infoData.SatelliteLongitude = satLatLon.Longitude
	infoData.SatelliteDaylight = satInDaylight
	infoData.HomeLatitude = homeLatLon.Latitude
	infoData.HomeLongitude = homeLatLon.Longitude
	infoData.SunLatitude = sunLatLon.Latitude
	infoData.SunLongitude = sunLatLon.Longitude
	infoData.MoonLatitude = moonLatLon.Latitude
	infoData.MoonLongitude = moonLatLon.Longitude
  infoData.MoonIllumination = illumination
	infoData.IsUTC = utc
	infoData.Is12Hr = is12Hr
	infoData.HomeEnabled = config.EnableHome
	infoData.SunEnabled = config.EnableSun
	infoData.MoonEnabled = config.EnableMoon
	if config.AccurateMoonPhase {
		infoData.MoonPhase = modDegrees(phase) / 360
	} else {
		infoData.MoonPhase = 0
	}
	packet := WSPacket{"info", infoData}
	wsContent, _ := json.Marshal(packet)
	sendAllWS(wsContent)
	infoDataMutex.Unlock()
}

func degreesToRadians(deg float64) float64 {
	return math.Mod(deg*degreesToRadiansMul, twoPi)
}

func radiansToDegrees(rad float64) float64 {
	return math.Mod(rad/degreesToRadiansMul, 360)
}

func modDegrees(i float64) float64 {
	return math.Mod(math.Mod(i, 360)+360, 360)

}

func mod2Pi(i float64) float64 {
	return math.Mod(math.Mod(i, twoPi)+twoPi, twoPi)
}

func main() {
	wsConns = make(map[int]*wsConn)
	fmt.Println("READING CONFIG")
	readConfigFile()
	fmt.Println("READING CONFIG DONE")
	newOrbitRefreshRate = make(chan bool)
	go ommFileDownloadTimer()
	fmt.Println("HANDLE OMM FILE")
	handleOMMFile(false)
	fmt.Println("HANDLE OMM FILE DONE")
	fmt.Println("SELECT OMM")
	setChosenOMM(config.DefaultSatellite)
	fmt.Println("SELECT OMM DONE")

	go devicesRestartTimer()
	go getTAStm32COMPaths()

	go setHomeSatSunMoonInputsLoop()

  //setTimeDif(time.Unix(1663251565, 0))

	go switchOMM()
	go httpListen()
	startServer()
}
