package main

// YOLO

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "odns-updater"
	app.Usage = "Update OpenDNS ip"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "username, u",
			Usage:    "Username used for authentication",
			Required: true,
		},
		cli.StringFlag{
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

		if len(c.String("username")) > 0 {
			username = c.String("u")
		}

		if len(c.String("password")) > 0 {
			password = c.String("p")
		}

		if len(c.String("network")) > 0 {
			network = c.String("n")
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	basicAuth(username, password, network)
}

func basicAuth(username string, password string, network string) {
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
