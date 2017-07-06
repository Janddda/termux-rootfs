package cmd

import (
	"fmt"
	"os"

	"transfersh/server"

	"github.com/fatih/color"
	"github.com/minio/cli"
)

var globalFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "listener",
		Usage: "hostname and port for HTTP",
		Value: "127.0.0.1:8080",
	},
	cli.StringFlag{
		Name:  "tls-listener",
		Usage: "hostname and port for HTTPS",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tls-cert-file",
		Usage: "path to TLS certificate",
		Value: "",
	},
	cli.StringFlag{
		Name:  "tls-private-key",
		Usage: "path to TLS private key",
		Value: "",
	},
	cli.BoolFlag{
		Name:  "force-https",
		Usage: "redirect from HTTP to HTTPS",
	},
	cli.StringFlag{
		Name:  "basedir",
		Usage: "path to storage of uploaded files",
		Value: "",
	},
	cli.StringFlag{
		Name:  "temp-path",
		Usage: "path to temporary files",
		Value: os.TempDir(),
	},
	cli.IntFlag{
		Name:  "rate-limit",
		Usage: "limit requests per minute",
		Value: 0,
	},
	cli.BoolFlag{
		Name:  "profiler",
		Usage: "enable profiling",
	},
}

type Cmd struct {
	*cli.App
}

func New() *Cmd {
	app := cli.NewApp()
	app.Name = "transfer.sh"
	app.Version = "1.0"
	app.Author = ""
	app.Usage = "transfer.sh"
	app.Description = `Easy file sharing from the command line`
	app.Flags = globalFlags

	app.CustomAppHelpTemplate = `
 Usage: {{.Name}} {{if .Flags}}[flags]{{end}}

 {{.Description}}

 {{if .Flags}}Flags:
 {{range .Flags}}  {{.}}
 {{end}}{{end}}
`

	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.Action = func(c *cli.Context) {
		options := []server.OptionFn{}

		if v := c.String("listener"); v != "" {
			options = append(options, server.Listener(v))
		}

		if v := c.String("tls-listener"); v != "" {
			options = append(options, server.TLSListener(v))

			if cert := c.String("tls-cert-file"); cert == "" {
				fmt.Println("TLS certificate file is not set.")
				os.Exit(1)
			} else if pk := c.String("tls-private-key"); pk == "" {
				fmt.Println("TLS private key file is not set.")
				os.Exit(1)
			} else {
				options = append(options, server.TLSConfig(cert, pk))
			}

			if c.Bool("force-https") {
				options = append(options, server.ForceHTTPs())
			}
		} else {
			if c.String("tls-cert-file") != "" {
				fmt.Println("Flag '--tls-cert-file' requires tls-listener specified.")
				os.Exit(1)
			}

			if c.String("tls-private-key") != "" {
				fmt.Println("Flag '--tls-private-key' requires tls-listener specified.")
				os.Exit(1)
			}

			if c.Bool("force-https") {
				fmt.Println("Flag '--force-https' requires tls-listener specified.")
				os.Exit(1)
			}
		}

		if v := c.String("temp-path"); v != "" {
			options = append(options, server.TempPath(v))
		}

		if v := c.Int("rate-limit"); v > 0 {
			options = append(options, server.RateLimit(v))
		}

		if c.Bool("profiler") {
			options = append(options, server.EnableProfiler())
		}

		if v := c.String("basedir"); v == "" {
			fmt.Println("You need to set a directory where uploaded files will be stored.")
			fmt.Println("Use flag '--basedir'.")
			os.Exit(1)
		} else if storage, err := server.NewLocalStorage(v); err != nil {
			panic(err)
		} else {
			options = append(options, server.UseStorage(storage))
		}

		srvr, err := server.New(
			options...,
		)

		if err != nil {
			fmt.Println(color.RedString("Error starting server: %s", err.Error()))
			os.Exit(1)
		}

		srvr.Run()
	}

	return &Cmd{
		App: app,
	}
}
