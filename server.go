package main

import (
	f "fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	// "strings"
)

func genKeyBank(rootPath string) {
	// We need a million bytes per file for a MB
	const size = 1000000
	// Loop to generate a hundred folders
	for i := 0; i < 100; i++ {
		// l helps handle the first 10 folders
		l := strconv.Itoa(i)
		// If the folder name has only one digit (0-9)
		if i < 10 {
			// Add a 0 to make it of lenght two (00-09)
			l = "0" + l
		}
		// Create the folder structure:
		//               BankName  /  FolderNumber
		os.MkdirAll("./"+rootPath+"/"+l, os.ModePerm)
		// Loop to generate every file
		for j := 0; j < 100; j++ {
			// k helps handle the first 10 files
			k := strconv.Itoa(j)
			// If the file name has only one digit (0-9)
			if j < 10 {
				// Add a 0 to make it of lenght two (00-09)
				k = "0" + k
			}
			// Initialize a variable which will contain the data for the file
			packet := make([]byte, size) // variable will contain a million bytes
			// Create the random data for the file
			rand.Read(packet)
			// Use the first byte as an ID for the file in it's name
			name := strconv.Itoa(int(packet[0]))
			// Write the file:
			// KeyBank/FolderNumber/FileNumber-ID
			ioutil.WriteFile("./"+rootPath+"/"+l+"/"+k+"-"+name, packet, os.ModePerm)
		}
	}
}

func InitDB(rootPath string) {
	// All the 100 folders in the keybank directory
	allFolders, _ := ioutil.ReadDir(rootPath)
	// Select the first folder (alphabetically) from all of them
	firstFolder := allFolders[0].Name()
	// All the 100 files in the keybank directory
	allFiles, _ := ioutil.ReadDir(rootPath + "/" + firstFolder)
	// Select the first file (alphabetically) from all of them
	firstFile := allFiles[0].Name()
	// Read the data from the file and store it in memory
	keyData, _ := ioutil.ReadFile(rootPath + "/" + firstFolder + "/" + firstFile)
	// Delete the file since all it's data is in memoy
	os.Remove(rootPath + "/" + firstFolder + "/" + firstFile)
	// TODO: Delete the folder if it's empty
	// Delete the order number from the name and only keep the ID
	firstFile = firstFile[3:]
	// TODO: Check if the folder exists, change name if so...
	// Create the directory which will contain the first random data file
	os.MkdirAll(rootPath+"/../"+firstFile, os.ModePerm)
	// Write the file to the new folder
	ioutil.WriteFile(rootPath+"/../"+firstFile+"/"+firstFile, keyData, os.ModePerm)
}

func AppendDB(keyBankPath, ID string, MBsize int) {
	folderSize, _ := DirSize(keyBankPath + "/../" + ID)
	for wantedSize := int64(MBsize - 1); wantedSize > folderSize; {
		allFolders, _ := ioutil.ReadDir(keyBankPath)
		firstFolder := allFolders[0].Name()
		allFiles, _ := ioutil.ReadDir(keyBankPath + "/" + firstFolder)
		f.Printf("%v: Type: %T", len(allFiles), len(allFiles))
		firstFile := allFiles[0].Name()
		keyData, _ := ioutil.ReadFile(keyBankPath + "/" + firstFolder + "/" + firstFile)
		os.Remove(keyBankPath + "/" + firstFolder + "/" + firstFile)
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
