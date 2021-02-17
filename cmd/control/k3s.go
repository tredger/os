package control

import (
	"github.com/burmilla/os/config"
	"github.com/burmilla/os/pkg/compose"
	"github.com/burmilla/os/pkg/log"

	"github.com/codegangsta/cli"
	"github.com/docker/libcompose/project/options"
	"golang.org/x/net/context"
)

func k3sSubcommands() []cli.Command {
	return []cli.Command{
		{
			Name:   "enable",
			Usage:  "enable/switch k3s",
			Action: k3sEnable,
		},
		{
			Name:   "disable",
			Usage:  "disable k3s",
			Action: k3sDisable,
		},
	}
}

func k3sEnable(c *cli.Context) error {
	cfg := config.LoadConfig()

	project, err := compose.GetProject(cfg, true, false, true)
	if err != nil {
		log.Fatal(err)
	}

	if err = project.Stop(context.Background(), 10, "k3s"); err != nil {
		log.Fatal(err)
	}

	if err = compose.LoadSpecialService(project, cfg, "k3s", "k3s"); err != nil {
		log.Fatal(err)
	}

	if err = project.Up(context.Background(), options.Up{}, "k3s"); err != nil {
		log.Fatal(err)
	}

	return nil
}

func k3sDisable(c *cli.Context) error {
	cfg := config.LoadConfig()

	project, err := compose.GetProject(cfg, true, false, true)
	if err != nil {
		log.Fatal(err)
	}

	if err = project.Down(context.Background(), options.Down{}, "k3s"); err != nil {
		log.Fatal(err)
	}
	return nil
}
