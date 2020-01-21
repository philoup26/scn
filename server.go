package main

import (
	// f "fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	// "strings"
)

func genKeyBank(rootPath string) {
	const size = 1000000
	for i := 0; i < 100; i++ {
		os.MkdirAll("./"+rootPath+"/"+strconv.Itoa(i), os.ModePerm)
		for j := 0; j < 100; j++ {
			k := strconv.Itoa(j)
			if j < 10 {
				k = "0" + k
			}
			packet := make([]byte, size)
			rand.Read(packet)
			name := strconv.Itoa(int(packet[0]))
			ioutil.WriteFile("./"+rootPath+"/"+strconv.Itoa(i)+"/"+k+"-"+name, packet, os.ModePerm)
		}
	}
}

type packet struct {
	data []byte
}

func InitDB(rootPath string) {
	allFolders, _ := ioutil.ReadDir(rootPath)
	firstFolder := allFolders[0].Name()
	allFiles, _ := ioutil.ReadDir(rootPath + "/" + firstFolder)
	firstFile := allFiles[0].Name()
	keyData, _ := ioutil.ReadFile(rootPath + "/" + firstFolder + "/" + firstFile)
	os.Remove(rootPath + "/" + firstFolder + "/" + firstFile)
	// TODO: Delete the folder if it's empty
	if firstFile == nil{
		os.Remove(rootPath + "/" + firstFolder)
	}
	
	firstFile = firstFile[3:]
	// TODO: Check if the folder exists, change name if so...
	os.MkdirAll(rootPath+"/../"+firstFile, os.ModePerm)
	ioutil.WriteFile(rootPath+"/../"+firstFile+"/"+firstFile, keyData, os.ModePerm)
}

func AppendDB(keyBankPath, ID string, MBsize int) {
	folderSize, _ := DirSize(keyBankPath + "/../" + ID)
	for wantedSize := int64(MBsize - 1); wantedSize > folderSize; {
		allFolders, _ := ioutil.ReadDir(keyBankPath)
		firstFolder := allFolders[0].Name()
		allFiles, _ := ioutil.ReadDir(keyBankPath + "/" + firstFolder)
		firstFile := allFiles[0].Name()
		keyData, _ := ioutil.ReadFile(keyBankPath + "/" + firstFolder + "/" + firstFile)
		os.Remove(keyBankPath + "/" + firstFolder + "/" + firstFile)
		if firstFile == nil{
			os.Remove(rootPath + "/" + firstFolder)
		}
			
		firstFile = firstFile[3:]
		ioutil.WriteFile(keyBankPath+"/../"+ID+"/"+firstFile, keyData, os.ModePerm)
		folderSize, _ = DirSize(keyBankPath + "/../" + ID)
	}
}

func DirSize(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return (size / 1024.0 / 1024.0), err
}
