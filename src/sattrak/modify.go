package main

import (
	"encoding/json"
)

type ModifyPacket struct {
	PacketType string     `json:"packetType"`
	Packet     ModifyData `json:"packet"`
}

type ModifyData struct {
	DefaultSatellite        int64   `json:"default-satellite"`
	DefaultSatelliteName    string  `json:"default-satellite-name"`
	SatelliteSwapRate       int     `json:"satellite-swap-rate"`
	EnableSatelliteSwap     bool    `json:"enable-satellite-swap"`
	OrbitRefreshRate        int     `json:"orbit-refresh-rate"`
	OrbitSourceListPath     string  `json:"orbit-src-list-path"`
	OrbitCachePath          string  `json:"orbit-cache-path"`
	HomeLatitude            float64 `json:"home-latitude"`
	HomeLongitude           float64 `json:"home-longitude"`
	EnableHome              bool    `json:"enable-home"`
	EnableSun               bool    `json:"enable-sun"`
	EnableMoon              bool    `json:"enable-moon"`
	AccurateMoonPhase       bool    `json:"accurate-moon-phase"`
	TimeUTC                 bool    `json:"time-utc"`
	Time12Hr                bool    `json:"time-12-hr"`
	StartTAStm32Immediately bool    `json:"start-tastm32-immediately"`
}

func getModifyDataValues() (wsContent []byte) {
	ommsMutex.Lock()
	ommsListTmp := omms.OMMs[:]
	ommsMutex.Unlock()

	modifyData := ModifyData{}
	modifyData.DefaultSatellite = config.DefaultSatellite
	satOMMIndex, _ := searchSortedOMMList(config.DefaultSatellite, 0, len(ommsListTmp), ommsListTmp)
	modifyData.DefaultSatellite = config.DefaultSatellite

	modifyData.DefaultSatelliteName = ommsListTmp[satOMMIndex].OBJECT_NAME

	modifyData.SatelliteSwapRate = config.SatelliteSwapRate
	modifyData.EnableSatelliteSwap = config.EnableSatelliteSwap
	modifyData.OrbitRefreshRate = config.OrbitRefreshRate
	modifyData.OrbitSourceListPath = config.OrbitSourceListPath
	modifyData.OrbitCachePath = config.OrbitCachePath
	modifyData.HomeLatitude = config.HomeLatitude
	modifyData.HomeLongitude = config.HomeLongitude
	modifyData.EnableHome = config.EnableHome
	modifyData.EnableSun = config.EnableSun
	modifyData.EnableMoon = config.EnableMoon
	modifyData.AccurateMoonPhase = config.AccurateMoonPhase
	//modifyData.TimeCurrent = config.TimeCurrent
	modifyData.TimeUTC = config.TimeUTC
	modifyData.Time12Hr = config.Time12Hr
	//modifyData.StartTime = config.StartTime
	modifyData.StartTAStm32Immediately = config.StartTAStm32Immediately
	packet := WSPacket{"modify", modifyData}
	wsContent, _ = json.Marshal(packet)
	return
}

func sendModifyDataValues() {
	wsContent := getModifyDataValues()
	sendAllWS(wsContent)
}
