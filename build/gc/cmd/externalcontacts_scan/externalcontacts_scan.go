package externalcontacts_scan

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_scan", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/scan/contacts")
	externalcontacts_scanCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_scan"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_scanCmd)
}

func Cmdexternalcontacts_scan() *cobra.Command {
	return externalcontacts_scanCmd
}
