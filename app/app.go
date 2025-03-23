package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Consulta"
	app.Usage = "IP Servidores"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "ip",
			Aliases: []string{"i"},
			Usage:   "Busca IPS de endereços na internet",
			Flags:   flags,
			Action:  searchIps,
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Busca domínios na internet",
			Flags:   flags,
			Action:  searchServer,
		},
	}
	return app
}

func searchIps(ctx *cli.Context) {
	host := ctx.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(ctx *cli.Context) {
	host := ctx.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
