package main

import (
	"log"
	"os"
)

func check (e error){
	if e != nil {
		log.Println("Error: " + e.Error())
	}
}

func main(){
	lines := []byte("package main\n\nimport \"fmt\"\n\n")
	path, err := os.Getwd()
	filename := os.Args[1]
	filepath := path + "/" + filename
	f, err := os.Create(filepath)
	check(err)

	_, err = f.Write(lines)
	check(err)

	err = f.Sync()
	check(err)
}
