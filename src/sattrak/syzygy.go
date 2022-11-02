package main

import (
	"github.com/joshuaferrara/go-satellite"

  "math"
  "encoding/json"
  "os/exec"
)

type Syzygy struct {
  MoonECEF satellite.Vector3 `json:"moonECEF"`
  SunECEF satellite.Vector3 `json:"sunECEF"`
  MoonIllumination float64 `json:"moonIllumination"`
  MoonPhase float64 `json:"moonPhase"`
}

type Vector struct {
  X float64 `json:"x"`
  Y float64 `json:"y"`
  Z float64 `json:"z"`
}

func getSunMoonECIMoonIlluminationPhase(unixTimestamp string, gmst float64) (sunECI, moonECI satellite.Vector3, moonIllumination, moonPhase float64) {
	out, err := exec.Command("syzygy", unixTimestamp).Output()
	if err != nil {
		return
	}
  s := new(Syzygy)
	json.Unmarshal(out, &s)
  moonIllumination = s.MoonIllumination
  moonPhase = s.MoonPhase
  sunECI = ecefToECI(s.SunECEF, gmst)
  moonECI = ecefToECI(s.MoonECEF, gmst)
  return
}

func moonPhaseToSpriteOffset(i float64) uint8 {
  return uint8((int(math.Round(i*16))%16+16)%16)
}
