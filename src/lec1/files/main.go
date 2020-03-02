package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	data, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))

	if err := ioutil.WriteFile("output.txt", []byte("output \nto file"), 0666); err != nil {
		fmt.Println(err)
		return
	}

	file3, err := os.Create("output2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file3.Close()
	for i := 0; i < 10; i++ {
		n, err := file3.WriteString(strconv.Itoa(i))
		fmt.Println(n, err)
		if err := file3.Sync(); err != nil {
			fmt.Println(err)
			break
		}
	}

}
