package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	"net/smtp"
	"os"
	"strings"
)

const (
	logFileName  string = "_.log"
	postsDirName string = "posts"
	tumblrEmail  string = "j8r8uvvjtpjn8@tumblr.com"
)

func main() {
	logFile, err := readOrCreateLogFile()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := logFile.Close(); err != nil {
			panic(err)
		}
	}()

	postsDir, err := readFromPostDir()
	if err != nil {
		panic(err)
	}

	lastFile := postsDir[len(postsDir)-1]
	fullFileName := fmt.Sprintf("posts/%v", lastFile.Name())

	contents, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		panic(err)
	}

	username := flag.String("u", "", "Gmail username")
	password := flag.String("p", "", "Gmail password")
	flag.Parse()

	var subject string
	splitFileNames := strings.Split(lastFile.Name(), ".")
	if len(splitFileNames) == 0 {
		subject = splitFileNames[0]
	} else {
		subject = ""
	}

	err = sendEmail(*username, *password, subject, string(contents))
	if err != nil {
		panic(err)
	}
}

func sendEmail(username, password, title, body string) error {
	contents := []byte(fmt.Sprintf("Subject:%s \n %s", title, body))
	auth := smtp.PlainAuth("", username, password, "smtp.gmail.com")
	to := []string{tumblrEmail}

	err := smtp.SendMail("smtp.gmail.com:587", auth, username, to, contents)
	if err != nil {
		return err
	}

	return nil
}

func readOrCreateLogFile() (*os.File, error) {
	logFile, err := os.Open(logFileName)
	if err != nil {
		fmt.Println("no log file, creating new one")

		logFile, err = os.Create(logFileName)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("error creating log file", err))
		}

		fmt.Printf("created %s file \n", logFileName)
	}

	return logFile, nil
}

func readFromPostDir() ([]os.FileInfo, error) {
	postsDir, err := ioutil.ReadDir(postsDirName)
	if err != nil {
		return nil, err
	}

	if len(postsDir) == 0 {
		return nil, errors.New(fmt.Sprintf("no post files in %s directory", postsDirName))
	}

	return postsDir, nil
}
