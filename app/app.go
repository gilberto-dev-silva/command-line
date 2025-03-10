package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Consulta de IP e Whois de Domínios em Go"
	app.Usage = "Consulta de IP e Whois de Domínios em Go"

	app.Commands = []cli.Command{
		{
			Name:    "ip",
			Aliases: []string{"i"},
			Usage:   "Busca IPS de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com",
				},
			},
			Action: searchIps,
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
