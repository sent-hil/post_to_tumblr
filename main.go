package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
	"strings"
)

const (
	logFileName string = "_.log"

	// put your posts in this dir
	postsDirName string = "posts"

	// this is my tumblr email address, you probably can't use this
	tumblrEmail string = "j8r8uvvjtpjn8@tumblr.com"
)

func main() {
	logFile, err := readOrCreateLogFile()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := logFile.Close()
		if err != nil {
			panic(err)
		}
	}()

	postsDir, err := readFromPostDir()
	if err != nil {
		panic(err)
	}

	var numberOfPosts int = len(postsDir)
	lastFile := postsDir[numberOfPosts-1]
	fullFileName := fmt.Sprintf("posts/%v", lastFile.Name())

	contents, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		panic(err)
	}

	joinedTitle := extractTitleFromFileName(lastFile.Name())
	err = sendEmail(joinedTitle, string(contents))
	if err != nil {
		panic(err)
	}
}

// In the future when this script will be run via cron job, this log file
// will be useful.
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

	log.Printf("found %v posts in dir %s", len(postsDir), postsDirName)

	return postsDir, nil
}

// Posts are expected to be in `1_Hello_world.md` format. First number
// indicates order. Everything after that till `.md` is assumed to be
// post title. This parses the _ and . and provides a human friendly
// title from the file.
func extractTitleFromFileName(fileName string) string {
	var nameWithoutExtension string
	var joinedTitle string

	splitFileNames := strings.Split(fileName, ".")
	nameWithoutExtension = splitFileNames[0]

	splitFileNames = strings.Split(nameWithoutExtension, "_")
	for _, str := range splitFileNames[1:] {
		joinedTitle += fmt.Sprintf("%s ", str)
	}

	log.Printf("parsed title: %s from file %s", joinedTitle, fileName)

	return joinedTitle
}

// Sends email via Gmail.
func sendEmail(subject, body string) error {
	from := flag.String("f", "", "Gmail username")
	password := flag.String("p", "", "Gmail password")
	flag.Parse()

	// tumblr looks for !m in subject so it can parse the email as markdown
	subject = subject + " !m"

	header := map[string]string{
		"From":         *from,
		"To":           tumblrEmail, // required for email to reach
		"Subject":      subject,
		"MIME-Version": "1.0",
	}

	var message string
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	message = fmt.Sprintf("%s\n%s", message, body)
	to := []string{tumblrEmail}
	auth := smtp.PlainAuth("", *from, *password, "smtp.gmail.com")
	err := smtp.SendMail("smtp.gmail.com:587", auth, *from, to, []byte(message))
	if err != nil {
		return err
	}

	log.Printf("email sent to %v", tumblrEmail)

	return nil
}
