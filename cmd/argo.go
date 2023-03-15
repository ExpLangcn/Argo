package main

import (
	"argo/pkg/engine"
	"argo/pkg/log"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/projectdiscovery/gologger"
	cli "github.com/urfave/cli/v2"
)

func SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, os.Interrupt, os.Kill, syscall.SIGKILL)
	go func() {
		<-c
		fmt.Println("ctrl+c exit")
		os.Exit(0)
	}()
}

func main() {
	SetupCloseHandler()
	app := cli.NewApp()
	app.Name = "argo"
	app.Authors = []*cli.Author{&cli.Author{Name: "Recar", Email: "https://github.com/Ciyfly"}}
	app.Usage = " -t http://testphp.vulnweb.com/"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "target",
			Aliases: []string{"t"},
			Value:   "",
			Usage:   "Specify the entry point for testing",
		},
		&cli.BoolFlag{
			Name:    "unheadless",
			Aliases: []string{"uh"},
			Value:   false,
			Usage:   "Is the default interface disabled? Specify 'uh' to enable the interface",
		},
		&cli.BoolFlag{
			Name:  "trace",
			Value: false,
			Usage: "Whether to display the elements of operation after opening the interface",
		},
		&cli.Float64Flag{
			Name:  "slow",
			Value: 1000,
			Usage: "The default delay time for operating after enabling ",
		},
		&cli.StringFlag{
			Name:    "username",
			Aliases: []string{"u"},
			Value:   "argo",
			Usage:   "If logging in, the default username ",
		},
		&cli.StringFlag{
			Name:    "password",
			Aliases: []string{"p"},
			Value:   "argo123",
			Usage:   "If logging in, the default password",
		},
		&cli.StringFlag{
			Name:  "email",
			Value: "argo@recar.com",
			Usage: "If logging in, the default email",
		},
		&cli.StringFlag{
			Name:  "phone",
			Value: "18888888888",
			Usage: "If logging in, the default phone",
		},
		&cli.StringFlag{
			Name:  "playback",
			Usage: "Support replay like headless YAML scripts",
		},
		&cli.BoolFlag{
			Name:  "testplayback",
			Usage: "If opened, then directly end after executing the specified playback script",
		},
		&cli.StringFlag{
			Name:  "proxy",
			Value: "",
			Usage: "Set up a proxy, for example, 127.0.0.1:3128",
		},
		&cli.IntFlag{
			Name:    "tabcount",
			Aliases: []string{"c"},
			Value:   10,
			Usage:   "The maximum number of tab pages that can be opened",
		},
		&cli.IntFlag{
			Name:  "tabtimeout",
			Value: 180,
			Usage: "Set the maximum running time for the tab, and close the tab if it exceeds the limit. The unit is in seconds",
		},
		&cli.IntFlag{
			Name:  "browsertimeout",
			Value: 3600,
			Usage: "Set the maximum running time for the browser, and close the browser if it exceeds the limit. The unit is in seconds",
		},
		&cli.StringFlag{
			Name:  "save",
			Usage: "The default name for the saved result is 'target' without a file extension. For example, to save as 'test', use the command '--save test'",
		},
		&cli.StringFlag{
			Name:  "format",
			Value: "txt,json",
			Usage: "Result output format separated by commas, multiple formats can be output at one time, and the supported formats include txt, json, xlsx, and html",
		},
		&cli.BoolFlag{
			Name:  "debug",
			Value: false,
			Usage: "Do you want to output debug information?",
		},
	}
	app.Action = RunMain

	err := app.Run(os.Args)
	if err != nil {
		gologger.Fatal().Msgf("cli.RunApp err: %s", err.Error())
		return
	}
}

func RunMain(c *cli.Context) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("painc err: %s\n", err)
		}
	}()
	target := c.String("target")
	if target == "" {
		fmt.Println("you need input target -h look look")
	}
	debug := c.Bool("debug")
	log.Init(debug)
	log.Logger.Info("[argo start]")
	// init
	engine.InitEngine(c)
	engine.Start()

	return nil
}