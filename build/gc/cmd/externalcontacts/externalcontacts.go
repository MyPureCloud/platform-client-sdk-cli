package externalcontacts

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts", "SWAGGER_OVERRIDE_/api/v2/externalcontacts")
	externalcontactsCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontactsCmd)
}

func Cmdexternalcontacts() *cobra.Command {
	return externalcontactsCmd
}
