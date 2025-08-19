package dataextensions

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("dataextensions", "SWAGGER_OVERRIDE_/api/v2/dataextensions")
	dataextensionsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dataextensions"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(dataextensionsCmd)
}

func Cmddataextensions() *cobra.Command {
	return dataextensionsCmd
}
