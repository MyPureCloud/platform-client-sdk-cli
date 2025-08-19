package license

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("license", "SWAGGER_OVERRIDE_/api/v2/license")
	licenseCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("license"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(licenseCmd)
}

func Cmdlicense() *cobra.Command {
	return licenseCmd
}
