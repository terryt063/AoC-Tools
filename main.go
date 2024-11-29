package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/joho/godotenv"
)

func main() {
	// Set the -d and -y flags to get the day and year for the challenges
	dayFlag := flag.String("d", "", "The day number")
	yearFlag := flag.String("y", "", "The year")
	// Set the flag variables
	flag.Parse()

	//Print the help text if not set...
	if *dayFlag == "" || *yearFlag == "" {
		fmt.Println("You must enter an input. The details are below")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *dayFlag == "all" {
		for i := 1; i < 26; i++ {
			err := writeFiles(*yearFlag, fmt.Sprint(i))
			if err != nil {
				log.Println(err.Error())
				os.Exit(1)
			}
			time.Sleep(200 * time.Nanosecond)
		}
	} else {
		err := writeFiles(*yearFlag, *dayFlag)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
	}

}

func writeFiles(year string, day string) error {
	fmt.Println("Getting the details for AoC", year, "Day", day, "...")
	// Create the directory for the file...
	err := os.MkdirAll(fmt.Sprintf("./%s/day%s/", year, day), os.ModePerm)
	if err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("./%s/day%s/day%s.md", year, day, day))
	if err != nil {
		return err
	}
	markdown, err := getInstructions(year, day)
	if err != nil {
		return err
	}
	file.WriteString(markdown)
	defer file.Close()

	// Now get the Input text
	inputFile, err := os.Create(fmt.Sprintf("./%s/day%s/input.txt", year, day))
	if err != nil {
		return err
	}
	inputText, err := getInput(year, day)
	if err != nil {
		return err
	}
	inputFile.WriteString(inputText)
	defer inputFile.Close()
	return nil

}

func requestData(url string) (*http.Response, error) {
	// Grab the session ID from the .env file...
	sessionID := loadSessionVariable()
	// setup the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic("Unable to create the request")
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: sessionID})
	client := &http.Client{}
	response, err := client.Do(req)
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %s", fmt.Sprint(response.StatusCode))
	}
	if err != nil {
		log.Panic("The request failed. Error -", err.Error())
	}
	return response, nil
}

func loadSessionVariable() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Cannot load session ID")
	}
	return os.Getenv("SESSION_ID")
}

func getInstructions(year string, day string) (string, error) {

	// Setup the request...
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s", year, day)

	response, err := requestData(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("status code: %s", fmt.Sprint(response.StatusCode))
	}

	bodyBytes, _ := io.ReadAll(response.Body)
	body := string(bodyBytes)
	// body := fmt.Sprintf("# Day %s\n\n %s", day, url)

	converter := md.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(body)
	if err != nil {
		return "", err
	}
	return markdown, nil

}

func getInput(year string, day string) (string, error) {

	// Setup the request...
	url := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)

	response, err := requestData(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("status code: %s", fmt.Sprint(response.StatusCode))
	}

	bodyBytes, _ := io.ReadAll(response.Body)
	body := string(bodyBytes)

	// body := fmt.Sprintf("# Day %s Input\n\n %s", day, url)

	return body, nil

}
