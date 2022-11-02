package main

import (
	//"io/ioutil"
	"fmt"
)

func main() {
	dat := []byte{}
	for i := 0; i < 3*32; i++ {
		dat = append(dat, 0)
	}
	dat = append(dat, []byte{0, 0, 0x80}...)
	for i := 0; i < 26; i++ {
		dat = append(dat, 0x81)
	}
	dat = append(dat, []byte{0x82, 0, 0}...)
	for i := 0; i < 22; i++ {
		dat = append(dat, []byte{0, 0, 0x90}...)
		for j := 0; j < 26; j++ {
			dat = append(dat, 0)
		}
		dat = append(dat, []byte{0x92, 0, 0}...)
	}
	dat = append(dat, []byte{0, 0, 0xA0}...)
	for i := 0; i < 26; i++ {
		dat = append(dat, 0xA1)
	}
	dat = append(dat, []byte{0xA2, 0, 0}...)
	for i := 0; i < 3*32; i++ {
		dat = append(dat, 0)
	}
	fmt.Printf(string(dat))
}
