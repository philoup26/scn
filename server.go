package main

import (
	f "fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func genKeyBank(rootPath string) {
	var myPacket packet
	for i := 0; i < 100; i++ {
		os.MkdirAll("./"+rootPath+"/"+strconv.Itoa(i), os.ModePerm)
		for j := 0; j < 100; j++ {
			myPacket.initRand()
			k := strconv.Itoa(j)
			if j < 10 {
				k = "0" + k
			}
			name := strconv.Itoa(myPacket.data[0]) + strconv.Itoa(myPacket.data[1]) + strconv.Itoa(myPacket.data[2])
			os.Create("./" + rootPath + "/" + strconv.Itoa(i) + "/" + k + "-" + name)
			firstFile, _ := os.OpenFile("./"+rootPath+"/"+strconv.Itoa(i)+"/"+k+"-"+name, os.O_RDWR, os.ModePerm)
			formattedPacket := f.Sprintf("%x", myPacket.data[3:])
			formattedPacket = strings.Replace(formattedPacket, " ", "", -1)
			firstFile.WriteString(formattedPacket)
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

func InitDB(rootPath string) {
	allFolders, _ := ioutil.ReadDir(rootPath)
	firstFolder := allFolders[0].Name()
	allFiles, _ := ioutil.ReadDir(rootPath + "/" + firstFolder)
	firstFile := allFiles[0].Name()
	keyData, _ := ioutil.ReadFile(rootPath + "/" + firstFolder + "/" + firstFile)
	firstFile = firstFile[3:]
	os.MkdirAll(rootPath+"/../"+firstFile, os.ModePerm)
	ioutil.WriteFile(rootPath+"/../"+firstFile+"/"+firstFile, keyData, os.ModePerm)
}
