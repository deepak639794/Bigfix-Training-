package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func read() {

	f, err := os.Open("../Files/fixlets.csv")

	if err != nil {
		log.Fatal("Not able to open a file  ", err.Error())
	}
	file := csv.NewReader(f)
	if _, err1 := file.Read(); err1 != nil {
		panic("another error while reading the file ")
	}

	for {
		data, err := file.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			println("there is some error ", err.Error())
		}

		for i := 0; i < len(data); i++ {
			fmt.Print(data[i])
		}
		fmt.Println()

	}

}

func write() {
	//need to ask for the permission append and if the file does not exit then it will create a file
	f, err := os.OpenFile("../Files/fixlets.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	fmt.Println("writing into a file")
	writer := csv.NewWriter(f)

	arr := []string{"1fdsfsfdsf", "50121701000", "MS22-AUG: Security Update for Windows Server 2022 - Windows Server 2022 - KB50121701000 (x64)", "Critical", "96"}
	fmt.Print("writting to file")
	err1 := writer.Write(arr)

	if err1 != nil {
		log.Fatal("there is an error while writing to file ", err1.Error())
	}
	writer.Flush()
	if err2 := writer.Error(); err2 != nil {
		log.Fatal("there is some error", err2.Error())
	}
	fmt.Println("file written succesfully")

}

func delete() {

	f, err := os.Open("../Files/fixlets.csv")

	if err != nil {
		log.Fatal("there is some error while opening file ", err.Error())
	}

	reader := csv.NewReader(f)
	records, err1 := reader.ReadAll()

	if err1 != nil {
		log.Fatal(err1.Error())
	}
	//this will create a file or override the existing file
	newFile, err := os.Create("../Files/fixlets.csv")

	writer := csv.NewWriter(newFile)

	for indx, record := range records {
		if indx > len(records)-2 {
			break
		}
		err := writer.Write(record)

		if err != nil {
			log.Fatal(err.Error())
		}

	}
	writer.Flush()
	fmt.Println("<<<<<<<<<<the last record has been deleted>>>>>>>>>")

}

func main() {

	choice := 0

	for {
		fmt.Println("Enter your choice as listed")
		fmt.Print("1. Read the File \n2. Write to existing file \n3. Delete record \n->")

		fmt.Scan(&choice)
		if choice > 3 {
			fmt.Print("lkjfas")
			break
		}
		switch choice {
		case 1:
			read()
		case 2:
			write()
		case 3:
			delete()
		default:
			fmt.Print("wrong choice")
		}
	}

}
