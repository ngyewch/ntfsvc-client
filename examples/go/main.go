package main

import (
    "errors"
    "os"
    "github.com/ngyewch/go-ntfsvc/client"
)

func doMain() error {
    url := os.Getenv("NTFSVC_URL")
    if url == "" {
		return errors.New("NTFSVC_URL not specified")
    }

    apiKey := os.Getenv("NTFSVC_APIKEY")
    if apiKey == "" {
		return errors.New("NTFSVC_APIKEY not specified")
    }

	topic := os.Args[1]
	if topic == "" {
		return errors.New("topic not specified")
	}

	message := os.Args[2]
	if message == "" {
		return errors.New("message not specified")
	}

	notificationService := client.NewNotificationService(url, apiKey)
	err := notificationService.SendNotification(topic, message)
	if err != nil {
	    return err
	}

	return nil
}

func main() {
    err := doMain()
    if err != nil {
        panic(err)
    }
}
