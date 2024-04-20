package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the command line application ready to be used
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application"
	app.Usage = "Search Ips and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "mercadolivre.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ips on the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "Servers where the ip are hosted",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func searchServers(c *cli.Context) {
	host := c.String("host")
	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers {
		fmt.Println(server)

	}

}
