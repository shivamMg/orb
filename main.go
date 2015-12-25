package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/nlopes/slack"
)

const (
	API_KEY             = "{{YOUR_SLACK_API_KEY}}"
	TARGET_CHANNEL_NAME = "brainteasers"
	SCHEDULED_HOUR      = 21
	SCHEDULED_MINUTE    = 30
)

var repoDir = os.Getenv("OPENSHIFT_REPO_DIR")
var assetsDir = repoDir + "prog_problems_assets/"

func main() {
	if !isScheduledTime() {
		return
	}

	api := slack.New(API_KEY)

	// Get channel ID from channel name
	targetChannelID := ""
	channels, _ := api.GetChannels(true)
	for _, ch := range channels {
		if ch.Name == TARGET_CHANNEL_NAME {
			targetChannelID = ch.ID
			break
		}
	}
	if targetChannelID == "" {
		fmt.Println(TARGET_CHANNEL_NAME, "does not exist")
		return
	}

	nextProblem := getNextProblem()
	nextProbFilePath := assetsDir + "problems/" + nextProblem

	params := slack.FileUploadParameters{
		Title:    "Programming Problem",
		Filetype: "post",
		File:     nextProbFilePath,
		Channels: []string{targetChannelID},
	}

	file, err := api.UploadFile(params)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Update last_problem.txt
	filepath := assetsDir + "last_problem.txt"
	err = ioutil.WriteFile(filepath, []byte(nextProblem), 0666)
	if err != nil {
		fmt.Println("Could not write to last_problem.txt:", err)
	}

	fmt.Println("Problem Name:", file.Name)
	fmt.Println("Uploaded file URL:", file.URL)
}

func isScheduledTime() bool {
	t := time.Now()
	indianTimeZone, err := time.LoadLocation("Asia/Calcutta")
	if err != nil {
		fmt.Println(err)
		return false
	}

	timeInIndia := t.In(indianTimeZone)
	hour := timeInIndia.Hour()
	min := timeInIndia.Minute()
	// For min, check if it is in range(30, 40)
	if !(hour == SCHEDULED_HOUR && min/10 == SCHEDULED_MINUTE/10) {
		fmt.Println("Not now")
		return false
	}
	return true
}
