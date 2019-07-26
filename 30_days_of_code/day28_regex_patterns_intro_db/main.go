package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type DB struct {
	storage []string
}

func (db *DB) AddUser(firstName, email string) bool {
	validUser := ValidateEmail(email)
	if validUser {
		db.storage = append(db.storage, firstName)
	}
	return validUser
}

func (db DB) GetAllEmail() []string {
	sort.Strings(db.storage)
	return db.storage
}

func ValidateEmail(email string) bool {
	matched, _ := regexp.MatchString(".*@gmail.com", email)
	return matched
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	db := DB{}
	for NItr := 0; NItr < int(N); NItr++ {
		firstNameEmailID := strings.Split(readLine(reader), " ")
		firstName := firstNameEmailID[0]
		emailID := firstNameEmailID[1]
		db.AddUser(firstName, emailID)
	}
	for _, firstName := range db.GetAllEmail() {
		fmt.Println(firstName)
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
