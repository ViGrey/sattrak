package main

import (
	"strconv"
	"strings"
)

type SatellitePacket struct {
	PacketType string     `json:"packetType"`
	Packet     DeviceData `json:"packet"`
}

func searchSatName(name string, maxLen int) (satList []string) {
	ommsMutex.Lock()
	ommsListTmp := omms.OMMs[:]
	ommsMutex.Unlock()

	var satListPerfectMatch []string
	var satListStartMatch []string
	var satListInnerMatch []string
	name = strings.ToLower(name)

	for _, omm := range ommsListTmp {
		satName := strings.ToLower(omm.OBJECT_NAME)
		ind := strings.Index(satName, name)
		if ind >= 0 {
			if satName == name {
				satListPerfectMatch = append(satListPerfectMatch, omm.OBJECT_NAME+" (NORAD ID: "+strconv.FormatInt(omm.NORAD_CAT_ID, 10)+")")
			} else if ind == 0 {
				satListStartMatch = append(satListStartMatch, omm.OBJECT_NAME+" (NORAD ID: "+strconv.FormatInt(omm.NORAD_CAT_ID, 10)+")")
			} else {
				satListInnerMatch = append(satListInnerMatch, omm.OBJECT_NAME+" (NORAD ID: "+strconv.FormatInt(omm.NORAD_CAT_ID, 10)+")")
			}
		}
	}
	satList = append(satList, satListPerfectMatch...)
	satList = append(satList, satListStartMatch...)
	satList = append(satList, satListInnerMatch...)
	if maxLen > 0 {
		satList = satList[:maxLen]
	}
	return
}
