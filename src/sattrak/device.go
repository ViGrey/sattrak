package main

import (
	"github.com/mikepb/go-serial"

	"bytes"
	"encoding/json"
	"sync"
	"time"
)

const (
	DEVICE_READY    = 0
	DEVICE_IN_USE   = 1
	DEVICE_DFU_MODE = 2
)

var (
	deviceVIDPID = map[string][]int{
		"TAStm32": []int{0xb07, 0x7a5},
	}
	readBuffer      []byte
	readPause       bool
	readWait        chan bool
	indexBase       int
	devices         = map[*Device]*sync.RWMutex{}
	devicesMutex    = sync.RWMutex{}
	deviceData      DeviceData
	deviceDataMutex = sync.RWMutex{}
)

type DevicesPacket struct {
	PacketType string     `json:"packetType"`
	Packet     DeviceData `json:"packet"`
}

type Device struct {
	device       *serial.Port `json:"-"`
	Name         string       `json:"name,omitempty"`
	Path         string       `json:"path,omitempty"`
	Status       int          `json:"status"`
	Index        int          `json:"index"`
	Alive        bool         `json:"alive"`
	resetWaiting bool         `json:"-"`
	setupWaiting bool         `json:"-"`
	devR         chan bool    `json:"-"`
	devS         chan bool    `json:"-"`
}

type DeviceData struct {
	Devices []*Device `json:"devices"`
}

func getTAStm32COMPaths() {
	for {
		infoList, err := serial.ListPorts()
		if err == nil {
			for _, list := range infoList {
				vid, pid, _ := list.USBVIDPID()
				for trd, y := range deviceVIDPID {
					if y[0] == vid && y[1] == pid {
						go addDevice(trd, list.Name())
					}
				}
			}
		}
		go func() {
			packet := WSPacket{"devices", deviceData}
			wsContent, _ := json.Marshal(packet)
			sendAllWS(wsContent)
		}()
		time.Sleep(1 * time.Second)
	}
}

func addDevice(trd, devPath string) {
	newDev := true
	for _, d := range deviceData.Devices {
		if d.Path == devPath {
			newDev = false
			break
		}
	}
	if newDev {
		d := new(Device)
		options := serial.RawOptions
		options.Mode = serial.MODE_READ_WRITE
		options.BitRate = 115200
		device, err := options.Open(devPath)
		if err != nil {
			return
		}
		d.device = device
		d.Name = trd
		d.Path = devPath
		d.Status = DEVICE_READY
		d.Alive = false
		if config.StartTAStm32Immediately {
			d.Status = DEVICE_IN_USE
		}
		d.Index = indexBase
		devices[d] = &sync.RWMutex{}

		devicesMutex.Lock()
		deviceData.Devices = append(deviceData.Devices, d)
		devicesMutex.Unlock()

		go deviceRead(d)

		deviceReset(d)
		deviceSetup(d)

		indexBase++
		writeToDevice(d, satInputs)
		writeToDevice(d, homeSunMoonInputs)
	}
}

func writeToDevice(d *Device, content []byte) {
	if d.Status != DEVICE_IN_USE || d.Alive == false {
		return
	}
	var final []byte
	for _, x := range content {
		final = append(final, []byte{'A', x}...)
	}
	if len(final) > 0 {
		devices[d].Lock()
		d.device.Write(final)
		devices[d].Unlock()
	}
}

func deviceRead(d *Device) {
	for {
		d.device.SetDeadline(time.Now().Add(100 * time.Millisecond))
		buf := make([]byte, 4096)
		bufN, err := d.device.Read(buf)
		if bufN > 0 {
			if bytes.ContainsRune(buf[:bufN], 'A') {
				if d.Status == DEVICE_IN_USE && !d.Alive {
					d.Alive = true
				}
			}
			if bytes.ContainsRune(buf[:bufN], 'R') {
				if d.resetWaiting {
					d.resetWaiting = false
					d.devR <- true
				}
			}
			if bytes.ContainsRune(buf[:bufN], 'S') {
				if d.setupWaiting {
					d.setupWaiting = false
					d.devS <- true
				}
			}
		} else {
			if err == serial.ErrTimeout {
				if d.Alive == true && d.Status == DEVICE_IN_USE {
					go func() {
						d.Alive = false
						deviceReset(d)
						deviceSetup(d)
					}()
				}
			} else {
				d.device.Close()
				break
			}
		}
	}
	for x := range deviceData.Devices {
		if deviceData.Devices[x].Index == d.Index {
			deviceData.Devices = append(deviceData.Devices[:x], deviceData.Devices[x+1:]...)
			break
		}
	}
}

func writeToAllDevices(content []byte) {
	for _, d := range deviceData.Devices {
		writeToDevice(d, content)
	}
}

func modifyDevicesStatus(index, status int) {
	for _, d := range deviceData.Devices {
		if d.Index == index {
			d.Status = status
			if d.Status == DEVICE_READY {
				deviceReset(d)
			} else if d.Status == DEVICE_IN_USE {
				deviceReset(d)
				deviceSetup(d)
			}
			break
		}
	}
}

func deviceResetConsole(d *Device) {
	devices[d].Lock()
	d.device.Write([]byte("P0"))
	time.Sleep(1000 * time.Millisecond)
	d.device.Write([]byte("P1"))
	devices[d].Unlock()
}

func deviceReset(d *Device) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	d.devR = make(chan bool)
	for {
		devices[d].Lock()
		d.resetWaiting = true
		d.device.Write([]byte("R"))
		select {
		case <-ticker.C:
		case <-d.devR:
			devices[d].Unlock()
			return
		}
		devices[d].Unlock()
	}
}

func deviceSetup(d *Device) {
	ticker := time.NewTicker(1000 * time.Millisecond)
	d.devS = make(chan bool)
	for {
		devices[d].Lock()
		d.setupWaiting = true
		d.device.Write([]byte("SAN\x80\x44A\x00"))
		select {
		case <-ticker.C:
			devices[d].Unlock()
			deviceReset(d)
		case <-d.devS:
			devices[d].Unlock()
			return
		}
	}
}

func devicesRestartTimer() {
	ticker := time.NewTicker(240000 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			for _, d := range deviceData.Devices {
				if d.Status == DEVICE_IN_USE {
					deviceResetConsole(d)
				}
			}
		}
	}
}
