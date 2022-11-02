package main

import (
	"time"
)

var (
	lastCheckedTimeUnix int64
	timeDif             time.Duration
)

func checkNewSecond(t time.Time) bool {
	timeUnix := t.Unix()
	if timeUnix > lastCheckedTimeUnix {
		lastCheckedTimeUnix = timeUnix
		return true
	}
	return false
}

func setHomeSatSunMoonInputsLoop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			t := time.Now().Add(timeDif)
			if checkNewSecond(t) {
				handleNewSecondData(t, chosenOMM)

				writeToAllConns(homeSunMoonInputs)
				writeToAllConns(satInputs)

				writeToAllDevices(homeSunMoonInputs)
				writeToAllDevices(satInputs)
			}
		}
	}
}

func setTimeDif(t time.Time) {
	timeDif = t.UTC().Sub(time.Now().UTC())
	lastCheckedTimeUnix = 0
}

func switchOMM() {
	satSwapTicker = time.NewTicker(1 * time.Second)
	satSwapRate := config.SatelliteSwapRate
	if satSwapRate > 0 && config.EnableSatelliteSwap {
		satSwapTicker.Reset(time.Duration(config.SatelliteSwapRate) * time.Second)
	} else {
		satSwapTicker.Stop()
	}
	for {
		select {
		case <-satSwapTicker.C:
			setChosenOMM(0)
		}
	}
}
