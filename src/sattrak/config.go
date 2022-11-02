package main

import (
	"github.com/go-yaml/yaml"

	"os"
	"strings"
	"time"
)

var (
	config     = Config{}
	basePath   = os.Getenv("HOME") + "/.sattrak"
	configFile = basePath + "/config.yaml"
)

type Config struct {
	EditKey                 string  `yaml:"edit-key,omitempty"`
	DefaultSatellite        int64   `yaml:"default-satellite"`
	SatelliteSwapRate       int     `yaml:"satellite-swap-rate"`
	EnableSatelliteSwap     bool    `yaml:"enable-satellite-swap"`
	OrbitRefreshRate        int     `yaml:"orbit-refresh-rate"`
	OrbitSourceListPath     string  `yaml:"orbit-src-list-path"`
	OrbitCachePath          string  `yaml:"orbit-cache-path"`
	HomeLatitude            float64 `yaml:"home-latitude"`
	HomeLongitude           float64 `yaml:"home-longitude"`
	EnableHome              bool    `yaml:"enable-home"`
	EnableSun               bool    `yaml:"enable-sun"`
	EnableMoon              bool    `yaml:"enable-moon"`
	AccurateMoonPhase       bool    `yaml:"accurate-moon-phase"`
	TimeUTC                 bool    `yaml:"time-utc"`
	Time12Hr                bool    `yaml:"time-12-hr"`
	StartTAStm32Immediately bool    `yaml:"start-tastm32-immediately"`
}

type ConfigPacket struct {
	PacketType string `yaml:"packetType"`
	Packet     Config `yaml:"packet"`
}

func readConfigFile() {
	os.Mkdir(basePath, 0755)
	content, _ := os.ReadFile(configFile)
	yaml.Unmarshal(content, &config)
	if config.OrbitRefreshRate == 0 {
		config.OrbitRefreshRate = 86400
	}
	if config.OrbitSourceListPath == "" {
		config.OrbitSourceListPath = "$BASE/cache/orbits-list.txt"
	}
	if config.OrbitCachePath == "" {
		config.OrbitCachePath = "$BASE/cache/orbits.xml"
	}
	config.OrbitSourceListPath = strings.ReplaceAll(config.OrbitSourceListPath, "$BASE", basePath)
	config.OrbitCachePath = strings.ReplaceAll(config.OrbitCachePath, "$BASE", basePath)
	isUTC = config.TimeUTC
	is12Hr = config.Time12Hr
	homeEnabled = config.EnableHome
	sunEnabled = config.EnableSun
	moonEnabled = config.EnableMoon
	accurateMoonPhase = config.AccurateMoonPhase
	autoStartDevices = config.StartTAStm32Immediately
}

func editConfig(content Config) {
	validConfigKey := true
	if content.EditKey == "default-satellite" {
		if content.DefaultSatellite > 0 {
			config.DefaultSatellite = content.DefaultSatellite
			setChosenOMM(content.DefaultSatellite)
		} else {
			setChosenOMM(0)
		}
		if config.SatelliteSwapRate > 0 && config.EnableSatelliteSwap {
			satSwapTicker.Reset(time.Duration(config.SatelliteSwapRate) * time.Second)
		} else {
			satSwapTicker.Stop()
		}
	} else if content.EditKey == "enable-satellite-swap" {
		config.EnableSatelliteSwap = content.EnableSatelliteSwap
		if config.SatelliteSwapRate > 0 && config.EnableSatelliteSwap {
			satSwapTicker.Reset(time.Duration(config.SatelliteSwapRate) * time.Second)
		} else {
			satSwapTicker.Stop()
		}
	} else if content.EditKey == "enable-home" {
		config.EnableHome = content.EnableHome
		homeEnabled = content.EnableHome
	} else if content.EditKey == "enable-sun" {
		config.EnableSun = content.EnableSun
		sunEnabled = content.EnableSun
	} else if content.EditKey == "enable-moon" {
		config.EnableMoon = content.EnableMoon
		moonEnabled = content.EnableMoon
	} else if content.EditKey == "accurate-moon-phase" {
		config.AccurateMoonPhase = content.AccurateMoonPhase
		accurateMoonPhase = content.AccurateMoonPhase
	} else if content.EditKey == "time-utc" {
		config.TimeUTC = content.TimeUTC
		isUTC = content.TimeUTC
	} else if content.EditKey == "time-12-hr" {
		config.Time12Hr = content.Time12Hr
		is12Hr = content.Time12Hr
	} else if content.EditKey == "start-tastm32-immediately" {
		config.StartTAStm32Immediately = content.StartTAStm32Immediately
		autoStartDevices = content.StartTAStm32Immediately
	} else if content.EditKey == "satellite-swap-rate" {
		if content.SatelliteSwapRate > 0 {
			config.SatelliteSwapRate = content.SatelliteSwapRate
			if config.SatelliteSwapRate > 0 && config.EnableSatelliteSwap {
				satSwapTicker.Reset(time.Duration(config.SatelliteSwapRate) * time.Second)
			}
		} else {
			validConfigKey = false
		}
	} else if content.EditKey == "home-latitude" {
		if content.HomeLatitude >= -90 || content.HomeLatitude <= 90 {
			config.HomeLatitude = content.HomeLatitude
		} else {
			validConfigKey = false
		}
	} else if content.EditKey == "home-longitude" {
		if content.HomeLongitude >= -180 || content.HomeLongitude <= 180 {
			config.HomeLongitude = content.HomeLongitude
		} else {
			validConfigKey = false
		}
	} else {
		validConfigKey = false
	}
	if validConfigKey {
		writeConfig()
	}
}

func writeConfig() {
	config.OrbitSourceListPath = strings.ReplaceAll(config.OrbitSourceListPath, basePath, "$BASE")
	config.OrbitCachePath = strings.ReplaceAll(config.OrbitCachePath, basePath, "$BASE")
	content, _ := yaml.Marshal(config)
	os.WriteFile(configFile, content, 0655)
}
