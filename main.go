package main

// YOLO

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "odns-updater"
	app.Usage = "Update OpenDNS ip"
	app.Version = "1.0.0"
	app.HideVersion = true

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "username, u",
			Usage:    "Username used for authentication",
			Required: true,
		},
		cli.StringSliceFlag{
			Name:     "password, p",
			Usage:    "Password used for authentication",
			Required: true,
		},
		cli.StringFlag{
			Name:     "network, n",
			Usage:    "Name of network you want to update",
			Required: true,
		},
	}

	var username string
	var password string
	var network string

	app.Action = func(c *cli.Context) error {
		username = c.String("u")
		password = c.String("p")
		network = c.String("n")

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	match, _ := regexp.MatchString("\\[.+\\]", password)
	if match == true {
		replacer := strings.NewReplacer("[", "", "]", "")
		password = replacer.Replace(password)
	}

	BasicAuth(username, password, network)
}

func BasicAuth(username string, password string, network string) {
	reqUrl := fmt.Sprintf("https://updates.opendns.com/nic/update?hostname=%s", network)

	client := &http.Client{}

	req, err := http.NewRequest("GET", reqUrl, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
