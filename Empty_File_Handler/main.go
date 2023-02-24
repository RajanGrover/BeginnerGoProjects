// we gona create a project to find empty files in a directory and store the names of files
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Handler")
	mydir, _ := os.Getwd()
	fmt.Println(mydir)

	args := os.Args[1:]
	// Here we have to pass the directory
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}
	// now we gona use ReadDir to read all files in an given dir
	files, err := ioutil.ReadDir(args[0])
	// ReadDir return a []os.Fileinfo it conatains method like Name and Size
	if err != nil {
		fmt.Println(err)
		return
	}
	var names []byte // created empty slice to store empty file names

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')

			// fmt.Println("File Name:", name)
		}
	}
	err = ioutil.WriteFile("out.txt", names, 0644)
	//0644 is permission paramater like which users can read and write
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%s", names)
}

// to find the dir
// dir,_=os.Getwd()

//To optimize the code use the below code in for range loop

// var names []byte          // created empty slice to store empty file names
// 	total := len(files) * 256 // alocating the total memory at once
// 	names = make([]byte, 0, total)
// 	// better aproach
// 	for _, file := range files {
// 		if file.Size() == 0 {
// 			total += len(file.Name()) + 1
// 		}
// 	}
