package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"crypto/md5"
)

//Если размер файлов небольшой
func defaultRead(){
	fmt.Println("<Start>" + "\n")
	fileInfo, err := ioutil.ReadDir("files")
	if err != nil{
		fmt.Println(err)
		return
	}

	for _, file := range fileInfo{
		envFileOpen, err := ioutil.ReadFile("files/" + file.Name())
		if err != nil{
			fmt.Println(err, "can't open file " + file.Name())
			continue
		}
		sum := md5.Sum(envFileOpen)
		fmt.Printf("%x --- File: %s\n", sum, file.Name())
	}
}
//Если файл огромный
func lastBytes(){
	fileInfo, err := ioutil.ReadDir("files")
	if err != nil{
		fmt.Println(err)
		return
	}

	for _, file := range fileInfo{
		fileOpenInfo, err := os.Stat("files/" + file.Name())
		if err != nil{
			fmt.Println(err, "can't open file " + file.Name())
			continue
		}
		fileOpen, err := os.Open("files/" + file.Name())
		defer fileOpen.Close()
		if err != nil {
			fmt.Println(err, "can't open file " + file.Name())
			continue
		}
		byteSlice := make([]byte, 16)
		_, err = fileOpen.ReadAt(byteSlice, fileOpenInfo.Size() - 16)
		if err != nil{
			fmt.Println(err,"can't open file " + file.Name())
			continue
		}
		fmt.Printf("%x --- File: %s\n", md5.Sum(byteSlice), file.Name())
	} 
}


func main(){
	//defaultRead()
	lastBytes()
}