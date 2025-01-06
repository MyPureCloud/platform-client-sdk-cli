package dataprivacy

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("dataprivacy", "SWAGGER_OVERRIDE_/api/v2/dataprivacy")
	dataprivacyCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("dataprivacy"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(dataprivacyCmd)
}

func Cmddataprivacy() *cobra.Command {
	return dataprivacyCmd
}
