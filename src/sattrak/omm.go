package main

import (
	"github.com/joshuaferrara/go-satellite"

	"crypto/rand"
	"encoding/xml"
	"math/big"
	"strings"
	//"time"
)

type OMM struct {
	OBJECT_NAME         string  `xml:"body>segment>metadata>OBJECT_NAME"`
	OBJECT_ID           string  `xml:"body>segment>metadata>OBJECT_ID"`
	CENTER_NAME         string  `xml:"body>segment>metadata>CENTER_NAME"`
	EPOCH               string  `xml:"body>segment>data>meanElements>EPOCH"`
	MEAN_MOTION         float64 `xml:"body>segment>data>meanElements>MEAN_MOTION"`
	ECCENTRICITY        float64 `xml:"body>segment>data>meanElements>ECCENTRICITY"`
	INCLINATION         float64 `xml:"body>segment>data>meanElements>INCLINATION"`
	RA_OF_ASC_NODE      float64 `xml:"body>segment>data>meanElements>RA_OF_ASC_NODE"`
	ARG_OF_PERICENTER   float64 `xml:"body>segment>data>meanElements>ARG_OF_PERICENTER"`
	MEAN_ANOMALY        float64 `xml:"body>segment>data>meanElements>MEAN_ANOMALY"`
	EPHEMERIS_TYPE      int64   `xml:"body>segment>data>tleParameters>EPHEMERIS_TYPE"`
	CLASSIFICATION_TYPE string  `xml:"body>segment>data>tleParameters>CLASSIFICATION_TYPE"`
	NORAD_CAT_ID        int64   `xml:"body>segment>data>tleParameters>NORAD_CAT_ID"`
	ELEMENT_SET_NO      int64   `xml:"body>segment>data>tleParameters>ELEMENT_SET_NO"`
	REV_AT_EPOCH        int64   `xml:"body>segment>data>tleParameters>REV_AT_EPOCH"`
	BSTAR               float64 `xml:"body>segment>data>tleParameters>BSTAR"`
	MEAN_MOTION_DOT     float64 `xml:"body>segment>data>tleParameters>MEAN_MOTION_DOT"`
	MEAN_MOTION_DDOT    float64 `xml:"body>segment>data>tleParameters>MEAN_MOTION_DDOT"`
}

type OMMs struct {
	XMLName xml.Name `xml:"ndm"`
	OMMs    []OMM    `xml:"omm"`
}

func ommToSatelliteOMM(omm OMM) (sat satellite.Satellite) {
  /*
	satOMM.ObjectName = omm.OBJECT_NAME
	satOMM.ObjectID = omm.OBJECT_ID
	satOMM.CenterName = omm.CENTER_NAME
	satOMM.Epoch, _ = time.Parse(getSubsecondFormat(omm.EPOCH), omm.EPOCH)
	satOMM.MeanMotion = omm.MEAN_MOTION
	satOMM.Eccentricity = omm.ECCENTRICITY
	satOMM.Inclination = omm.INCLINATION
	satOMM.RAOfAscNode = omm.RA_OF_ASC_NODE
	satOMM.ArgOfPericenter = omm.ARG_OF_PERICENTER
	satOMM.MeanAnomaly = omm.MEAN_ANOMALY
	satOMM.EphemerisType = omm.EPHEMERIS_TYPE
	satOMM.ClassificationType = omm.CLASSIFICATION_TYPE
	satOMM.NORADCatID = omm.NORAD_CAT_ID
	satOMM.ElementSetNo = omm.ELEMENT_SET_NO
	satOMM.RevAtEpoch = omm.REV_AT_EPOCH
	satOMM.BStar = omm.BSTAR
	satOMM.MeanMotionDot = omm.MEAN_MOTION_DOT
	satOMM.MeanMotionDdot = omm.MEAN_MOTION_DDOT
  */
	return
}

func getSubsecondFormat(timeString string) (timeFormat string) {
	timeFormat = "2006-01-02T15:04:05"
	x := strings.LastIndex(timeString, ".")
	if x > -1 {
		timeFormat += "."
		for x < len(timeString)-1 {
			timeFormat += "0"
			x++
		}
	}
	return
}

func setChosenOMM(noradIDVal int64) {

	ommsMutex.Lock()
	ommsListTmp := omms.OMMs[:]
	ommsMutex.Unlock()

	ommLen := len(ommsListTmp)

	if noradIDVal > 0 {
		ommIndex, noradID := searchSortedOMMList(noradIDVal, 0, ommLen, ommsListTmp)
		if noradID == noradIDVal {
			chosenOMM = ommsListTmp[ommIndex]
			return
		}
	}
	if ommLen > 0 {
		offset, _ := rand.Int(rand.Reader, big.NewInt(int64(ommLen-1)))
		chosenOMM = ommsListTmp[int(offset.Int64())]
	}
}

func searchSortedOMMList(noradIDVal int64, start, length int, ommsList []OMM) (ommIndex int, noradID int64) {
	if length > 0 {
		indexPoint := (start + (start + length - 1)) / 2
		noradIDAtIndexPoint := ommsList[indexPoint].NORAD_CAT_ID
		if noradIDVal == noradIDAtIndexPoint {
			ommIndex = indexPoint
			noradID = noradIDAtIndexPoint
		} else if noradIDVal < noradIDAtIndexPoint {
			ommIndex, noradID = searchSortedOMMList(noradIDVal, start, indexPoint-start, ommsList)
		} else {
			ommIndex, noradID = searchSortedOMMList(noradIDVal, indexPoint+1, length-(indexPoint+1-start), ommsList)
		}
	}
	return
}
