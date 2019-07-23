package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type ContactBook map[string]string

func (cb ContactBook) AddContact(name, telephoneNumber string) ContactBook {
	cb[name] = telephoneNumber
	return cb
}

func (cb ContactBook) GetContact(name string) string {
	if telephoneNumber, found := cb[name]; found {
		return fmt.Sprintf("%s=%s", name, telephoneNumber)
	}
	return "Not found"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	var arrTemp []string
	var name, telephoneNumber string
	contactBook := make(ContactBook)
	for i := 0; i < n; i++ {
		arrTemp = strings.Split(readLine(reader), " ")
		name = arrTemp[0]
		telephoneNumber = arrTemp[1]
		contactBook = contactBook.AddContact(name, telephoneNumber)
	}
	for {
		name = readLine(reader)
		if name == "" {
			break
		}
		fmt.Println(contactBook.GetContact(name))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
