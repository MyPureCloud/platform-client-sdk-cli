package certificate

import (
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/utils"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/services"

	"github.com/spf13/cobra"
)

var (
	Description = utils.FormatUsageDescription("certificate", "SWAGGER_OVERRIDE_/api/v2/certificate")
	certificateCmd = &cobra.Command{
		Use:   utils.FormatUsageDescription("certificate"),
		Short: Description,
		Long:  Description,
	}
	CommandService services.CommandService
)

func init() {
	CommandService = services.NewCommandService(certificateCmd)
}

func Cmdcertificate() *cobra.Command {
	return certificateCmd
}
