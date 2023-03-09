package infrastructureascode

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("infrastructureascode", "SWAGGER_OVERRIDE_/api/v2/infrastructureascode")
	infrastructureascodeCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("infrastructureascode"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(infrastructureascodeCmd)
}

func Cmdinfrastructureascode() *cobra.Command {
	return infrastructureascodeCmd
}
