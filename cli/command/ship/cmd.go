package ship

import (
	"github.com/spf13/cobra"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/service"
	"strings"
)

// NewShipCommand returns a cobra command for `service` subcommands
// nolint: interfacer
func NewShipCommand(dockerCli *command.DockerCli) *cobra.Command {
	cmdService := service.NewServiceCommand(dockerCli)
	transformToShip(cmdService)

	cmd := &cobra.Command{
		Use:   "ship",
		Short: "Manage ships",
		Args:  cli.NoArgs,
		RunE:  command.ShowHelp(dockerCli.Err()),
		Tags:  map[string]string{"version": "1.24"},
	}

	subCommands := cmdService.Commands()
	cmd.AddCommand(subCommands...)

	return cmd
}

func transformToShip(cmd *cobra.Command) {
	cmd.Use = serviceToShip(cmd.Use)
	cmd.Short = serviceToShip(cmd.Short)
	cmd.Long = serviceToShip(cmd.Long)
	cmd.Example = serviceToShip(cmd.Example)
}

func serviceToShip(serviceStr string) (shipStr string) {
	shipStr = serviceStr
	shipStr = strings.Replace(shipStr, "service", "ship", -1)
	shipStr = strings.Replace(shipStr, "Service", "Ship", -1)
	return
}
