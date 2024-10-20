package externalcontacts_import

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("externalcontacts_import", "SWAGGER_OVERRIDE_/api/v2/externalcontacts/import")
	externalcontacts_importCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("externalcontacts_import"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(externalcontacts_importCmd)
}

func Cmdexternalcontacts_import() *cobra.Command {
	return externalcontacts_importCmd
}
