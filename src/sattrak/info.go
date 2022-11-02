package main

import (
	"sync"
	"time"
)

var (
	infoData      InfoData
	infoDataMutex = sync.RWMutex{}
)

type InfoPacket struct {
	PacketType string   `json:"packetType"`
	Packet     InfoData `json:"packet"`
}

type InfoData struct {
	SatelliteName      string    `json:"satelliteName"`
	Time               time.Time `json:"time"`
	Altitude           float64   `json:"altitude"`
	NoradID            int64     `json:"noradID"`
	SatelliteLatitude  float64   `json:"satelliteLatitude"`
	SatelliteLongitude float64   `json:"satelliteLongitude"`
	SatelliteDaylight  bool      `json:"satelliteDaylight"`
	HomeLatitude       float64   `json:"homeLatitude"`
	HomeLongitude      float64   `json:"homeLongitude"`
	SunLatitude        float64   `json:"sunLatitude"`
	SunLongitude       float64   `json:"sunLongitude"`
	MoonLatitude       float64   `json:"moonLatitude"`
	MoonLongitude      float64   `json:"moonLongitude"`
	MoonPhase          float64   `json:"moonPhase"`
  MoonIllumination   float64   `json:"moonIllumination"`
	IsUTC              bool      `json:"isUTC"`
	Is12Hr             bool      `json:"is12Hr"`
	HomeEnabled        bool      `json:"homeEnabled"`
	SunEnabled         bool      `json:"sunEnabled"`
	MoonEnabled        bool      `json:"moonEnabled"`
}
