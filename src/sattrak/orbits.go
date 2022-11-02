package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"sort"
	"time"
)

var (
	orbitURLList          []string
	ommFileDownloadTicker *time.Ticker
)

func ommFileDownloadTimer() {
	ommFileDownloadTicker = time.NewTicker(time.Duration(config.OrbitRefreshRate) * time.Second)
	for {
		select {
		case <-ommFileDownloadTicker.C:
			handleOMMFile(true)
		case <-newOrbitRefreshRate:
			ommFileDownloadTicker = time.NewTicker(time.Duration(config.OrbitRefreshRate) * time.Second)
		}
	}
}

func handleOMMFile(force bool) {
	f, err := os.Stat(config.OrbitCachePath)
	if err == nil {
		if force || time.Now().Sub(f.ModTime()) >= time.Duration(config.OrbitRefreshRate)*time.Second {
			downloadOMMFile()
		}
		xmlDat, _ := os.ReadFile(config.OrbitCachePath)

		ommsNew := new(OMMs)
		xml.Unmarshal(xmlDat, ommsNew)
		sortOMMs(*ommsNew)
	} else {
		downloadOMMFile()
	}
}

func downloadOMMFile() {
	ommsNew := new(OMMs)
	sats := make(map[int64]bool)
	getOrbitSourceList()
	writeCacheFile := true
	for _, url := range orbitURLList {
		for attempt := 0; attempt < 3 && len(url) > 0; attempt++ {
			client := http.Client{
				Timeout: 5 * time.Second,
			}
			resp, err := client.Get(url)
			if err == nil {
				defer resp.Body.Close()
				content, err := io.ReadAll(resp.Body)
				// SUCCESSFUL CONTENT READ
				if err == nil {
					ommsTmp := new(OMMs)
					xml.Unmarshal(content, &ommsTmp)
					for _, omm := range ommsTmp.OMMs {
						if sats[omm.NORAD_CAT_ID] == false {
							sats[omm.NORAD_CAT_ID] = true
							ommsNew.OMMs = append(ommsNew.OMMs, omm)
						}
					}
					break
				}
			}
			if attempt == 2 {
				writeCacheFile = false
				break
			}
		}
		if !writeCacheFile {
			break
		}
	}
	if writeCacheFile {
		sortOMMs(*ommsNew)
		ommsCacheContent, _ := xml.Marshal(ommsNew)
		err := os.WriteFile(config.OrbitCachePath+"-bak", ommsCacheContent, 0644)
		if err == nil {
			os.Rename(config.OrbitCachePath+"-bak", config.OrbitCachePath)
		}
	}
}

func sortOMMs(ommsList OMMs) {
	sort.Slice(ommsList.OMMs, func(i, j int) bool {
		return ommsList.OMMs[i].NORAD_CAT_ID < ommsList.OMMs[j].NORAD_CAT_ID
	})
	ommsMutex.Lock()
	omms.OMMs = ommsList.OMMs[:]
	ommsMutex.Unlock()
}
