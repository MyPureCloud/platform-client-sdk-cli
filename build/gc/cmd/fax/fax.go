package fax

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("fax", "SWAGGER_OVERRIDE_/api/v2/fax")
	faxCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("fax"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(faxCmd)
}

func Cmdfax() *cobra.Command {
	return faxCmd
}
