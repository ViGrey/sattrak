package main

import (
	"io/ioutil"
	"strings"
)

func getOrbitSourceList() {
	content, _ := ioutil.ReadFile(config.OrbitSourceListPath)
	orbitURLList = strings.Split(string(content), "\n")
}
