package main

import (
	f "fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func genKeyBank(path string) {
	var myPacket packet
	for i := 0; i < 100; i++ {
		os.MkdirAll("./"+path+"/"+strconv.Itoa(i), os.ModePerm)
		for j := 0; j < 100; j++ {
			os.Create("./" + path + "/" + strconv.Itoa(i) + "/" + strconv.Itoa(j))
			file, _ := os.OpenFile("./"+path+"/"+strconv.Itoa(i)+"/"+strconv.Itoa(j), os.O_RDWR, 0644)
			myPacket.initRand()
			formattedPacket := f.Sprintf("%x", myPacket.data)
			formattedPacket = strings.Replace(formattedPacket, " ", "", -1)
			file.WriteString(formattedPacket)
		}
	}
}

type packet struct {
	data []int
}

func (slice *packet) initRand() {
	slice.data = slice.data[:0]
	for i := 0; i < 1000000; i++ {
		slice.data = append(slice.data, rand.Intn(10))
	}
}
